// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"go.cadenya.com/cadenya-go/internal/apijson"
	"go.cadenya.com/cadenya-go/packages/param"
	"go.cadenya.com/cadenya-go/packages/respjson"
	"time"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// AccountResourceMetadata is used to represent a resource that is associated to an
// account but not to a workspace.
type AccountResourceMetadata struct {
	// Unique identifier for the resource (prefixed ULID, e.g., "apikey_01HXK...")
	ID string `json:"id" api:"required"`
	// Account this resource belongs to for multi-tenant isolation (prefixed ULID)
	AccountID string `json:"accountId" api:"required"`
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool") Required for resources that users interact with directly
	Name      string    `json:"name" api:"required"`
	ProfileID string    `json:"profileId" api:"required"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID string `json:"externalId"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
	Labels map[string]string `json:"labels"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		Name        respjson.Field
		ProfileID   respjson.Field
		CreatedAt   respjson.Field
		ExternalID  respjson.Field
		Labels      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountResourceMetadata) RawJSON() string { return r.JSON.raw }
func (r *AccountResourceMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BareMetadata contains the minimal metadata for a resource: the ID and an
// optional human-readable name. These are used for reference fields where the full
// metadata (account scoping, timestamps, labels, external IDs) is not needed —
// e.g., the tool references inside an agent variation spec or the tools assigned
// to an objective. Both fields are server-populated; clients provide IDs through
// sibling fields rather than by constructing a BareMetadata themselves.
type BareMetadata struct {
	ID string `json:"id"`
	// Human-readable name of the referenced resource, populated by the server on reads
	// for convenience. Absent on references to resources that do not have a name
	// (e.g., objective tasks).
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BareMetadata) RawJSON() string { return r.JSON.raw }
func (r *BareMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreateOperationMetadata contains the user-provided fields for creating an
// operation. Read-only fields (id, account_id, workspace_id, created_at,
// profile_id) are excluded since they are set by the server.
type CreateOperationMetadataParam struct {
	// External ID for the operation (e.g., a workflow ID from an external system)
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"priority": "high", "source": "api",
	// "workflow": "onboarding"}
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r CreateOperationMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateOperationMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateOperationMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreateResourceMetadata contains the user-provided fields for creating a
// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
// profile_id, created_at) are excluded since they are set by the server.
//
// The property Name is required.
type CreateResourceMetadataParam struct {
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool")
	Name string `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r CreateResourceMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateResourceMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateResourceMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for ephemeral operations and activities (e.g., objectives, executions,
// runs)
type OperationMetadata struct {
	// Unique identifier for the operation (prefixed ULID, e.g., "obj_01HXK...")
	ID string `json:"id" api:"required"`
	// Account this operation belongs to for multi-tenant isolation (prefixed ULID)
	AccountID string `json:"accountId" api:"required"`
	// Timestamp when this operation was created ULID includes timestamp information,
	// but this explicit field enables easier querying
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// ID of the actor (user or service account) that created this operation
	ProfileID string `json:"profileId" api:"required"`
	// Workspace this operation belongs to for organizational grouping (prefixed ULID)
	WorkspaceID string `json:"workspaceId" api:"required"`
	// External ID for the operation (e.g., a workflow ID from an external system)
	ExternalID string `json:"externalId"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"priority": "high", "source": "api",
	// "workflow": "onboarding"}
	Labels map[string]string `json:"labels"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		CreatedAt   respjson.Field
		ProfileID   respjson.Field
		WorkspaceID respjson.Field
		ExternalID  respjson.Field
		Labels      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationMetadata) RawJSON() string { return r.JSON.raw }
func (r *OperationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
type ResourceMetadata struct {
	// Unique identifier for the resource (prefixed ULID, e.g., "agent_01HXK...")
	ID string `json:"id" api:"required"`
	// Account this resource belongs to for multi-tenant isolation (prefixed ULID)
	AccountID string `json:"accountId" api:"required"`
	// Timestamp when this resource was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool") Required for resources that users interact with directly
	Name string `json:"name" api:"required"`
	// ID of the actor (user or service account) that created this resource
	ProfileID string `json:"profileId" api:"required"`
	// Workspace this resource belongs to for organizational grouping (prefixed ULID)
	WorkspaceID string `json:"workspaceId" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID string `json:"externalId"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
	Labels map[string]string `json:"labels"`
	// Timestamp when this resource was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		ProfileID   respjson.Field
		WorkspaceID respjson.Field
		ExternalID  respjson.Field
		Labels      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceMetadata) RawJSON() string { return r.JSON.raw }
func (r *ResourceMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpdateResourceMetadata contains the user-provided fields for updating a
// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
// profile_id, created_at) are excluded since they are set by the server.
//
// The property Name is required.
type UpdateResourceMetadataParam struct {
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool")
	Name string `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r UpdateResourceMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateResourceMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateResourceMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
