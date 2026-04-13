// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cadenya/cadenya-go/internal/apijson"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/shared"
)

// UploadService issues short-lived presigned URLs for direct client-to-object-
// storage uploads at the WORKSPACE level. Created uploads can be referenced by id
// when creating or updating resources that accept binary content (e.g.,
// MemoryEntry).
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// UploadService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadService] method instead.
type UploadService struct {
	Options []option.RequestOption
}

// NewUploadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUploadService(opts ...option.RequestOption) (r *UploadService) {
	r = &UploadService{}
	r.Options = opts
	return
}

// Issues a short-lived presigned URL for direct upload to object storage. The
// returned id is used to reference the upload from resources that accept binary
// content.
func (r *UploadService) New(ctx context.Context, body UploadNewParams, opts ...option.RequestOption) (res *Upload, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/uploads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current state of an upload, including its lifecycle status
func (r *UploadService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Upload, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/uploads/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Upload is a workspace-scoped handle representing a single file upload flow.
// Clients call CreateUpload to receive a short-lived presigned URL, PUT the file
// directly to object storage, then reference the upload by id when creating or
// updating resources that accept binary content.
//
// Uploads are one-shot: once consumed by a creating or updating resource, the
// upload transitions to UPLOAD_STATUS_CONSUMED and cannot be reused. Unused
// uploads expire and are garbage-collected by the runtime.
type Upload struct {
	Info UploadInfo `json:"info" api:"required"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     UploadSpec              `json:"spec" api:"required"`
	JSON     uploadJSON              `json:"-"`
}

// uploadJSON contains the JSON metadata for the struct [Upload]
type uploadJSON struct {
	Info        apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Upload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadJSON) RawJSON() string {
	return r.raw
}

type UploadInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile `json:"createdBy"`
	// Lifecycle state. Transitions PENDING → COMPLETE (storage confirms the object
	// exists) → CONSUMED (a resource referenced this upload), or → EXPIRED (URL
	// elapsed without a PUT).
	Status UploadInfoStatus `json:"status"`
	// Presigned PUT URL. Short-lived. The client must PUT with the exact Content-Type
	// declared in the spec, and the body length must match size_bytes.
	UploadURL string `json:"uploadUrl"`
	// Absolute time at which upload_url stops working.
	UploadURLExpiresAt time.Time      `json:"uploadUrlExpiresAt" format:"date-time"`
	JSON               uploadInfoJSON `json:"-"`
}

// uploadInfoJSON contains the JSON metadata for the struct [UploadInfo]
type uploadInfoJSON struct {
	CreatedBy          apijson.Field
	Status             apijson.Field
	UploadURL          apijson.Field
	UploadURLExpiresAt apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UploadInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadInfoJSON) RawJSON() string {
	return r.raw
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

func (r UploadInfoStatus) IsKnown() bool {
	switch r {
	case UploadInfoStatusUploadStatusUnspecified, UploadInfoStatusUploadStatusPending, UploadInfoStatusUploadStatusComplete, UploadInfoStatusUploadStatusConsumed, UploadInfoStatusUploadStatusExpired:
		return true
	}
	return false
}

type UploadSpec struct {
	// MIME type the client will send. Baked into the presigned URL's signature — the
	// PUT must match exactly or object storage will reject it.
	ContentType string `json:"contentType" api:"required"`
	// Client-supplied filename. Used for audit and display only; does not control the
	// object's storage path.
	Filename string `json:"filename" api:"required"`
	// Expected size of the upload in bytes. Baked into the presigned URL as a
	// Content-Length constraint.
	SizeBytes string         `json:"sizeBytes" api:"required"`
	JSON      uploadSpecJSON `json:"-"`
}

// uploadSpecJSON contains the JSON metadata for the struct [UploadSpec]
type uploadSpecJSON struct {
	ContentType apijson.Field
	Filename    apijson.Field
	SizeBytes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UploadSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uploadSpecJSON) RawJSON() string {
	return r.raw
}

type UploadSpecParam struct {
	// MIME type the client will send. Baked into the presigned URL's signature — the
	// PUT must match exactly or object storage will reject it.
	ContentType param.Field[string] `json:"contentType" api:"required"`
	// Client-supplied filename. Used for audit and display only; does not control the
	// object's storage path.
	Filename param.Field[string] `json:"filename" api:"required"`
	// Expected size of the upload in bytes. Baked into the presigned URL as a
	// Content-Length constraint.
	SizeBytes param.Field[string] `json:"sizeBytes" api:"required"`
}

func (r UploadSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UploadNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	Spec     param.Field[UploadSpecParam]                    `json:"spec" api:"required"`
}

func (r UploadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
