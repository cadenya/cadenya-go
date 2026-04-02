// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyago

import (
	"github.com/cadenya/cadenya-go/internal/apierror"
	"github.com/cadenya/cadenya-go/shared"
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
