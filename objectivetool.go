// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cadenya/cadenya-go/internal/apijson"
	"github.com/cadenya/cadenya-go/internal/apiquery"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/packages/pagination"
	"github.com/cadenya/cadenya-go/shared"
)

// ObjectiveToolService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveToolService] method instead.
type ObjectiveToolService struct {
	Options []option.RequestOption
}

// NewObjectiveToolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectiveToolService(opts ...option.RequestOption) (r *ObjectiveToolService) {
	r = &ObjectiveToolService{}
	r.Options = opts
	return
}

// Lists all tools that were assigned to an objective
func (r *ObjectiveToolService) List(ctx context.Context, objectiveID string, query ObjectiveToolListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveTool], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/tools", objectiveID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *ObjectiveToolService) ListAutoPaging(ctx context.Context, objectiveID string, query ObjectiveToolListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveTool] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, objectiveID, query, opts...))
}

// ObjectiveTool represents a tool that was assigned to an objective.
type ObjectiveTool struct {
	// BareMetadata contains the minimal metadata for a resource, including the ID.
	// These are used sparingly in Cadenya for resources where the full metadata is not
	// needed. You will come across them in list responses and other places where the
	// full metadata is not required like listing the tools that were assigned to an
	// objective. Because these types records are commonly created by other processes
	// in Cadenya, they do not have things like external IDs, labels, or names.
	Metadata shared.BareMetadata `json:"metadata" api:"required"`
	// Snapshot of the tool at the time it was assigned to the objective. Because tools
	// can change over time, snapshots are used to ensure tools don't change
	// unexpectedly during an objective's lifecycle.
	Snapshot Tool              `json:"snapshot"`
	JSON     objectiveToolJSON `json:"-"`
}

// objectiveToolJSON contains the JSON metadata for the struct [ObjectiveTool]
type objectiveToolJSON struct {
	Metadata    apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolJSON) RawJSON() string {
	return r.raw
}

type ObjectiveToolListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ObjectiveToolListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveToolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
