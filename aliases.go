// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

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

// BareMetadata contains the minimal metadata for a resource: the ID and an
// optional human-readable name. These are used for reference fields where the full
// metadata (account scoping, timestamps, labels, external IDs) is not needed —
// e.g., the tool references inside an agent variation spec or the tools assigned
// to an objective. Both fields are server-populated; clients provide IDs through
// sibling fields rather than by constructing a BareMetadata themselves.
//
// This is an alias to an internal type.
type BareMetadata = shared.BareMetadata

// BareMetadata contains the minimal metadata for a resource: the ID and an
// optional human-readable name. These are used for reference fields where the full
// metadata (account scoping, timestamps, labels, external IDs) is not needed —
// e.g., the tool references inside an agent variation spec or the tools assigned
// to an objective. Both fields are server-populated; clients provide IDs through
// sibling fields rather than by constructing a BareMetadata themselves.
//
// This is an alias to an internal type.
type BareMetadataParam = shared.BareMetadataParam

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
