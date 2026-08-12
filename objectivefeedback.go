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

// ObjectiveFeedbackService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveFeedbackService] method instead.
type ObjectiveFeedbackService struct {
	options []option.RequestOption
}

// NewObjectiveFeedbackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectiveFeedbackService(opts ...option.RequestOption) (r ObjectiveFeedbackService) {
	r = ObjectiveFeedbackService{}
	r.options = opts
	return
}

// Submits feedback for an objective's execution. Feedback scores are used by the
// agent variation scoring system to evaluate and rank variation performance.
func (r *ObjectiveFeedbackService) New(ctx context.Context, objectiveID string, params ObjectiveFeedbackNewParams, opts ...option.RequestOption) (res *ObjectiveFeedback, err error) {
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
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/feedback", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Lists all feedback submitted for an objective
func (r *ObjectiveFeedbackService) List(ctx context.Context, objectiveID string, params ObjectiveFeedbackListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveFeedback], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/feedback", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
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

// Lists all feedback submitted for an objective
func (r *ObjectiveFeedbackService) ListAutoPaging(ctx context.Context, objectiveID string, params ObjectiveFeedbackListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveFeedback] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, objectiveID, params, opts...))
}

// ObjectiveFeedback represents feedback submitted for an objective's execution.
// Feedback is used to score agent variations and improve agent performance over
// time.
type ObjectiveFeedback struct {
	Data ObjectiveFeedbackData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	Info     ObjectiveFeedbackInfo    `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Metadata    respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveFeedback) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveFeedback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveFeedbackData struct {
	// Optional human-readable comment explaining the feedback
	Comment string `json:"comment"`
	// A score between -1.0 and 1.0 representing the quality of the objective's
	// execution. -1.0 is the worst possible score, 0.0 is neutral, and 1.0 is the
	// best.
	Score float64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comment     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveFeedbackData) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveFeedbackData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ObjectiveFeedbackData to a ObjectiveFeedbackDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ObjectiveFeedbackDataParam.Overrides()
func (r ObjectiveFeedbackData) ToParam() ObjectiveFeedbackDataParam {
	return param.Override[ObjectiveFeedbackDataParam](json.RawMessage(r.RawJSON()))
}

type ObjectiveFeedbackDataParam struct {
	// Optional human-readable comment explaining the feedback
	Comment param.Opt[string] `json:"comment,omitzero"`
	// A score between -1.0 and 1.0 representing the quality of the objective's
	// execution. -1.0 is the worst possible score, 0.0 is neutral, and 1.0 is the
	// best.
	Score param.Opt[float64] `json:"score,omitzero"`
	paramObj
}

func (r ObjectiveFeedbackDataParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveFeedbackDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveFeedbackDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveFeedbackInfo struct {
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	AgentVariation shared.BareMetadata `json:"agentVariation"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Objective shared.BareMetadata `json:"objective"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	SubmittedBy Profile `json:"submittedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentVariation respjson.Field
		Objective      respjson.Field
		SubmittedBy    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveFeedbackInfo) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveFeedbackInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveFeedbackNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string]          `path:"workspaceId,omitzero" api:"required" json:"-"`
	Data        ObjectiveFeedbackDataParam `json:"data,omitzero" api:"required"`
	// CreateOperationMetadata contains the user-provided fields for creating an
	// operation. Read-only fields (id, account_id, workspace_id, created_at,
	// profile_id) are excluded since they are set by the server.
	Metadata shared.CreateOperationMetadataParam `json:"metadata,omitzero" api:"required"`
	paramObj
}

func (r ObjectiveFeedbackNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveFeedbackNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveFeedbackNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveFeedbackListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectiveFeedbackListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveFeedbackListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
