// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyago

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/cadenya/cadenya-go/internal/apijson"
	"github.com/cadenya/cadenya-go/internal/apiquery"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
)

// SearchService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Searches for tools or tool sets in the workspace
func (r *SearchService) SearchToolsOrToolSets(ctx context.Context, query SearchSearchToolsOrToolSetsParams, opts ...option.RequestOption) (res *SearchSearchToolsOrToolSetsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/search/tools_or_tool_sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SearchSearchToolsOrToolSetsResponse struct {
	Agents   []Agent                                 `json:"agents"`
	Tools    []Tool                                  `json:"tools"`
	ToolSets []ToolSet                               `json:"toolSets"`
	JSON     searchSearchToolsOrToolSetsResponseJSON `json:"-"`
}

// searchSearchToolsOrToolSetsResponseJSON contains the JSON metadata for the
// struct [SearchSearchToolsOrToolSetsResponse]
type searchSearchToolsOrToolSetsResponseJSON struct {
	Agents      apijson.Field
	Tools       apijson.Field
	ToolSets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchSearchToolsOrToolSetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchSearchToolsOrToolSetsResponseJSON) RawJSON() string {
	return r.raw
}

type SearchSearchToolsOrToolSetsParams struct {
	Query param.Field[string] `query:"query"`
}

// URLQuery serializes [SearchSearchToolsOrToolSetsParams]'s query parameters as
// `url.Values`.
func (r SearchSearchToolsOrToolSetsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
