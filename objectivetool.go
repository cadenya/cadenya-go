// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"go.cadenya.com/cadenya-go/internal/apijson"
	"go.cadenya.com/cadenya-go/internal/apiquery"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/pagination"
	"go.cadenya.com/cadenya-go/packages/param"
	"go.cadenya.com/cadenya-go/packages/respjson"
	"go.cadenya.com/cadenya-go/shared"
	"net/http"
	"net/url"
	"slices"
)

// ObjectiveToolService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveToolService] method instead.
type ObjectiveToolService struct {
	options []option.RequestOption
}

// NewObjectiveToolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectiveToolService(opts ...option.RequestOption) (r ObjectiveToolService) {
	r = ObjectiveToolService{}
	r.options = opts
	return
}

// Lists all tools that were assigned to an objective
func (r *ObjectiveToolService) List(ctx context.Context, objectiveID string, params ObjectiveToolListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveTool], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.WorkspaceID, precfg.WorkspaceID)
	if params.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tools", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all tools that were assigned to an objective
func (r *ObjectiveToolService) ListAutoPaging(ctx context.Context, objectiveID string, params ObjectiveToolListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveTool] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, objectiveID, params, opts...))
}

// ObjectiveTool represents a tool that was assigned to an objective.
type ObjectiveTool struct {
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Metadata shared.BareMetadata `json:"metadata" api:"required"`
	// Snapshot of the tool at the time it was assigned to the objective. Because tools
	// can change over time, snapshots are used to ensure tools don't change
	// unexpectedly during an objective's lifecycle.
	Snapshot Tool `json:"snapshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Snapshot    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveTool) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectiveToolListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
