// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cadenya/cadenya-go/internal/apiquery"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/packages/pagination"
)

// AgentService manages AI agents at the WORKSPACE level. Agents are
// workspace-scoped resources that define AI behavior and tool access. All
// operations are implicitly scoped to the workspace determined by the JWT token.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// AgentFeedbackService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentFeedbackService] method instead.
type AgentFeedbackService struct {
	Options []option.RequestOption
}

// NewAgentFeedbackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentFeedbackService(opts ...option.RequestOption) (r *AgentFeedbackService) {
	r = &AgentFeedbackService{}
	r.Options = opts
	return
}

// Lists feedback submitted across all objectives belonging to an agent. Supports
// search by comment, sentiment filter, agent variation filter, and creation date
// range. Results are ordered by creation time, newest first.
func (r *AgentFeedbackService) List(ctx context.Context, workspaceID string, agentID string, query AgentFeedbackListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveFeedback], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/feedback", workspaceID, agentID)
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

// Lists feedback submitted across all objectives belonging to an agent. Supports
// search by comment, sentiment filter, agent variation filter, and creation date
// range. Results are ordered by creation time, newest first.
func (r *AgentFeedbackService) ListAutoPaging(ctx context.Context, workspaceID string, agentID string, query AgentFeedbackListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveFeedback] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, agentID, query, opts...))
}

type AgentFeedbackListParams struct {
	// Optional filter to limit results to feedback on objectives run by a single agent
	// variation. Supports "external_id:" prefix for external IDs.
	AgentVariationID param.Field[string] `query:"agentVariationId"`
	// Inclusive lower bound on feedback creation time.
	CreatedAfter param.Field[time.Time] `query:"createdAfter" format:"date-time"`
	// Exclusive upper bound on feedback creation time.
	CreatedBefore param.Field[time.Time] `query:"createdBefore" format:"date-time"`
	// Pagination cursor from previous response.
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return.
	Limit param.Field[int64] `query:"limit"`
	// Free-text search applied to the feedback comment. Case-insensitive substring
	// match.
	Query param.Field[string] `query:"query"`
	// Filter by sentiment. UNSPECIFIED returns feedback regardless of score.
	Sentiment param.Field[AgentFeedbackListParamsSentiment] `query:"sentiment"`
}

// URLQuery serializes [AgentFeedbackListParams]'s query parameters as
// `url.Values`.
func (r AgentFeedbackListParams) URLQuery() (v url.Values) {
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

func (r AgentFeedbackListParamsSentiment) IsKnown() bool {
	switch r {
	case AgentFeedbackListParamsSentimentFeedbackSentimentUnspecified, AgentFeedbackListParamsSentimentFeedbackSentimentPositive, AgentFeedbackListParamsSentimentFeedbackSentimentNegative:
		return true
	}
	return false
}
