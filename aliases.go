// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"github.com/cadenya/cadenya-sdk-go/internal/apierror"
	"github.com/cadenya/cadenya-sdk-go/shared"
)

type Error = apierror.Error

// AccountResourceMetadata is used to represent a resource that is associated to an
// account but not to a workspace.
//
// This is an alias to an internal type.
type AccountResourceMetadata = shared.AccountResourceMetadata

// AccountResourceMetadata is used to represent a resource that is associated to an
// account but not to a workspace.
//
// This is an alias to an internal type.
type AccountResourceMetadataParam = shared.AccountResourceMetadataParam

// BareMetadata contains the minimal metadata for a resource, including the ID.
// These are used sparingly in Cadenya for resources where the full metadata is not
// needed. You will come across them in list responses and other places where the
// full metadata is not required like listing the tools that were assigned to an
// objective. Because these types records are commonly created by other processes
// in Cadenya, they do not have things like external IDs, labels, or names.
//
// This is an alias to an internal type.
type BareMetadata = shared.BareMetadata

// CallableTool is a union that represents a tool that can be called by an agent.
// In Cadenya, a tool that is used within an agent objective might be a
// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
// or a Cadenya Tool (one Cadenya provides).
//
// This is an alias to an internal type.
type CallableTool = shared.CallableTool

// CreateOperationMetadata contains the user-provided fields for creating an
// operation. Read-only fields (id, account_id, workspace_id, created_at,
// profile_id) are excluded since they are set by the server.
//
// This is an alias to an internal type.
type CreateOperationMetadataParam = shared.CreateOperationMetadataParam

// CreateResourceMetadata contains the user-provided fields for creating a
// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
// profile_id, created_at) are excluded since they are set by the server.
//
// This is an alias to an internal type.
type CreateResourceMetadataParam = shared.CreateResourceMetadataParam

// Metadata for ephemeral operations and activities (e.g., objectives, executions,
// runs)
//
// This is an alias to an internal type.
type OperationMetadata = shared.OperationMetadata

// Profile represents a human user at the account level. Profiles are
// account-scoped resources that can be associated with multiple workspaces through
// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
//
// This is an alias to an internal type.
type Profile = shared.Profile

// Profile represents a human user at the account level. Profiles are
// account-scoped resources that can be associated with multiple workspaces through
// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
//
// This is an alias to an internal type.
type ProfileParam = shared.ProfileParam

// ProfileSpec contains the profile-specific fields
//
// This is an alias to an internal type.
type ProfileSpec = shared.ProfileSpec

// Type is the type of profile. User's are humans, API keys are computers. You know
// the deal.
//
// This is an alias to an internal type.
type ProfileSpecType = shared.ProfileSpecType

// This is an alias to an internal value.
const ProfileSpecTypeProfileTypeUser = shared.ProfileSpecTypeProfileTypeUser

// This is an alias to an internal value.
const ProfileSpecTypeProfileTypeAPIKey = shared.ProfileSpecTypeProfileTypeAPIKey

// This is an alias to an internal value.
const ProfileSpecTypeProfileTypeSystem = shared.ProfileSpecTypeProfileTypeSystem

// ProfileSpec contains the profile-specific fields
//
// This is an alias to an internal type.
type ProfileSpecParam = shared.ProfileSpecParam

// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
//
// This is an alias to an internal type.
type ResourceMetadata = shared.ResourceMetadata

// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
//
// This is an alias to an internal type.
type ResourceMetadataParam = shared.ResourceMetadataParam

// UpdateResourceMetadata contains the user-provided fields for updating a
// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
// profile_id, created_at) are excluded since they are set by the server.
//
// This is an alias to an internal type.
type UpdateResourceMetadataParam = shared.UpdateResourceMetadataParam

// This is an alias to an internal type.
type Workspace = shared.Workspace
