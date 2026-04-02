// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/internal/apijson"
	"github.com/cadenya/cadenya-sdk-go/internal/param"
)

// AccountResourceMetadata is used to represent a resource that is associated to an
// account but not to a workspace.
type AccountResourceMetadata struct {
	// Unique identifier for the resource (prefixed ULID, e.g., "apikey_01HXK...")
	ID string `json:"id" api:"required"`
	// Account this resource belongs to for multi-tenant isolation (prefixed ULID)
	AccountID string `json:"accountId" api:"required"`
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool") Required for resources that users interact with directly
	Name      string `json:"name" api:"required"`
	ProfileID string `json:"profileId" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID string `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels map[string]string           `json:"labels"`
	JSON   accountResourceMetadataJSON `json:"-"`
}

// accountResourceMetadataJSON contains the JSON metadata for the struct
// [AccountResourceMetadata]
type accountResourceMetadataJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	ExternalID  apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountResourceMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountResourceMetadataJSON) RawJSON() string {
	return r.raw
}

// AccountResourceMetadata is used to represent a resource that is associated to an
// account but not to a workspace.
type AccountResourceMetadataParam struct {
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool") Required for resources that users interact with directly
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r AccountResourceMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// BareMetadata contains the minimal metadata for a resource, including the ID.
// These are used sparingly in Cadenya for resources where the full metadata is not
// needed. You will come across them in list responses and other places where the
// full metadata is not required like listing the tools that were assigned to an
// objective. Because these types records are commonly created by other processes
// in Cadenya, they do not have things like external IDs, labels, or names.
type BareMetadata struct {
	ID   string           `json:"id" api:"required"`
	JSON bareMetadataJSON `json:"-"`
}

// bareMetadataJSON contains the JSON metadata for the struct [BareMetadata]
type bareMetadataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BareMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bareMetadataJSON) RawJSON() string {
	return r.raw
}

// CallableTool is a union that represents a tool that can be called by an agent.
// In Cadenya, a tool that is used within an agent objective might be a
// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
// or a Cadenya Tool (one Cadenya provides).
type CallableTool struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent ResourceMetadata `json:"agent"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	CadenyaProvidedTool ResourceMetadata `json:"cadenyaProvidedTool"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Tool ResourceMetadata `json:"tool"`
	JSON callableToolJSON `json:"-"`
}

// callableToolJSON contains the JSON metadata for the struct [CallableTool]
type callableToolJSON struct {
	Agent               apijson.Field
	CadenyaProvidedTool apijson.Field
	Tool                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CallableTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callableToolJSON) RawJSON() string {
	return r.raw
}

// CreateOperationMetadata contains the user-provided fields for creating an
// operation. Read-only fields (id, account_id, workspace_id, created_at,
// profile_id) are excluded since they are set by the server.
type CreateOperationMetadataParam struct {
	// External ID for the operation (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"priority": "high", "source": "api", "workflow": "onboarding"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r CreateOperationMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CreateResourceMetadata contains the user-provided fields for creating a
// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
// profile_id, created_at) are excluded since they are set by the server.
type CreateResourceMetadataParam struct {
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool")
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r CreateResourceMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"priority": "high", "source": "api", "workflow": "onboarding"}
	Labels map[string]string     `json:"labels"`
	JSON   operationMetadataJSON `json:"-"`
}

// operationMetadataJSON contains the JSON metadata for the struct
// [OperationMetadata]
type operationMetadataJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	CreatedAt   apijson.Field
	ProfileID   apijson.Field
	WorkspaceID apijson.Field
	ExternalID  apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationMetadataJSON) RawJSON() string {
	return r.raw
}

// Profile represents a human user at the account level. Profiles are
// account-scoped resources that can be associated with multiple workspaces through
// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
type Profile struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata AccountResourceMetadata `json:"metadata" api:"required"`
	// ProfileSpec contains the profile-specific fields
	Spec ProfileSpec `json:"spec" api:"required"`
	JSON profileJSON `json:"-"`
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

// Profile represents a human user at the account level. Profiles are
// account-scoped resources that can be associated with multiple workspaces through
// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
type ProfileParam struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata param.Field[AccountResourceMetadataParam] `json:"metadata" api:"required"`
	// ProfileSpec contains the profile-specific fields
	Spec param.Field[ProfileSpecParam] `json:"spec" api:"required"`
}

func (r ProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ProfileSpec contains the profile-specific fields
type ProfileSpec struct {
	// Type is the type of profile. User's are humans, API keys are computers. You know
	// the deal.
	Type ProfileSpecType `json:"type" api:"required"`
	// Email address of the user (required, unique per account)
	Email string `json:"email"`
	// Display name for the user (e.g., "Bobby Tables")
	Name string          `json:"name"`
	JSON profileSpecJSON `json:"-"`
}

// profileSpecJSON contains the JSON metadata for the struct [ProfileSpec]
type profileSpecJSON struct {
	Type        apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileSpecJSON) RawJSON() string {
	return r.raw
}

// Type is the type of profile. User's are humans, API keys are computers. You know
// the deal.
type ProfileSpecType string

const (
	ProfileSpecTypeProfileTypeUser   ProfileSpecType = "PROFILE_TYPE_USER"
	ProfileSpecTypeProfileTypeAPIKey ProfileSpecType = "PROFILE_TYPE_API_KEY"
	ProfileSpecTypeProfileTypeSystem ProfileSpecType = "PROFILE_TYPE_SYSTEM"
)

func (r ProfileSpecType) IsKnown() bool {
	switch r {
	case ProfileSpecTypeProfileTypeUser, ProfileSpecTypeProfileTypeAPIKey, ProfileSpecTypeProfileTypeSystem:
		return true
	}
	return false
}

// ProfileSpec contains the profile-specific fields
type ProfileSpecParam struct {
	// Type is the type of profile. User's are humans, API keys are computers. You know
	// the deal.
	Type param.Field[ProfileSpecType] `json:"type" api:"required"`
	// Email address of the user (required, unique per account)
	Email param.Field[string] `json:"email"`
	// Display name for the user (e.g., "Bobby Tables")
	Name param.Field[string] `json:"name"`
}

func (r ProfileSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels map[string]string    `json:"labels"`
	JSON   resourceMetadataJSON `json:"-"`
}

// resourceMetadataJSON contains the JSON metadata for the struct
// [ResourceMetadata]
type resourceMetadataJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	WorkspaceID apijson.Field
	ExternalID  apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceMetadataJSON) RawJSON() string {
	return r.raw
}

// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
type ResourceMetadataParam struct {
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool") Required for resources that users interact with directly
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r ResourceMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UpdateResourceMetadata contains the user-provided fields for updating a
// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
// profile_id, created_at) are excluded since they are set by the server.
type UpdateResourceMetadataParam struct {
	// Human-readable name for the resource (e.g., "Customer Support Agent", "Email
	// Tool")
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r UpdateResourceMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Workspace struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata AccountResourceMetadata                `json:"metadata" api:"required"`
	Spec     gocadenyacomcadenyasdkgo.WorkspaceSpec `json:"spec" api:"required"`
	JSON     workspaceJSON                          `json:"-"`
}

// workspaceJSON contains the JSON metadata for the struct [Workspace]
type workspaceJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Workspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceJSON) RawJSON() string {
	return r.raw
}
