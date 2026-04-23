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

// ObjectiveFeedbackService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveFeedbackService] method instead.
type ObjectiveFeedbackService struct {
	Options []option.RequestOption
}

// NewObjectiveFeedbackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectiveFeedbackService(opts ...option.RequestOption) (r *ObjectiveFeedbackService) {
	r = &ObjectiveFeedbackService{}
	r.Options = opts
	return
}

// Submits feedback for an objective's execution. Feedback scores are used by the
// agent variation scoring system to evaluate and rank variation performance.
func (r *ObjectiveFeedbackService) New(ctx context.Context, objectiveID string, body ObjectiveFeedbackNewParams, opts ...option.RequestOption) (res *ObjectiveFeedback, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/feedback", objectiveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists all feedback submitted for an objective
func (r *ObjectiveFeedbackService) List(ctx context.Context, objectiveID string, query ObjectiveFeedbackListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveFeedback], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/feedback", objectiveID)
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

// Lists all feedback submitted for an objective
func (r *ObjectiveFeedbackService) ListAutoPaging(ctx context.Context, objectiveID string, query ObjectiveFeedbackListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveFeedback] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, objectiveID, query, opts...))
}

// ObjectiveFeedback represents feedback submitted for an objective's execution.
// Feedback is used to score agent variations and improve agent performance over
// time.
type ObjectiveFeedback struct {
	Data ObjectiveFeedbackData `json:"data" api:"required"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Metadata shared.BareMetadata   `json:"metadata" api:"required"`
	Info     ObjectiveFeedbackInfo `json:"info"`
	JSON     objectiveFeedbackJSON `json:"-"`
}

// objectiveFeedbackJSON contains the JSON metadata for the struct
// [ObjectiveFeedback]
type objectiveFeedbackJSON struct {
	Data        apijson.Field
	Metadata    apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveFeedback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveFeedbackJSON) RawJSON() string {
	return r.raw
}

type ObjectiveFeedbackData struct {
	// Arbitrary key-value pairs to identify the source of the feedback. Since the
	// submitting profile is typically an API key, use this to pass through
	// application-specific identifiers (e.g., {"user_id": "usr_123", "session_id":
	// "abc"}).
	Attributes map[string]string `json:"attributes"`
	// Optional human-readable comment explaining the feedback
	Comment string `json:"comment"`
	// A score between -1.0 and 1.0 representing the quality of the objective's
	// execution. -1.0 is the worst possible score, 0.0 is neutral, and 1.0 is the
	// best.
	Score float64                   `json:"score"`
	JSON  objectiveFeedbackDataJSON `json:"-"`
}

// objectiveFeedbackDataJSON contains the JSON metadata for the struct
// [ObjectiveFeedbackData]
type objectiveFeedbackDataJSON struct {
	Attributes  apijson.Field
	Comment     apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveFeedbackData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveFeedbackDataJSON) RawJSON() string {
	return r.raw
}

type ObjectiveFeedbackDataParam struct {
	// Arbitrary key-value pairs to identify the source of the feedback. Since the
	// submitting profile is typically an API key, use this to pass through
	// application-specific identifiers (e.g., {"user_id": "usr_123", "session_id":
	// "abc"}).
	Attributes param.Field[map[string]string] `json:"attributes"`
	// Optional human-readable comment explaining the feedback
	Comment param.Field[string] `json:"comment"`
	// A score between -1.0 and 1.0 representing the quality of the objective's
	// execution. -1.0 is the worst possible score, 0.0 is neutral, and 1.0 is the
	// best.
	Score param.Field[float64] `json:"score"`
}

func (r ObjectiveFeedbackDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	SubmittedBy Profile                   `json:"submittedBy"`
	JSON        objectiveFeedbackInfoJSON `json:"-"`
}

// objectiveFeedbackInfoJSON contains the JSON metadata for the struct
// [ObjectiveFeedbackInfo]
type objectiveFeedbackInfoJSON struct {
	AgentVariation apijson.Field
	Objective      apijson.Field
	SubmittedBy    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObjectiveFeedbackInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveFeedbackInfoJSON) RawJSON() string {
	return r.raw
}

type ObjectiveFeedbackNewParams struct {
	Data param.Field[ObjectiveFeedbackDataParam] `json:"data" api:"required"`
}

func (r ObjectiveFeedbackNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveFeedbackListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ObjectiveFeedbackListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveFeedbackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
