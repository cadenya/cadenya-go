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

// AIProviderKey is a credential for an AI provider, scoped to a workspace. Most
// keys are customer-provided (BYOK); Cadenya also provisions promotional keys (see
// AIProviderKeyInfo.is_promotional), which cannot be modified or deleted by
// account administrators. The secret value is never returned in responses.
type AIProviderKey struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     AIProviderKeySpec       `json:"spec" api:"required"`
	// AIProviderKeyInfo carries server-derived, read-only details about a key, for AI
	// provider management UIs.
	Info AIProviderKeyInfo `json:"info"`
	JSON aiProviderKeyJSON `json:"-"`
}

// aiProviderKeyJSON contains the JSON metadata for the struct [AIProviderKey]
type aiProviderKeyJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeyJSON) RawJSON() string {
	return r.raw
}

// AIProviderKeyInfo carries server-derived, read-only details about a key, for AI
// provider management UIs.
type AIProviderKeyInfo struct {
	// Number of disabled models provisioned on this key.
	DisabledModelCount int64 `json:"disabledModelCount"`
	// Number of enabled models provisioned on this key.
	EnabledModelCount int64 `json:"enabledModelCount"`
	// Cadenya includes promotional keys (one for onboarding, and potentially more in
	// the future). These are not added or maintained by account administrators.
	IsPromotional bool                  `json:"isPromotional"`
	JSON          aiProviderKeyInfoJSON `json:"-"`
}

// aiProviderKeyInfoJSON contains the JSON metadata for the struct
// [AIProviderKeyInfo]
type aiProviderKeyInfoJSON struct {
	DisabledModelCount apijson.Field
	EnabledModelCount  apijson.Field
	IsPromotional      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AIProviderKeyInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeyInfoJSON) RawJSON() string {
	return r.raw
}

type AIProviderKeySpec struct {
	// AIProviderConfig holds non-secret, provider-specific settings. The set case must
	// correspond to AIProviderKeySpec.provider. Providers with no settings (Anthropic,
	// Gemini) simply leave this unset. The endpoint of a named provider is fixed and
	// intentionally not overridable here; use the OpenAI-compatible provider to target
	// a custom endpoint.
	Config AIProviderKeySpecConfig `json:"config"`
	// AIProviderCredential is the secret material used to authenticate with a
	// provider. The set case must correspond to AIProviderKeySpec.provider. The server
	// encrypts the serialized message at rest and never returns it on reads.
	Credentials AIProviderKeySpecCredentials `json:"credentials"`
	// The AI provider this key authenticates against.
	Provider AIProviderKeySpecProvider `json:"provider"`
	JSON     aiProviderKeySpecJSON     `json:"-"`
}

// aiProviderKeySpecJSON contains the JSON metadata for the struct
// [AIProviderKeySpec]
type aiProviderKeySpecJSON struct {
	Config      apijson.Field
	Credentials apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKeySpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecJSON) RawJSON() string {
	return r.raw
}

// AIProviderConfig holds non-secret, provider-specific settings. The set case must
// correspond to AIProviderKeySpec.provider. Providers with no settings (Anthropic,
// Gemini) simply leave this unset. The endpoint of a named provider is fixed and
// intentionally not overridable here; use the OpenAI-compatible provider to target
// a custom endpoint.
type AIProviderKeySpecConfig struct {
	// OpenAIConfig holds OpenAI-specific settings.
	OpenAI AIProviderKeySpecConfigOpenAI `json:"openai"`
	// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
	// Completions API. The base URL is required and its model catalog is discovered
	// live via GET {base_url}/models.
	OpenAICompatible AIProviderKeySpecConfigOpenAICompatible `json:"openaiCompatible"`
	// OpenRouterConfig holds OpenRouter-specific settings.
	Openrouter AIProviderKeySpecConfigOpenrouter `json:"openrouter"`
	JSON       aiProviderKeySpecConfigJSON       `json:"-"`
}

// aiProviderKeySpecConfigJSON contains the JSON metadata for the struct
// [AIProviderKeySpecConfig]
type aiProviderKeySpecConfigJSON struct {
	OpenAI           apijson.Field
	OpenAICompatible apijson.Field
	Openrouter       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AIProviderKeySpecConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecConfigJSON) RawJSON() string {
	return r.raw
}

// OpenAIConfig holds OpenAI-specific settings.
type AIProviderKeySpecConfigOpenAI struct {
	// Sent as the OpenAI-Organization header when set.
	OrganizationID string `json:"organizationId"`
	// Sent as the OpenAI-Project header when set.
	ProjectID string                            `json:"projectId"`
	JSON      aiProviderKeySpecConfigOpenAIJSON `json:"-"`
}

