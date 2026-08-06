// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
	Config AIProviderKeySpecConfig `json:"config"`
	// AIProviderCredential is the secret material used to authenticate with a
	// provider. The set case must correspond to AIProviderKeySpec.provider. The server
	// encrypts the serialized message at rest and never returns it on reads.
	Credentials AIProviderKeySpecCredentials `json:"credentials"`
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
	// The JSON name of the variant set in `config` (e.g. "openrouter"). Required from
	// clients on writes, filled by the server on reads; drives the discriminated union
	// in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenAI           respjson.Field
		OpenAICompatible respjson.Field
		Openrouter       respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKeySpecConfig) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpecConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIConfig holds OpenAI-specific settings.
type AIProviderKeySpecConfigOpenAI struct {
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
func (r AIProviderKeySpecConfigOpenAI) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpecConfigOpenAI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
// Completions API. The base URL is required and its model catalog is discovered
// live via GET {base_url}/models.
type AIProviderKeySpecConfigOpenAICompatible struct {
	BaseURL string `json:"baseUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKeySpecConfigOpenAICompatible) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpecConfigOpenAICompatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenRouterConfig holds OpenRouter-specific settings.
type AIProviderKeySpecConfigOpenrouter struct {
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
func (r AIProviderKeySpecConfigOpenrouter) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpecConfigOpenrouter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// The JSON name of the variant set in `credential` (e.g. "apiKey"). Required on
	// input; never returned (the credential is write-only). Drives the discriminated
	// union in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		Headers     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKeySpecCredentials) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpecCredentials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialAPIKey carries a single bearer/header API key.
type AIProviderKeySpecCredentialsAPIKey struct {
	APIKey string `json:"apiKey"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKeySpecCredentialsAPIKey) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpecCredentialsAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
type AIProviderKeySpecCredentialsHeaders struct {
	Headers map[string]string `json:"headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIProviderKeySpecCredentialsHeaders) RawJSON() string { return r.JSON.raw }
func (r *AIProviderKeySpecCredentialsHeaders) UnmarshalJSON(data []byte) error {
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
	Config AIProviderKeySpecConfigParam `json:"config,omitzero"`
	// AIProviderCredential is the secret material used to authenticate with a
	// provider. The set case must correspond to AIProviderKeySpec.provider. The server
	// encrypts the serialized message at rest and never returns it on reads.
	Credentials AIProviderKeySpecCredentialsParam `json:"credentials,omitzero"`
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

// AIProviderConfig holds non-secret, provider-specific settings. The set case must
// correspond to AIProviderKeySpec.provider. Providers with no settings (Anthropic,
// Gemini) simply leave this unset. The endpoint of a named provider is fixed and
// intentionally not overridable here; use the OpenAI-compatible provider to target
// a custom endpoint.
type AIProviderKeySpecConfigParam struct {
	// The JSON name of the variant set in `config` (e.g. "openrouter"). Required from
	// clients on writes, filled by the server on reads; drives the discriminated union
	// in the generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	// OpenAIConfig holds OpenAI-specific settings.
	OpenAI AIProviderKeySpecConfigOpenAIParam `json:"openai,omitzero"`
	// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
	// Completions API. The base URL is required and its model catalog is discovered
	// live via GET {base_url}/models.
	OpenAICompatible AIProviderKeySpecConfigOpenAICompatibleParam `json:"openaiCompatible,omitzero"`
	// OpenRouterConfig holds OpenRouter-specific settings.
	Openrouter AIProviderKeySpecConfigOpenrouterParam `json:"openrouter,omitzero"`
	paramObj
}

func (r AIProviderKeySpecConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAIConfig holds OpenAI-specific settings.
type AIProviderKeySpecConfigOpenAIParam struct {
	// Sent as the OpenAI-Organization header when set.
	OrganizationID param.Opt[string] `json:"organizationId,omitzero"`
	// Sent as the OpenAI-Project header when set.
	ProjectID param.Opt[string] `json:"projectId,omitzero"`
	paramObj
}

func (r AIProviderKeySpecConfigOpenAIParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecConfigOpenAIParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecConfigOpenAIParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenAICompatibleConfig configures a generic endpoint that speaks the OpenAI Chat
// Completions API. The base URL is required and its model catalog is discovered
// live via GET {base_url}/models.
//
// The property BaseURL is required.
type AIProviderKeySpecConfigOpenAICompatibleParam struct {
	BaseURL string `json:"baseUrl" api:"required"`
	paramObj
}

func (r AIProviderKeySpecConfigOpenAICompatibleParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecConfigOpenAICompatibleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecConfigOpenAICompatibleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OpenRouterConfig holds OpenRouter-specific settings.
type AIProviderKeySpecConfigOpenrouterParam struct {
	// Data-residency region (e.g. "us", "eu"). Empty uses the provider default.
	Region param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r AIProviderKeySpecConfigOpenrouterParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecConfigOpenrouterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecConfigOpenrouterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AIProviderCredential is the secret material used to authenticate with a
// provider. The set case must correspond to AIProviderKeySpec.provider. The server
// encrypts the serialized message at rest and never returns it on reads.
type AIProviderKeySpecCredentialsParam struct {
	// The JSON name of the variant set in `credential` (e.g. "apiKey"). Required on
	// input; never returned (the credential is write-only). Drives the discriminated
	// union in the generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	// CredentialAPIKey carries a single bearer/header API key.
	APIKey AIProviderKeySpecCredentialsAPIKeyParam `json:"apiKey,omitzero"`
	// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
	// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
	Headers AIProviderKeySpecCredentialsHeadersParam `json:"headers,omitzero"`
	paramObj
}

func (r AIProviderKeySpecCredentialsParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecCredentialsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecCredentialsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialAPIKey carries a single bearer/header API key.
type AIProviderKeySpecCredentialsAPIKeyParam struct {
	APIKey param.Opt[string] `json:"apiKey,omitzero"`
	paramObj
}

func (r AIProviderKeySpecCredentialsAPIKeyParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecCredentialsAPIKeyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecCredentialsAPIKeyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CredentialHeaders carries arbitrary HTTP headers sent with every request to the
// provider (e.g. {"Authorization": "Bearer ...", "X-Api-Key": "..."}).
type AIProviderKeySpecCredentialsHeadersParam struct {
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r AIProviderKeySpecCredentialsHeadersParam) MarshalJSON() (data []byte, err error) {
	type shadow AIProviderKeySpecCredentialsHeadersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIProviderKeySpecCredentialsHeadersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
