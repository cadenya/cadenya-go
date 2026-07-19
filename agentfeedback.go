// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"go.cadenya.com/cadenya-go/internal/apiquery"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/pagination"
	"go.cadenya.com/cadenya-go/packages/param"
	"net/http"
	"net/url"
	"slices"
	"time"
)

// Manage AI agents within a workspace. Agents define AI behavior and tool access.
//
// AgentFeedbackService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentFeedbackService] method instead.
type AgentFeedbackService struct {
	options []option.RequestOption
}

// NewAgentFeedbackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentFeedbackService(opts ...option.RequestOption) (r AgentFeedbackService) {
	r = AgentFeedbackService{}
	r.options = opts
	return
}

// Lists feedback submitted across all objectives belonging to an agent. Supports
// search by comment, sentiment filter, agent variation filter, and creation date
// range. Results are ordered by creation time, newest first.
func (r *AgentFeedbackService) List(ctx context.Context, agentID string, params AgentFeedbackListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveFeedback], err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/feedback", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID))
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

// Lists feedback submitted across all objectives belonging to an agent. Supports
// search by comment, sentiment filter, agent variation filter, and creation date
// range. Results are ordered by creation time, newest first.
func (r *AgentFeedbackService) ListAutoPaging(ctx context.Context, agentID string, params AgentFeedbackListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveFeedback] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, agentID, params, opts...))
}

type AgentFeedbackListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Optional filter to limit results to feedback on objectives run by a single agent
	// variation. Supports "external_id:" prefix for external IDs.
	AgentVariationID param.Opt[string] `query:"agentVariationId,omitzero" json:"-"`
	// Inclusive lower bound on feedback creation time.
	CreatedAfter param.Opt[time.Time] `query:"createdAfter,omitzero" format:"date-time" json:"-"`
	// Exclusive upper bound on feedback creation time.
	CreatedBefore param.Opt[time.Time] `query:"createdBefore,omitzero" format:"date-time" json:"-"`
	// Pagination cursor from previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search applied to the feedback comment. Case-insensitive substring
	// match.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Filter by sentiment. UNSPECIFIED returns feedback regardless of score.
	//
	// Any of "FEEDBACK_SENTIMENT_UNSPECIFIED", "FEEDBACK_SENTIMENT_POSITIVE",
	// "FEEDBACK_SENTIMENT_NEGATIVE".
	Sentiment AgentFeedbackListParamsSentiment `query:"sentiment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentFeedbackListParams]'s query parameters as
// `url.Values`.
func (r AgentFeedbackListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by sentiment. UNSPECIFIED returns feedback regardless of score.
type AgentFeedbackListParamsSentiment string

const (
	AgentFeedbackListParamsSentimentFeedbackSentimentUnspecified AgentFeedbackListParamsSentiment = "FEEDBACK_SENTIMENT_UNSPECIFIED"
	AgentFeedbackListParamsSentimentFeedbackSentimentPositive    AgentFeedbackListParamsSentiment = "FEEDBACK_SENTIMENT_POSITIVE"
	AgentFeedbackListParamsSentimentFeedbackSentimentNegative    AgentFeedbackListParamsSentiment = "FEEDBACK_SENTIMENT_NEGATIVE"
)
