package cadenya

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.cadenya.com/cadenya-go/internal/apijson"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/param"
	"go.cadenya.com/cadenya-go/packages/respjson"
	"go.cadenya.com/cadenya-go/shared"
	"net/http"
	"net/url"
	"slices"
	"time"
)

// Issue short-lived presigned URLs for direct client-to-object-storage uploads.
// Created uploads can be referenced by id when creating or updating resources that
// accept binary content (e.g., MemoryEntry).
//
// UploadService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadService] method instead.
type UploadService struct {
	options []option.RequestOption
}

// NewUploadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUploadService(opts ...option.RequestOption) (r UploadService) {
	r = UploadService{}
	r.options = opts
	return
}

// Issues a short-lived presigned URL for direct upload to object storage. The
// returned id is used to reference the upload from resources that accept binary
// content.
func (r *UploadService) New(ctx context.Context, params UploadNewParams, opts ...option.RequestOption) (res *Upload, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.WorkspaceID, precfg.WorkspaceID)
	if params.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/uploads", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves the current state of an upload, including its lifecycle status
func (r *UploadService) Get(ctx context.Context, id string, query UploadGetParams, opts ...option.RequestOption) (res *Upload, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.WorkspaceID, precfg.WorkspaceID)
	if query.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/uploads/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A handle representing a single file upload flow. Clients call CreateUpload to
// receive a short-lived presigned URL, PUT the file directly to object storage,
// then reference the upload by id when creating or updating resources that accept
// binary content.
//
// Uploads are one-shot: once consumed by a creating or updating resource the
// upload transitions to UPLOAD_STATUS_CONSUMED and cannot be reused. Unused
// uploads expire and are garbage-collected.
type Upload struct {
	Info UploadInfo `json:"info" api:"required"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     UploadSpec              `json:"spec" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Upload) RawJSON() string { return r.JSON.raw }
func (r *Upload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Lifecycle state. Transitions PENDING → COMPLETE (storage confirms the object
	// exists) → CONSUMED (a resource referenced this upload), or → EXPIRED (URL
	// elapsed without a PUT).
	//
	// Any of "UPLOAD_STATUS_UNSPECIFIED", "UPLOAD_STATUS_PENDING",
	// "UPLOAD_STATUS_COMPLETE", "UPLOAD_STATUS_CONSUMED", "UPLOAD_STATUS_EXPIRED".
	Status UploadInfoStatus `json:"status"`
	// Presigned PUT URL. Short-lived. The client must PUT with the exact Content-Type
	// declared in the spec, and the body length must match size_bytes.
	UploadURL string `json:"uploadUrl"`
	// Absolute time at which upload_url stops working.
	UploadURLExpiresAt time.Time `json:"uploadUrlExpiresAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy          respjson.Field
		Status             respjson.Field
		UploadURL          respjson.Field
		UploadURLExpiresAt respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadInfo) RawJSON() string { return r.JSON.raw }
func (r *UploadInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state. Transitions PENDING → COMPLETE (storage confirms the object
// exists) → CONSUMED (a resource referenced this upload), or → EXPIRED (URL
// elapsed without a PUT).
type UploadInfoStatus string

const (
	UploadInfoStatusUploadStatusUnspecified UploadInfoStatus = "UPLOAD_STATUS_UNSPECIFIED"
	UploadInfoStatusUploadStatusPending     UploadInfoStatus = "UPLOAD_STATUS_PENDING"
	UploadInfoStatusUploadStatusComplete    UploadInfoStatus = "UPLOAD_STATUS_COMPLETE"
	UploadInfoStatusUploadStatusConsumed    UploadInfoStatus = "UPLOAD_STATUS_CONSUMED"
	UploadInfoStatusUploadStatusExpired     UploadInfoStatus = "UPLOAD_STATUS_EXPIRED"
)

type UploadSpec struct {
	// MIME type the client will send. Baked into the presigned URL's signature — the
	// PUT must match exactly or object storage will reject it.
	ContentType string `json:"contentType" api:"required"`
	// Client-supplied filename. Used for audit and display only; does not control the
	// object's storage path.
	Filename string `json:"filename" api:"required"`
	// Expected size of the upload in bytes. Baked into the presigned URL as a
	// Content-Length constraint.
	SizeBytes string `json:"sizeBytes" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Filename    respjson.Field
		SizeBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadSpec) RawJSON() string { return r.JSON.raw }
func (r *UploadSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UploadSpec to a UploadSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UploadSpecParam.Overrides()
func (r UploadSpec) ToParam() UploadSpecParam {
	return param.Override[UploadSpecParam](json.RawMessage(r.RawJSON()))
}

// The properties ContentType, Filename, SizeBytes are required.
type UploadSpecParam struct {
	// MIME type the client will send. Baked into the presigned URL's signature — the
	// PUT must match exactly or object storage will reject it.
	ContentType string `json:"contentType" api:"required"`
	// Client-supplied filename. Used for audit and display only; does not control the
	// object's storage path.
	Filename string `json:"filename" api:"required"`
	// Expected size of the upload in bytes. Baked into the presigned URL as a
	// Content-Length constraint.
	SizeBytes string `json:"sizeBytes" api:"required"`
	paramObj
}

func (r UploadSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow UploadSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UploadSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	Spec     UploadSpecParam                    `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r UploadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UploadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UploadNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
