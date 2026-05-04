// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cadenya API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	// Manage the authenticated account. Accounts are the top-level organizational unit
	// and contain one or more workspaces.
	Account *AccountService
	// Manage AI agents within a workspace. Agents define AI behavior and tool access.
	Agents     *AgentService
	Objectives *ObjectiveService
	// Manage memory layers and their entries. Layers are named containers that can be
	// composed into an objective's memory stack; entries are the keyed values within a
	// layer. System-managed layers (e.g., episodic layers created by the runtime)
	// cannot be mutated through this API.
	MemoryLayers *MemoryLayerService
	// Issue short-lived presigned URLs for direct client-to-object-storage uploads.
	// Created uploads can be referenced by id when creating or updating resources that
	// accept binary content (e.g., MemoryEntry).
	Uploads *UploadService
	// Manage LLM models available to a workspace. Models represent provider and family
	// pairs (e.g., "anthropic/claude-sonnet-4.6"). Workspaces are seeded with the
	// supported models and you can enable or disable each one.
	Models *ModelService
	Search *SearchService
	// Manage tool sets and the tools they contain. Tool sets group related tools, and
	// tools define specific capabilities available to agents.
	//
	// When a tool set is managed, only API key actors can modify its tools; human
	// (profile) actors cannot.
	ToolSets *ToolSetService
	// Issue, rotate, and revoke API keys for a workspace. Each API key belongs to
	// exactly one workspace, ensuring isolation between environments.
	APIKeys          *APIKeyService
	WorkspaceSecrets *WorkspaceSecretService
	// Manage workspaces within an account. Workspaces provide organizational grouping
	// and isolation for resources such as agents, tools, and API keys.
	Workspaces *WorkspaceService
	Webhooks   *WebhookService
	// Apply a declarative bundle of workspace resources — tool sets, memory layers,
	// agents, variations, assignments, and schedules — in a single asynchronous
	// operation.
	BulkWorkspaceResources *BulkWorkspaceResourceService
}

// DefaultClientOptions read from the environment (CADENYA_API_KEY,
// CADENYA_WEBHOOK_KEY, CADENYA_BASE_URL). This should be used to initialize new
// clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CADENYA_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("CADENYA_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("CADENYA_WEBHOOK_KEY"); ok {
		defaults = append(defaults, option.WithWebhookKey(o))
	}
	if o, ok := os.LookupEnv("CADENYA_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (CADENYA_API_KEY, CADENYA_WEBHOOK_KEY, CADENYA_BASE_URL). The option
// passed in as arguments are applied after these default arguments, and all option
// will be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = &Client{Options: opts}

	r.Account = NewAccountService(opts...)
	r.Agents = NewAgentService(opts...)
	r.Objectives = NewObjectiveService(opts...)
	r.MemoryLayers = NewMemoryLayerService(opts...)
	r.Uploads = NewUploadService(opts...)
	r.Models = NewModelService(opts...)
	r.Search = NewSearchService(opts...)
	r.ToolSets = NewToolSetService(opts...)
	r.APIKeys = NewAPIKeyService(opts...)
	r.WorkspaceSecrets = NewWorkspaceSecretService(opts...)
	r.Workspaces = NewWorkspaceService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.BulkWorkspaceResources = NewBulkWorkspaceResourceService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
