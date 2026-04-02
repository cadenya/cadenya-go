// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cadenya/cadenya-sdk-go/internal/apijson"
	"github.com/cadenya/cadenya-sdk-go/internal/apiquery"
	"github.com/cadenya/cadenya-sdk-go/internal/param"
	"github.com/cadenya/cadenya-sdk-go/internal/requestconfig"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/cadenya/cadenya-sdk-go/packages/pagination"
	"github.com/cadenya/cadenya-sdk-go/shared"
)

// ModelService manages LLM models at the WORKSPACE level. Models represent
// available LLM providers and families (e.g., "anthropic/claude-sonnet-4.6").
// Models are seeded into workspaces and can be enabled or disabled. All operations
// are implicitly scoped to the workspace determined by the JWT token.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// ModelService contains methods and other services that help with interacting with
// the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	return
}

// Retrieves a model by ID from the workspace
func (r *ModelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/models/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all models in the workspace
func (r *ModelService) List(ctx context.Context, query ModelListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Model], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/models"
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

// Lists all models in the workspace
func (r *ModelService) ListAutoPaging(ctx context.Context, query ModelListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Model] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Enables or disables a model in the workspace
func (r *ModelService) SetStatus(ctx context.Context, id string, body ModelSetStatusParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/models/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type Model struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// Model specification
	Spec ModelSpec `json:"spec" api:"required"`
	JSON modelJSON `json:"-"`
}

// modelJSON contains the JSON metadata for the struct [Model]
type modelJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Model) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelJSON) RawJSON() string {
	return r.raw
}

type ModelSpec struct {
	// The model family (e.g., "claude-sonnet-4.6", "gpt-5.4", "gemini-2.5-flash")
	Family string `json:"family"`
	// Cost per million input tokens in cents (e.g., 300 = $3.00)
	InputPricePerMillionTokens string `json:"inputPricePerMillionTokens"`
	// Maximum number of input tokens the model supports
	MaxInputTokens int64 `json:"maxInputTokens"`
	// Maximum number of output tokens the model can generate
	MaxOutputTokens int64 `json:"maxOutputTokens"`
	// Cost per million output tokens in cents (e.g., 1500 = $15.00)
	OutputPricePerMillionTokens string `json:"outputPricePerMillionTokens"`
	// The model provider (e.g., "anthropic", "openai", "google")
	Provider string `json:"provider"`
	// The status of the model in the workspace
	Status ModelSpecStatus `json:"status"`
	JSON   modelSpecJSON   `json:"-"`
}

// modelSpecJSON contains the JSON metadata for the struct [ModelSpec]
type modelSpecJSON struct {
	Family                      apijson.Field
	InputPricePerMillionTokens  apijson.Field
	MaxInputTokens              apijson.Field
	MaxOutputTokens             apijson.Field
	OutputPricePerMillionTokens apijson.Field
	Provider                    apijson.Field
	Status                      apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ModelSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelSpecJSON) RawJSON() string {
	return r.raw
}

// The status of the model in the workspace
type ModelSpecStatus string

const (
	ModelSpecStatusModelStatusUnspecified ModelSpecStatus = "MODEL_STATUS_UNSPECIFIED"
	ModelSpecStatusModelStatusEnabled     ModelSpecStatus = "MODEL_STATUS_ENABLED"
	ModelSpecStatusModelStatusDisabled    ModelSpecStatus = "MODEL_STATUS_DISABLED"
)

func (r ModelSpecStatus) IsKnown() bool {
	switch r {
	case ModelSpecStatusModelStatusUnspecified, ModelSpecStatusModelStatusEnabled, ModelSpecStatusModelStatusDisabled:
		return true
	}
	return false
}

type ModelListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter by name prefix
	Prefix param.Field[string] `query:"prefix"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by model status
	Status param.Field[ModelListParamsStatus] `query:"status"`
}

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by model status
type ModelListParamsStatus string

const (
	ModelListParamsStatusModelStatusUnspecified ModelListParamsStatus = "MODEL_STATUS_UNSPECIFIED"
	ModelListParamsStatusModelStatusEnabled     ModelListParamsStatus = "MODEL_STATUS_ENABLED"
	ModelListParamsStatusModelStatusDisabled    ModelListParamsStatus = "MODEL_STATUS_DISABLED"
)

func (r ModelListParamsStatus) IsKnown() bool {
	switch r {
	case ModelListParamsStatusModelStatusUnspecified, ModelListParamsStatusModelStatusEnabled, ModelListParamsStatusModelStatusDisabled:
		return true
	}
	return false
}

type ModelSetStatusParams struct {
	// The new status for the model
	Status param.Field[ModelSetStatusParamsStatus] `json:"status"`
}

func (r ModelSetStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The new status for the model
type ModelSetStatusParamsStatus string

const (
	ModelSetStatusParamsStatusModelStatusUnspecified ModelSetStatusParamsStatus = "MODEL_STATUS_UNSPECIFIED"
	ModelSetStatusParamsStatusModelStatusEnabled     ModelSetStatusParamsStatus = "MODEL_STATUS_ENABLED"
	ModelSetStatusParamsStatusModelStatusDisabled    ModelSetStatusParamsStatus = "MODEL_STATUS_DISABLED"
)

func (r ModelSetStatusParamsStatus) IsKnown() bool {
	switch r {
	case ModelSetStatusParamsStatusModelStatusUnspecified, ModelSetStatusParamsStatusModelStatusEnabled, ModelSetStatusParamsStatusModelStatusDisabled:
		return true
	}
	return false
}
