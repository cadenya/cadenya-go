package cadenya

import (
	"context"
	"encoding/json"
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

// AIProviderKeyService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIProviderKeyService] method instead.
type AIProviderKeyService struct {
	options []option.RequestOption
}

// NewAIProviderKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIProviderKeyService(opts ...option.RequestOption) (r AIProviderKeyService) {
	r = AIProviderKeyService{}
	r.options = opts
	return
}

// Creates a new customer-provided AI provider key in the workspace
func (r *AIProviderKeyService) New(ctx context.Context, params AIProviderKeyNewParams, opts ...option.RequestOption) (res *AIProviderKey, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.WorkspaceID, precfg.WorkspaceID)
	if params.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves an AI provider key by ID from the workspace
func (r *AIProviderKeyService) Get(ctx context.Context, id string, query AIProviderKeyGetParams, opts ...option.RequestOption) (res *AIProviderKey, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.WorkspaceID, precfg.WorkspaceID)
	if query.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an AI provider key's name or key value in the workspace
func (r *AIProviderKeyService) Update(ctx context.Context, id string, params AIProviderKeyUpdateParams, opts ...option.RequestOption) (res *AIProviderKey, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.WorkspaceID, precfg.WorkspaceID)
	if params.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all customer-provided AI provider keys in the workspace
func (r *AIProviderKeyService) List(ctx context.Context, params AIProviderKeyListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AIProviderKey], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys", url.PathEscape(params.WorkspaceID.Value))
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

// Lists all customer-provided AI provider keys in the workspace
func (r *AIProviderKeyService) ListAutoPaging(ctx context.Context, params AIProviderKeyListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AIProviderKey] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes an AI provider key from the workspace
func (r *AIProviderKeyService) Delete(ctx context.Context, id string, body AIProviderKeyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&body.WorkspaceID, precfg.WorkspaceID)
	if body.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/ai_provider_keys/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type AIProviderConfigOpenAI struct {
	// OpenAIConfig holds OpenAI-specific settings.
	OpenAI AIProviderConfigOpenAIOpenAI `json:"openai" api:"required"`
	// Any of "openai".
	Type AIProviderConfigOpenAIType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenAI      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderConfigOpenAI) RawJSON() string { return r.JSON.raw }
