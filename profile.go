// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/cadenya/cadenya-go/internal/apiquery"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/packages/pagination"
)

// Read account profiles. Profiles are the account-level principals (users and API
// keys) that can be granted access to workspaces.
//
// ProfileService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r *ProfileService) {
	r = &ProfileService{}
	r.Options = opts
	return
}

// Lists the profiles in the current account. Supports free-form search and a type
// filter, intended for member-picker UIs (e.g. choosing a profile to add to a
// workspace).
func (r *ProfileService) List(ctx context.Context, query ProfileListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Profile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/profiles"
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

// Lists the profiles in the current account. Supports free-form search and a type
// filter, intended for member-picker UIs (e.g. choosing a profile to add to a
// workspace).
func (r *ProfileService) ListAutoPaging(ctx context.Context, query ProfileListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Profile] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

type ProfileListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Free-form search over profile name and email, for member-picker UIs.
	Query param.Field[string] `query:"query"`
	// Filter by profile type. Defaults to all types when unset; pass PROFILE_TYPE_USER
	// to list only human users (e.g. for a member picker).
	Type param.Field[ProfileListParamsType] `query:"type"`
}

// URLQuery serializes [ProfileListParams]'s query parameters as `url.Values`.
func (r ProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by profile type. Defaults to all types when unset; pass PROFILE_TYPE_USER
// to list only human users (e.g. for a member picker).
type ProfileListParamsType string

const (
	ProfileListParamsTypeProfileTypeUser   ProfileListParamsType = "PROFILE_TYPE_USER"
	ProfileListParamsTypeProfileTypeAPIKey ProfileListParamsType = "PROFILE_TYPE_API_KEY"
	ProfileListParamsTypeProfileTypeSystem ProfileListParamsType = "PROFILE_TYPE_SYSTEM"
)

func (r ProfileListParamsType) IsKnown() bool {
	switch r {
	case ProfileListParamsTypeProfileTypeUser, ProfileListParamsTypeProfileTypeAPIKey, ProfileListParamsTypeProfileTypeSystem:
		return true
	}
	return false
}
