// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/cadenya/cadenya-go/internal/apijson"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/shared"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

// WebhookService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
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
	opts = slices.Concat(r.Options, opts)
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
	Type string                       `json:"type" api:"required"`
	JSON unsafeUnwrapWebhookEventJSON `json:"-"`
}

// unsafeUnwrapWebhookEventJSON contains the JSON metadata for the struct
// [UnsafeUnwrapWebhookEvent]
type unsafeUnwrapWebhookEventJSON struct {
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnsafeUnwrapWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unsafeUnwrapWebhookEventJSON) RawJSON() string {
	return r.raw
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
	Objective      shared.OperationMetadata                   `json:"objective" api:"required"`
	ObjectiveEvent UnsafeUnwrapWebhookEventDataObjectiveEvent `json:"objectiveEvent" api:"required"`
	JSON           unsafeUnwrapWebhookEventDataJSON           `json:"-"`
}

// unsafeUnwrapWebhookEventDataJSON contains the JSON metadata for the struct
// [UnsafeUnwrapWebhookEventData]
type unsafeUnwrapWebhookEventDataJSON struct {
	Agent          apijson.Field
	AgentVariation apijson.Field
	Objective      apijson.Field
	ObjectiveEvent apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UnsafeUnwrapWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unsafeUnwrapWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type UnsafeUnwrapWebhookEventDataObjectiveEvent struct {
	Data ObjectiveEventData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata        shared.OperationMetadata                       `json:"metadata" api:"required"`
	ContextWindowID string                                         `json:"contextWindowId"`
	Info            ObjectiveEventInfo                             `json:"info"`
	JSON            unsafeUnwrapWebhookEventDataObjectiveEventJSON `json:"-"`
}

// unsafeUnwrapWebhookEventDataObjectiveEventJSON contains the JSON metadata for
// the struct [UnsafeUnwrapWebhookEventDataObjectiveEvent]
type unsafeUnwrapWebhookEventDataObjectiveEventJSON struct {
	Data            apijson.Field
	Metadata        apijson.Field
	ContextWindowID apijson.Field
	Info            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UnsafeUnwrapWebhookEventDataObjectiveEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unsafeUnwrapWebhookEventDataObjectiveEventJSON) RawJSON() string {
	return r.raw
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
	Type string                 `json:"type" api:"required"`
	JSON unwrapWebhookEventJSON `json:"-"`
}

// unwrapWebhookEventJSON contains the JSON metadata for the struct
// [UnwrapWebhookEvent]
type unwrapWebhookEventJSON struct {
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnwrapWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unwrapWebhookEventJSON) RawJSON() string {
	return r.raw
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
	Objective      shared.OperationMetadata             `json:"objective" api:"required"`
	ObjectiveEvent UnwrapWebhookEventDataObjectiveEvent `json:"objectiveEvent" api:"required"`
	JSON           unwrapWebhookEventDataJSON           `json:"-"`
}

// unwrapWebhookEventDataJSON contains the JSON metadata for the struct
// [UnwrapWebhookEventData]
type unwrapWebhookEventDataJSON struct {
	Agent          apijson.Field
	AgentVariation apijson.Field
	Objective      apijson.Field
	ObjectiveEvent apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UnwrapWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unwrapWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

type UnwrapWebhookEventDataObjectiveEvent struct {
	Data ObjectiveEventData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata        shared.OperationMetadata                 `json:"metadata" api:"required"`
	ContextWindowID string                                   `json:"contextWindowId"`
	Info            ObjectiveEventInfo                       `json:"info"`
	JSON            unwrapWebhookEventDataObjectiveEventJSON `json:"-"`
}

// unwrapWebhookEventDataObjectiveEventJSON contains the JSON metadata for the
// struct [UnwrapWebhookEventDataObjectiveEvent]
type unwrapWebhookEventDataObjectiveEventJSON struct {
	Data            apijson.Field
	Metadata        apijson.Field
	ContextWindowID apijson.Field
	Info            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UnwrapWebhookEventDataObjectiveEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unwrapWebhookEventDataObjectiveEventJSON) RawJSON() string {
	return r.raw
}