func (r *AIProviderConfigOpenAI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIProviderConfigOpenAI to a AIProviderConfigOpenAIParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIProviderConfigOpenAIParam.Overrides()
func (r AIProviderConfigOpenAI) ToParam() AIProviderConfigOpenAIParam {
	return param.Override[AIProviderConfigOpenAIParam](json.RawMessage(r.RawJSON()))
}

// OpenAIConfig holds OpenAI-specific settings.
type AIProviderConfigOpenAIOpenAI struct {
	// Sent as the OpenAI-Organization header when set.
	OrganizationID string `json:"organizationId"`
	// Sent as the OpenAI-Project header when set.
	ProjectID string `json:"projectId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrganizationID respjson.Field
		ProjectID      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderConfigOpenAIOpenAI) RawJSON() string { return r.JSON.raw }
func (r *AIProviderConfigOpenAIOpenAI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderConfigOpenAIType string

const (
	AIProviderConfigOpenAITypeOpenAI AIProviderConfigOpenAIType = "openai"
)

// The properties OpenAI, Type are required.
type AIProviderConfigOpenAIParam struct {
	// OpenAIConfig holds OpenAI-specific settings.
	OpenAI AIProviderConfigOpenAIOpenAIParam `json:"openai,omitzero" api:"required"`
	// Any of "openai".
	Type AIProviderConfigOpenAIType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AIProviderConfigOpenAIParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderConfigOpenAIParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderConfigOpenAIParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIConfig holds OpenAI-specific settings.
type AIProviderConfigOpenAIOpenAIParam struct {
	// Sent as the OpenAI-Organization header when set.
	OrganizationID param.Opt[string] `json:"organizationId,omitzero"`
	// Sent as the OpenAI-Project header when set.
	ProjectID param.Opt[string] `json:"projectId,omitzero"`
	paramObj
}

func (r AIProviderConfigOpenAIOpenAIParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderConfigOpenAIOpenAIParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderConfigOpenAIOpenAIParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderConfigOpenAICompatible struct {
	// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
	// Completions API. The base URL is required and its model catalog is discovered
	// live via GET {base_url}/models.
	OpenAICompatible AIProviderConfigOpenAICompatibleOpenAICompatible `json:"openaiCompatible" api:"required"`
	// Any of "openaiCompatible".
	Type AIProviderConfigOpenAICompatibleType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenAICompatible respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderConfigOpenAICompatible) RawJSON() string { return r.JSON.raw }
func (r *AIProviderConfigOpenAICompatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIProviderConfigOpenAICompatible to a
// AIProviderConfigOpenAICompatibleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIProviderConfigOpenAICompatibleParam.Overrides()
func (r AIProviderConfigOpenAICompatible) ToParam() AIProviderConfigOpenAICompatibleParam {
	return param.Override[AIProviderConfigOpenAICompatibleParam](json.RawMessage(r.RawJSON()))
}

// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
// Completions API. The base URL is required and its model catalog is discovered
// live via GET {base_url}/models.
type AIProviderConfigOpenAICompatibleOpenAICompatible struct {
	BaseURL string `json:"baseUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderConfigOpenAICompatibleOpenAICompatible) RawJSON() string { return r.JSON.raw }
func (r *AIProviderConfigOpenAICompatibleOpenAICompatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderConfigOpenAICompatibleType string

const (
	AIProviderConfigOpenAICompatibleTypeOpenAICompatible AIProviderConfigOpenAICompatibleType = "openaiCompatible"
)

// The properties OpenAICompatible, Type are required.
type AIProviderConfigOpenAICompatibleParam struct {
	// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
	// Completions API. The base URL is required and its model catalog is discovered
	// live via GET {base_url}/models.
	OpenAICompatible AIProviderConfigOpenAICompatibleOpenAICompatibleParam `json:"openaiCompatible,omitzero" api:"required"`
	// Any of "openaiCompatible".
	Type AIProviderConfigOpenAICompatibleType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AIProviderConfigOpenAICompatibleParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderConfigOpenAICompatibleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderConfigOpenAICompatibleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
// Completions API. The base URL is required and its model catalog is discovered
// live via GET {base_url}/models.
//
// The property BaseURL is required.
type AIProviderConfigOpenAICompatibleOpenAICompatibleParam struct {
	BaseURL string `json:"baseUrl" api:"required"`
	paramObj
}

func (r AIProviderConfigOpenAICompatibleOpenAICompatibleParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderConfigOpenAICompatibleOpenAICompatibleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderConfigOpenAICompatibleOpenAICompatibleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderConfigOpenrouter struct {
	// OpenRouterConfig holds OpenRouter-specific settings.
	Openrouter AIProviderConfigOpenrouterOpenrouter `json:"openrouter" api:"required"`
	// Any of "openrouter".
	Type AIProviderConfigOpenrouterType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Openrouter  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderConfigOpenrouter) RawJSON() string { return r.JSON.raw }
func (r *AIProviderConfigOpenrouter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIProviderConfigOpenrouter to a
// AIProviderConfigOpenrouterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIProviderConfigOpenrouterParam.Overrides()
func (r AIProviderConfigOpenrouter) ToParam() AIProviderConfigOpenrouterParam {
	return param.Override[AIProviderConfigOpenrouterParam](json.RawMessage(r.RawJSON()))
}

// OpenRouterConfig holds OpenRouter-specific settings.
type AIProviderConfigOpenrouterOpenrouter struct {
	// Data-residency region (e.g. "us", "eu"). Empty uses the provider default.
	Region string `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderConfigOpenrouterOpenrouter) RawJSON() string { return r.JSON.raw }
func (r *AIProviderConfigOpenrouterOpenrouter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderConfigOpenrouterType string

const (
	AIProviderConfigOpenrouterTypeOpenrouter AIProviderConfigOpenrouterType = "openrouter"
)

// The properties Openrouter, Type are required.
type AIProviderConfigOpenrouterParam struct {
	// OpenRouterConfig holds OpenRouter-specific settings.
	Openrouter AIProviderConfigOpenrouterOpenrouterParam `json:"openrouter,omitzero" api:"required"`
	// Any of "openrouter".
	Type AIProviderConfigOpenrouterType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AIProviderConfigOpenrouterParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderConfigOpenrouterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderConfigOpenrouterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenRouterConfig holds OpenRouter-specific settings.
type AIProviderConfigOpenrouterOpenrouterParam struct {
	// Data-residency region (e.g. "us", "eu"). Empty uses the provider default.
	Region param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r AIProviderConfigOpenrouterOpenrouterParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderConfigOpenrouterOpenrouterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderConfigOpenrouterOpenrouterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderCredentialAPIKey struct {
	// CredentialAPIKey carries a single bearer/header API key.
	APIKey AIProviderCredentialAPIKeyAPIKey `json:"apiKey" api:"required"`
	// Any of "apiKey".
	Type AIProviderCredentialAPIKeyType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderCredentialAPIKey) RawJSON() string { return r.JSON.raw }
func (r *AIProviderCredentialAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIProviderCredentialAPIKey to a
// AIProviderCredentialAPIKeyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIProviderCredentialAPIKeyParam.Overrides()
func (r AIProviderCredentialAPIKey) ToParam() AIProviderCredentialAPIKeyParam {
	return param.Override[AIProviderCredentialAPIKeyParam](json.RawMessage(r.RawJSON()))
}

// CredentialAPIKey carries a single bearer/header API key.
type AIProviderCredentialAPIKeyAPIKey struct {
	APIKey string `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderCredentialAPIKeyAPIKey) RawJSON() string { return r.JSON.raw }
func (r *AIProviderCredentialAPIKeyAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderCredentialAPIKeyType string

const (
	AIProviderCredentialAPIKeyTypeAPIKey AIProviderCredentialAPIKeyType = "apiKey"
)

// The properties APIKey, Type are required.
type AIProviderCredentialAPIKeyParam struct {
	// CredentialAPIKey carries a single bearer/header API key.
	APIKey AIProviderCredentialAPIKeyAPIKeyParam `json:"apiKey,omitzero" api:"required"`
	// Any of "apiKey".
	Type AIProviderCredentialAPIKeyType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AIProviderCredentialAPIKeyParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderCredentialAPIKeyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderCredentialAPIKeyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialAPIKey carries a single bearer/header API key.
type AIProviderCredentialAPIKeyAPIKeyParam struct {
	APIKey param.Opt[string] `json:"apiKey,omitzero"`
	paramObj
}

func (r AIProviderCredentialAPIKeyAPIKeyParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderCredentialAPIKeyAPIKeyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderCredentialAPIKeyAPIKeyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderCredentialHeaders struct {
	// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
	// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
	Headers AIProviderCredentialHeadersHeaders `json:"headers" api:"required"`
	// Any of "headers".
	Type AIProviderCredentialHeadersType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderCredentialHeaders) RawJSON() string { return r.JSON.raw }
func (r *AIProviderCredentialHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIProviderCredentialHeaders to a
// AIProviderCredentialHeadersParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIProviderCredentialHeadersParam.Overrides()
func (r AIProviderCredentialHeaders) ToParam() AIProviderCredentialHeadersParam {
	return param.Override[AIProviderCredentialHeadersParam](json.RawMessage(r.RawJSON()))
}

// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
type AIProviderCredentialHeadersHeaders struct {
	Headers map[string]string `json:"headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderCredentialHeadersHeaders) RawJSON() string { return r.JSON.raw }
func (r *AIProviderCredentialHeadersHeaders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderCredentialHeadersType string

const (
	AIProviderCredentialHeadersTypeHeaders AIProviderCredentialHeadersType = "headers"
)

// The properties Headers, Type are required.
type AIProviderCredentialHeadersParam struct {
	// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
	// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
	Headers AIProviderCredentialHeadersHeadersParam `json:"headers,omitzero" api:"required"`
	// Any of "headers".
	Type AIProviderCredentialHeadersType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AIProviderCredentialHeadersParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderCredentialHeadersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderCredentialHeadersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
type AIProviderCredentialHeadersHeadersParam struct {
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r AIProviderCredentialHeadersHeadersParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderCredentialHeadersHeadersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderCredentialHeadersHeadersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKey) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	IsPromotional bool `json:"isPromotional"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisabledModelCount respjson.Field
		EnabledModelCount  respjson.Field
		IsPromotional      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKeyInfo) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderKeySpec struct {
	// AIProviderConfig holds non-secret, provider-specific settings. The set case must
	// correspond to AIProviderKeySpec.provider. Providers with no settings (Anthropic,
	// Gemini) simply leave this unset. The endpoint of a named provider is fixed and
	// intentionally not overridable here; use the OpenAI-compatible provider to target
	// a custom endpoint.
	Config AIProviderKeySpecConfigUnion `json:"config"`
	// AIProviderCredential is the secret material used to authenticate with a
	// provider. The set case must correspond to AIProviderKeySpec.provider. The server
	// encrypts the serialized message at rest and never returns it on reads.
	Credentials AIProviderKeySpecCredentialsUnion `json:"credentials"`
	// The AI provider this key authenticates against.
	//
	// Any of "AI_PROVIDER_UNSPECIFIED", "AI_PROVIDER_OPENROUTER",
	// "AI_PROVIDER_OPENAI", "AI_PROVIDER_ANTHROPIC", "AI_PROVIDER_GEMINI",
	// "AI_PROVIDER_OPENAI_COMPATIBLE".
	Provider AIProviderKeySpecProvider `json:"provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Credentials respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKeySpec) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIProviderKeySpec to a AIProviderKeySpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIProviderKeySpecParam.Overrides()
func (r AIProviderKeySpec) ToParam() AIProviderKeySpecParam {
	return param.Override[AIProviderKeySpecParam](json.RawMessage(r.RawJSON()))
}

// AIProviderKeySpecConfigUnion contains all possible properties and values from
// [AIProviderConfigOpenrouter], [AIProviderConfigOpenAI],
// [AIProviderConfigOpenAICompatible].
//
// Use the [AIProviderKeySpecConfigUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AIProviderKeySpecConfigUnion struct {
	// This field is from variant [AIProviderConfigOpenrouter].
	Openrouter AIProviderConfigOpenrouterOpenrouter `json:"openrouter"`
	// Any of "openrouter", "openai", "openaiCompatible".
	Type string `json:"type"`
	// This field is from variant [AIProviderConfigOpenAI].
	OpenAI AIProviderConfigOpenAIOpenAI `json:"openai"`
	// This field is from variant [AIProviderConfigOpenAICompatible].
	OpenAICompatible AIProviderConfigOpenAICompatibleOpenAICompatible `json:"openaiCompatible"`
	JSON             struct {
		Openrouter       respjson.Field
		Type             respjson.Field
		OpenAI           respjson.Field
		OpenAICompatible respjson.Field
		raw              string
	} `json:"-"`
}

// anyAIProviderKeySpecConfig is implemented by each variant of
// [AIProviderKeySpecConfigUnion] to add type safety for the return type of
// [AIProviderKeySpecConfigUnion.AsAny]
type anyAIProviderKeySpecConfig interface {
	implAIProviderKeySpecConfigUnion()
}

func (AIProviderConfigOpenrouter) implAIProviderKeySpecConfigUnion()       {}
func (AIProviderConfigOpenAI) implAIProviderKeySpecConfigUnion()           {}
func (AIProviderConfigOpenAICompatible) implAIProviderKeySpecConfigUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AIProviderKeySpecConfigUnion.AsAny().(type) {
//	case cadenya.AIProviderConfigOpenrouter:
//	case cadenya.AIProviderConfigOpenAI:
//	case cadenya.AIProviderConfigOpenAICompatible:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AIProviderKeySpecConfigUnion) AsAny() anyAIProviderKeySpecConfig {
	switch u.Type {
	case "openrouter":
		return u.AsOpenrouter()
	case "openai":
		return u.AsOpenAI()
	case "openaiCompatible":
		return u.AsOpenAICompatible()
	}
	return nil
}

func (u AIProviderKeySpecConfigUnion) AsOpenrouter() (v AIProviderConfigOpenrouter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIProviderKeySpecConfigUnion) AsOpenAI() (v AIProviderConfigOpenAI) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIProviderKeySpecConfigUnion) AsOpenAICompatible() (v AIProviderConfigOpenAICompatible) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AIProviderKeySpecConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *AIProviderKeySpecConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AIProviderKeySpecCredentialsUnion contains all possible properties and values
// from [AIProviderCredentialAPIKey], [AIProviderCredentialHeaders].
//
// Use the [AIProviderKeySpecCredentialsUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AIProviderKeySpecCredentialsUnion struct {
	// This field is from variant [AIProviderCredentialAPIKey].
	APIKey AIProviderCredentialAPIKeyAPIKey `json:"apiKey"`
	// Any of "apiKey", "headers".
	Type string `json:"type"`
	// This field is from variant [AIProviderCredentialHeaders].
	Headers AIProviderCredentialHeadersHeaders `json:"headers"`
	JSON    struct {
		APIKey  respjson.Field
		Type    respjson.Field
		Headers respjson.Field
		raw     string
	} `json:"-"`
}

// anyAIProviderKeySpecCredentials is implemented by each variant of
// [AIProviderKeySpecCredentialsUnion] to add type safety for the return type of
// [AIProviderKeySpecCredentialsUnion.AsAny]
type anyAIProviderKeySpecCredentials interface {
	implAIProviderKeySpecCredentialsUnion()
}

func (AIProviderCredentialAPIKey) implAIProviderKeySpecCredentialsUnion()  {}
func (AIProviderCredentialHeaders) implAIProviderKeySpecCredentialsUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AIProviderKeySpecCredentialsUnion.AsAny().(type) {
//	case cadenya.AIProviderCredentialAPIKey:
//	case cadenya.AIProviderCredentialHeaders:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AIProviderKeySpecCredentialsUnion) AsAny() anyAIProviderKeySpecCredentials {
	switch u.Type {
	case "apiKey":
		return u.AsAPIKey()
	case "headers":
		return u.AsHeaders()
	}
	return nil
}

func (u AIProviderKeySpecCredentialsUnion) AsAPIKey() (v AIProviderCredentialAPIKey) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIProviderKeySpecCredentialsUnion) AsHeaders() (v AIProviderCredentialHeaders) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AIProviderKeySpecCredentialsUnion) RawJSON() string { return u.JSON.raw }

func (r *AIProviderKeySpecCredentialsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type AIProviderKeySpecParam struct {
	// AIProviderConfig holds non-secret, provider-specific settings. The set case must
	// correspond to AIProviderKeySpec.provider. Providers with no settings (Anthropic,
	// Gemini) simply leave this unset. The endpoint of a named provider is fixed and
	// intentionally not overridable here; use the OpenAI-compatible provider to target
	// a custom endpoint.
	Config AIProviderKeySpecConfigUnionParam `json:"config,omitzero"`
	// AIProviderCredential is the secret material used to authenticate with a
	// provider. The set case must correspond to AIProviderKeySpec.provider. The server
	// encrypts the serialized message at rest and never returns it on reads.
	Credentials AIProviderKeySpecCredentialsUnionParam `json:"credentials,omitzero"`
	// The AI provider this key authenticates against.
	//
	// Any of "AI_PROVIDER_UNSPECIFIED", "AI_PROVIDER_OPENROUTER",
	// "AI_PROVIDER_OPENAI", "AI_PROVIDER_ANTHROPIC", "AI_PROVIDER_GEMINI",
	// "AI_PROVIDER_OPENAI_COMPATIBLE".
	Provider AIProviderKeySpecProvider `json:"provider,omitzero"`
	paramObj
}

func (r AIProviderKeySpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIProviderKeySpecConfigUnionParam struct {
	OfOpenrouter       *AIProviderConfigOpenrouterParam       `json:",omitzero,inline"`
	OfOpenAI           *AIProviderConfigOpenAIParam           `json:",omitzero,inline"`
	OfOpenAICompatible *AIProviderConfigOpenAICompatibleParam `json:",omitzero,inline"`
	paramUnion
}

func (u AIProviderKeySpecConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOpenrouter, u.OfOpenAI, u.OfOpenAICompatible)
}
func (u *AIProviderKeySpecConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[AIProviderKeySpecConfigUnionParam](
		"type",
		apijson.Discriminator[AIProviderConfigOpenrouterParam]("openrouter"),
		apijson.Discriminator[AIProviderConfigOpenAIParam]("openai"),
		apijson.Discriminator[AIProviderConfigOpenAICompatibleParam]("openaiCompatible"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIProviderKeySpecCredentialsUnionParam struct {
	OfAPIKey  *AIProviderCredentialAPIKeyParam  `json:",omitzero,inline"`
	OfHeaders *AIProviderCredentialHeadersParam `json:",omitzero,inline"`
	paramUnion
}

func (u AIProviderKeySpecCredentialsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAPIKey, u.OfHeaders)
}
func (u *AIProviderKeySpecCredentialsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[AIProviderKeySpecCredentialsUnionParam](
		"type",
		apijson.Discriminator[AIProviderCredentialAPIKeyParam]("apiKey"),
		apijson.Discriminator[AIProviderCredentialHeadersParam]("headers"),
	)
}

type AIProviderKeyNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	Spec     AIProviderKeySpecParam             `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r AIProviderKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderKeyGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AIProviderKeyUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Fields to update.
	UpdateMask param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	Spec     AIProviderKeySpecParam             `json:"spec,omitzero"`
	paramObj
}

func (r AIProviderKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIProviderKeyListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, populate each item's info (model counts), at the cost of extra
	// lookups.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter expression (query param: prefix)
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// When true, return only promotional keys (provided by Cadenya, e.g. for
	// onboarding). Defaults to returning all keys, customer-provided and promotional
	// alike.
	Promotional param.Opt[bool] `query:"promotional,omitzero" json:"-"`
	// Free-form search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIProviderKeyListParams]'s query parameters as
// `url.Values`.
func (r AIProviderKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIProviderKeyDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
