package cadenya

import (
	"errors"
	"go.cadenya.com/cadenya-go/internal/apijson"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/respjson"
	"go.cadenya.com/cadenya-go/shared"
	"net/http"
	"slices"
	"time"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

// WebhookService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.options = opts
	return
}

func (r *WebhookService) UnsafeUnwrap(payload []byte, opts ...option.RequestOption) (*UnsafeUnwrapWebhookEvent, error) {
	res := &UnsafeUnwrapWebhookEvent{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEvent, error) {
	opts = slices.Concat(r.options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookKey
	if key == "" {
		return nil, errors.New("The WebhookKey option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &UnwrapWebhookEvent{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// The envelope for an objective event webhook delivery. Contains timestamp, event
// type, and the webhook data payload.
type UnsafeUnwrapWebhookEvent struct {
	// The webhook data payload with flat top-level keys for agent, variation,
	// objective, and event.
	Data      UnsafeUnwrapWebhookEventData `json:"data" api:"required"`
	Timestamp time.Time                    `json:"timestamp" api:"required" format:"date-time"`
	// The event type, prefixed with objective_event. (e.g.,
	// objective_event.tool_result)
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnsafeUnwrapWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *UnsafeUnwrapWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook data payload with flat top-level keys for agent, variation,
// objective, and event.
type UnsafeUnwrapWebhookEventData struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent" api:"required"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	AgentVariation shared.ResourceMetadata `json:"agentVariation" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective      shared.OperationMetadata `json:"objective" api:"required"`
	ObjectiveEvent ObjectiveEvent           `json:"objectiveEvent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent          respjson.Field
		AgentVariation respjson.Field
		Objective      respjson.Field
		ObjectiveEvent respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnsafeUnwrapWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *UnsafeUnwrapWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The envelope for an objective event webhook delivery. Contains timestamp, event
// type, and the webhook data payload.
type UnwrapWebhookEvent struct {
	// The webhook data payload with flat top-level keys for agent, variation,
	// objective, and event.
	Data      UnwrapWebhookEventData `json:"data" api:"required"`
	Timestamp time.Time              `json:"timestamp" api:"required" format:"date-time"`
	// The event type, prefixed with objective_event. (e.g.,
	// objective_event.tool_result)
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook data payload with flat top-level keys for agent, variation,
// objective, and event.
type UnwrapWebhookEventData struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent" api:"required"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	AgentVariation shared.ResourceMetadata `json:"agentVariation" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective      shared.OperationMetadata `json:"objective" api:"required"`
	ObjectiveEvent ObjectiveEvent           `json:"objectiveEvent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent          respjson.Field
		AgentVariation respjson.Field
		Objective      respjson.Field
		ObjectiveEvent respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnwrapWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *UnwrapWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
