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
	"github.com/cadenya/cadenya-go/shared"
)

// AIProviderKeyService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIProviderKeyService] method instead.
type AIProviderKeyService struct {
	Options []option.RequestOption
}

// NewAIProviderKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIProviderKeyService(opts ...option.RequestOption) (r *AIProviderKeyService) {
	r = &AIProviderKeyService{}
	r.Options = opts
	return
}

// Creates a new customer-provided AI provider key in the workspace
func (r *AIProviderKeyService) New(ctx context.Context, workspaceID string, body AIProviderKeyNewParams, opts ...option.RequestOption) (res *AIProviderKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an AI provider key by ID from the workspace
func (r *AIProviderKeyService) Get(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (res *AIProviderKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an AI provider key's name or key value in the workspace
func (r *AIProviderKeyService) Update(ctx context.Context, workspaceID string, id string, body AIProviderKeyUpdateParams, opts ...option.RequestOption) (res *AIProviderKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all customer-provided AI provider keys in the workspace
func (r *AIProviderKeyService) List(ctx context.Context, workspaceID string, query AIProviderKeyListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AIProviderKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys", workspaceID)
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

// Lists all customer-provided AI provider keys in the workspace
func (r *AIProviderKeyService) ListAutoPaging(ctx context.Context, workspaceID string, query AIProviderKeyListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AIProviderKey] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Deletes an AI provider key from the workspace
func (r *AIProviderKeyService) Delete(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// AIProviderKey is a customer-provided (BYOK) credential for an AI provider,
// scoped to a workspace. The secret value is never returned in responses.
type AIProviderKey struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     AIProviderKeySpec       `json:"spec" api:"required"`
	JSON     aiProviderKeyJSON       `json:"-"`
}

// aiProviderKeyJSON contains the JSON metadata for the struct [AIProviderKey]
type aiProviderKeyJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeyJSON) RawJSON() string {
	return r.raw
}

type AIProviderKeySpec struct {
	// The provider credential. Accepted on create/update; never populated in responses
	// (the server returns an empty value to avoid leaking it).
	APIKey string `json:"apiKey"`
	// The AI provider this key authenticates against. Currently "openrouter".
	Provider string `json:"provider"`
	// The provider region. "us" or "eu". Defaults to "us".
	Region string                `json:"region"`
	JSON   aiProviderKeySpecJSON `json:"-"`
}

// aiProviderKeySpecJSON contains the JSON metadata for the struct
// [AIProviderKeySpec]
type aiProviderKeySpecJSON struct {
	APIKey      apijson.Field
	Provider    apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKeySpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecJSON) RawJSON() string {
	return r.raw
}

type AIProviderKeySpecParam struct {
	// The provider credential. Accepted on create/update; never populated in responses
	// (the server returns an empty value to avoid leaking it).
	APIKey param.Field[string] `json:"apiKey"`
	// The AI provider this key authenticates against. Currently "openrouter".
	Provider param.Field[string] `json:"provider"`
	// The provider region. "us" or "eu". Defaults to "us".
	Region param.Field[string] `json:"region"`
}

func (r AIProviderKeySpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIProviderKeyNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	Spec     param.Field[AIProviderKeySpecParam]             `json:"spec" api:"required"`
}

func (r AIProviderKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIProviderKeyUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	Spec     param.Field[AIProviderKeySpecParam]             `json:"spec"`
	// Fields to update.
	UpdateMask param.Field[string] `json:"updateMask" format:"field-mask"`
}

func (r AIProviderKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIProviderKeyListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [AIProviderKeyListParams]'s query parameters as
// `url.Values`.
func (r AIProviderKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
