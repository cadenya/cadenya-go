// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

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
)

// Issue, rotate, and revoke API keys for the account, and grant or revoke each
// key's access to individual workspaces.
//
// APIKeyAccessService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyAccessService] method instead.
type APIKeyAccessService struct {
	Options []option.RequestOption
}

// NewAPIKeyAccessService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIKeyAccessService(opts ...option.RequestOption) (r *APIKeyAccessService) {
	r = &APIKeyAccessService{}
	r.Options = opts
	return
}

// Lists the workspaces this API key has access to. Cursor-paginated.
func (r *APIKeyAccessService) List(ctx context.Context, id string, query APIKeyAccessListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Workspace], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/api_keys/%s/workspaces", id)
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

// Lists the workspaces this API key has access to. Cursor-paginated.
func (r *APIKeyAccessService) ListAutoPaging(ctx context.Context, id string, query APIKeyAccessListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Workspace] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, id, query, opts...))
}

// Grants this API key access to the specified workspace. Idempotent — adding an
// already-associated workspace is a no-op. Returns the updated API key with
// refreshed workspace preview and total.
func (r *APIKeyAccessService) Add(ctx context.Context, id string, body APIKeyAccessAddParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/api_keys/%s/workspaces", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Revokes this API key's access to the specified workspace. Idempotent. A key may
// have zero workspaces and remains valid.
func (r *APIKeyAccessService) Remove(ctx context.Context, id string, workspaceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	path := fmt.Sprintf("v1/account/api_keys/%s/workspaces/%s", id, workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type APIKeyAccessListParams struct {
	// Pagination cursor from previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [APIKeyAccessListParams]'s query parameters as `url.Values`.
func (r APIKeyAccessListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIKeyAccessAddParams struct {
	// The workspace to grant access to.
	WorkspaceID param.Field[string] `json:"workspaceId"`
}

func (r APIKeyAccessAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
