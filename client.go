// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyago

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cadenya API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	// AccountService manages account-level operations. Accounts are the top-level
	// organizational unit in the system. All operations are scoped to the
	// authenticated account determined by the JWT token.
	//
	// Authentication: Bearer token (JWT) Scope: Account-level operations
	Account *AccountService
	// AgentService manages AI agents at the WORKSPACE level. Agents are
	// workspace-scoped resources that define AI behavior and tool access. All
	// operations are implicitly scoped to the workspace determined by the JWT token.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	Agents     *AgentService
	Objectives *ObjectiveService
	// ModelService manages LLM models at the WORKSPACE level. Models represent
	// available LLM providers and families (e.g., "anthropic/claude-sonnet-4.6").
	// Models are seeded into workspaces and can be enabled or disabled. All operations
	// are implicitly scoped to the workspace determined by the JWT token.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	Models *ModelService
	Search *SearchService
	// ToolService manages tool sets and tools at the WORKSPACE level. Tool sets group
	// related tools, and tools define specific capabilities for agents. All operations
	// are implicitly scoped to the workspace determined by the JWT token.
	//
	// Note: When a ToolSet has managed=true, only API Key actors can modify its tools.
	// Profile actors (humans) are restricted from modifying managed tool sets.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	ToolSets *ToolSetService
	// APIKeyService manages workspace-scoped API Keys. Each API key belongs to a
	// single workspace, ensuring isolation between environments.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	APIKeys          *APIKeyService
	WorkspaceSecrets *WorkspaceSecretService
	// WorkspaceService manages workspaces at the ACCOUNT level. This service is
	// responsible for creating and listing workspaces within an account. Workspaces
	// provide organizational grouping for resources within an account.
	//
	// Authentication: Bearer token (JWT) Scope: Account-level operations (manages
	// workspaces themselves, not resources within workspaces)
	Workspaces *WorkspaceService
}

// DefaultClientOptions read from the environment (CADENYA_API_KEY,
// CADENYA_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CADENYA_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("CADENYA_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (CADENYA_API_KEY, CADENYA_BASE_URL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = &Client{Options: opts}

	r.Account = NewAccountService(opts...)
	r.Agents = NewAgentService(opts...)
	r.Objectives = NewObjectiveService(opts...)
	r.Models = NewModelService(opts...)
	r.Search = NewSearchService(opts...)
	r.ToolSets = NewToolSetService(opts...)
	r.APIKeys = NewAPIKeyService(opts...)
	r.WorkspaceSecrets = NewWorkspaceSecretService(opts...)
	r.Workspaces = NewWorkspaceService(opts...)

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