// aiProviderKeySpecConfigOpenAIJSON contains the JSON metadata for the struct
// [AIProviderKeySpecConfigOpenAI]
type aiProviderKeySpecConfigOpenAIJSON struct {
	OrganizationID apijson.Field
	ProjectID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIProviderKeySpecConfigOpenAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecConfigOpenAIJSON) RawJSON() string {
	return r.raw
}

// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
// Completions API. The base URL is required and its model catalog is discovered
// live via GET {base_url}/models.
type AIProviderKeySpecConfigOpenAICompatible struct {
	BaseURL string                                      `json:"baseUrl"`
	JSON    aiProviderKeySpecConfigOpenAICompatibleJSON `json:"-"`
}

// aiProviderKeySpecConfigOpenAICompatibleJSON contains the JSON metadata for the
// struct [AIProviderKeySpecConfigOpenAICompatible]
type aiProviderKeySpecConfigOpenAICompatibleJSON struct {
	BaseURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKeySpecConfigOpenAICompatible) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecConfigOpenAICompatibleJSON) RawJSON() string {
	return r.raw
}

// OpenRouterConfig holds OpenRouter-specific settings.
type AIProviderKeySpecConfigOpenrouter struct {
	// Data-residency region (e.g. "us", "eu"). Empty uses the provider default.
	Region string                                `json:"region"`
	JSON   aiProviderKeySpecConfigOpenrouterJSON `json:"-"`
}

// aiProviderKeySpecConfigOpenrouterJSON contains the JSON metadata for the struct
// [AIProviderKeySpecConfigOpenrouter]
type aiProviderKeySpecConfigOpenrouterJSON struct {
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKeySpecConfigOpenrouter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecConfigOpenrouterJSON) RawJSON() string {
	return r.raw
}

// AIProviderCredential is the secret material used to authenticate with a
// provider. The set case must correspond to AIProviderKeySpec.provider. The server
// encrypts the serialized message at rest and never returns it on reads.
type AIProviderKeySpecCredentials struct {
	// CredentialAPIKey carries a single bearer/header API key.
	APIKey AIProviderKeySpecCredentialsAPIKey `json:"apiKey"`
	// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
	// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
	Headers AIProviderKeySpecCredentialsHeaders `json:"headers"`
	JSON    aiProviderKeySpecCredentialsJSON    `json:"-"`
}

// aiProviderKeySpecCredentialsJSON contains the JSON metadata for the struct
// [AIProviderKeySpecCredentials]
type aiProviderKeySpecCredentialsJSON struct {
	APIKey      apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKeySpecCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecCredentialsJSON) RawJSON() string {
	return r.raw
}

// CredentialAPIKey carries a single bearer/header API key.
type AIProviderKeySpecCredentialsAPIKey struct {
	APIKey string                                 `json:"apiKey"`
	JSON   aiProviderKeySpecCredentialsAPIKeyJSON `json:"-"`
}

// aiProviderKeySpecCredentialsAPIKeyJSON contains the JSON metadata for the struct
// [AIProviderKeySpecCredentialsAPIKey]
type aiProviderKeySpecCredentialsAPIKeyJSON struct {
	APIKey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKeySpecCredentialsAPIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecCredentialsAPIKeyJSON) RawJSON() string {
	return r.raw
}

// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
type AIProviderKeySpecCredentialsHeaders struct {
	Headers map[string]string                       `json:"headers"`
	JSON    aiProviderKeySpecCredentialsHeadersJSON `json:"-"`
}

// aiProviderKeySpecCredentialsHeadersJSON contains the JSON metadata for the
// struct [AIProviderKeySpecCredentialsHeaders]
type aiProviderKeySpecCredentialsHeadersJSON struct {
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIProviderKeySpecCredentialsHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiProviderKeySpecCredentialsHeadersJSON) RawJSON() string {
	return r.raw
}

// The AI provider this key authenticates against.
type AIProviderKeySpecProvider string

const (
	AIProviderKeySpecProviderAIProviderUnspecified      AIProviderKeySpecProvider = "AI_PROVIDER_UNSPECIFIED"
	AIProviderKeySpecProviderAIProviderOpenrouter       AIProviderKeySpecProvider = "AI_PROVIDER_OPENROUTER"
	AIProviderKeySpecProviderAIProviderOpenAI           AIProviderKeySpecProvider = "AI_PROVIDER_OPENAI"
	AIProviderKeySpecProviderAIProviderAnthropic        AIProviderKeySpecProvider = "AI_PROVIDER_ANTHROPIC"
	AIProviderKeySpecProviderAIProviderGemini           AIProviderKeySpecProvider = "AI_PROVIDER_GEMINI"
	AIProviderKeySpecProviderAIProviderOpenAICompatible AIProviderKeySpecProvider = "AI_PROVIDER_OPENAI_COMPATIBLE"
)

func (r AIProviderKeySpecProvider) IsKnown() bool {
	switch r {
	case AIProviderKeySpecProviderAIProviderUnspecified, AIProviderKeySpecProviderAIProviderOpenrouter, AIProviderKeySpecProviderAIProviderOpenAI, AIProviderKeySpecProviderAIProviderAnthropic, AIProviderKeySpecProviderAIProviderGemini, AIProviderKeySpecProviderAIProviderOpenAICompatible:
		return true
	}
	return false
}

type AIProviderKeySpecParam struct {
	// AIProviderConfig holds non-secret, provider-specific settings. The set case must
	// correspond to AIProviderKeySpec.provider. Providers with no settings (Anthropic,
	// Gemini) simply leave this unset. The endpoint of a named provider is fixed and
	// intentionally not overridable here; use the OpenAI-compatible provider to target
	// a custom endpoint.
	Config param.Field[AIProviderKeySpecConfigParam] `json:"config"`
	// AIProviderCredential is the secret material used to authenticate with a
	// provider. The set case must correspond to AIProviderKeySpec.provider. The server
	// encrypts the serialized message at rest and never returns it on reads.
	Credentials param.Field[AIProviderKeySpecCredentialsParam] `json:"credentials"`
	// The AI provider this key authenticates against.
	Provider param.Field[AIProviderKeySpecProvider] `json:"provider"`
}

func (r AIProviderKeySpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AIProviderConfig holds non-secret, provider-specific settings. The set case must
// correspond to AIProviderKeySpec.provider. Providers with no settings (Anthropic,
// Gemini) simply leave this unset. The endpoint of a named provider is fixed and
// intentionally not overridable here; use the OpenAI-compatible provider to target
// a custom endpoint.
type AIProviderKeySpecConfigParam struct {
	// OpenAIConfig holds OpenAI-specific settings.
	OpenAI param.Field[AIProviderKeySpecConfigOpenAIParam] `json:"openai"`
	// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
	// Completions API. The base URL is required and its model catalog is discovered
	// live via GET {base_url}/models.
	OpenAICompatible param.Field[AIProviderKeySpecConfigOpenAICompatibleParam] `json:"openaiCompatible"`
	// OpenRouterConfig holds OpenRouter-specific settings.
	Openrouter param.Field[AIProviderKeySpecConfigOpenrouterParam] `json:"openrouter"`
}

func (r AIProviderKeySpecConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OpenAIConfig holds OpenAI-specific settings.
type AIProviderKeySpecConfigOpenAIParam struct {
	// Sent as the OpenAI-Organization header when set.
	OrganizationID param.Field[string] `json:"organizationId"`
	// Sent as the OpenAI-Project header when set.
	ProjectID param.Field[string] `json:"projectId"`
}

func (r AIProviderKeySpecConfigOpenAIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
// Completions API. The base URL is required and its model catalog is discovered
// live via GET {base_url}/models.
type AIProviderKeySpecConfigOpenAICompatibleParam struct {
	BaseURL param.Field[string] `json:"baseUrl"`
}

func (r AIProviderKeySpecConfigOpenAICompatibleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OpenRouterConfig holds OpenRouter-specific settings.
type AIProviderKeySpecConfigOpenrouterParam struct {
	// Data-residency region (e.g. "us", "eu"). Empty uses the provider default.
	Region param.Field[string] `json:"region"`
}

func (r AIProviderKeySpecConfigOpenrouterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AIProviderCredential is the secret material used to authenticate with a
// provider. The set case must correspond to AIProviderKeySpec.provider. The server
// encrypts the serialized message at rest and never returns it on reads.
type AIProviderKeySpecCredentialsParam struct {
	// CredentialAPIKey carries a single bearer/header API key.
	APIKey param.Field[AIProviderKeySpecCredentialsAPIKeyParam] `json:"apiKey"`
	// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
	// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
	Headers param.Field[AIProviderKeySpecCredentialsHeadersParam] `json:"headers"`
}

func (r AIProviderKeySpecCredentialsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CredentialAPIKey carries a single bearer/header API key.
type AIProviderKeySpecCredentialsAPIKeyParam struct {
	APIKey param.Field[string] `json:"apiKey"`
}

func (r AIProviderKeySpecCredentialsAPIKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
type AIProviderKeySpecCredentialsHeadersParam struct {
	Headers param.Field[map[string]string] `json:"headers"`
}

func (r AIProviderKeySpecCredentialsHeadersParam) MarshalJSON() (data []byte, err error) {
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
	// When true, populate each item's info (model counts), at the cost of extra
	// lookups.
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Field[string] `query:"labels"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// When true, return only promotional keys (provided by Cadenya, e.g. for
	// onboarding). Defaults to returning all keys, customer-provided and promotional
	// alike.
	Promotional param.Field[bool] `query:"promotional"`
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
