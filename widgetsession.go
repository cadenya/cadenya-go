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
	"time"
)

// Mint and manage widget sessions. Session creation is server-to-server only: the
// customer's backend authenticates its visitor, asserts tenant/subject context,
// attaches any per-visitor secrets, and receives a short-lived bearer token the
// browser uses against the widget host.
//
// WidgetSessionService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWidgetSessionService] method instead.
type WidgetSessionService struct {
	options []option.RequestOption
}

// NewWidgetSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWidgetSessionService(opts ...option.RequestOption) (r WidgetSessionService) {
	r = WidgetSessionService{}
	r.options = opts
	return
}

// Mints a session against a widget and returns the session bearer token
// (`spec.token`, returned only on creation) plus the authoritative widget hostname
// (`info.host`). Asserting a tenant upserts the tenant record; attached secrets
// flow to every conversation the session creates.
func (r *WidgetSessionService) New(ctx context.Context, params WidgetSessionNewParams, opts ...option.RequestOption) (res *WidgetSession, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/widget_sessions", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a widget session. The bearer token is never returned on reads.
func (r *WidgetSessionService) Get(ctx context.Context, id string, query WidgetSessionGetParams, opts ...option.RequestOption) (res *WidgetSession, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/widget_sessions/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists widget sessions in a workspace, filterable by widget, tenant, subject, and
// state
func (r *WidgetSessionService) List(ctx context.Context, params WidgetSessionListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[WidgetSession], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/widget_sessions", url.PathEscape(params.WorkspaceID.Value))
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

// Lists widget sessions in a workspace, filterable by widget, tenant, subject, and
// state
func (r *WidgetSessionService) ListAutoPaging(ctx context.Context, params WidgetSessionListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[WidgetSession] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a session and its secrets. The session's conversations are
// disassociated, not deleted; use the tenant-level delete for full erasure.
func (r *WidgetSessionService) Delete(ctx context.Context, id string, body WidgetSessionDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/widget_sessions/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Deletes every session belonging to a tenant across all widgets in the workspace,
// along with the conversations those sessions created — built for GDPR erasure
// requests. The tenant is required; an empty value is rejected rather than
// matching everything.
func (r *WidgetSessionService) DeleteTenant(ctx context.Context, params WidgetSessionDeleteTenantParams, opts ...option.RequestOption) (res *WidgetSessionDeleteTenantResponse, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/widget_sessions", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Transitions a session to STATE_REVOKED. Outstanding tokens stop working
// immediately, open event streams close within seconds, and the session's secrets
// are deleted. Terminal.
func (r *WidgetSessionService) Revoke(ctx context.Context, id string, body WidgetSessionRevokeParams, opts ...option.RequestOption) (res *WidgetSession, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.WorkspaceID, precfg.WorkspaceID)
	if body.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/widget_sessions/%s:revoke", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// SubjectAssertion identifies a person within a tenant in the customer's own
// namespace — typically their user id. Asserting a subject upserts the subject
// record under the asserted tenant and associates the created resource with it. A
// subject assertion is only valid alongside a tenant assertion: subject
// identifiers are scoped to their tenant.
type SubjectAssertion struct {
	// The subject identifier in the customer's namespace (e.g. their user id). Stored
	// as the subject record's external_id; unique within the tenant.
	ID string `json:"id" api:"required"`
	// Optional human-readable name for the subject. Updates the subject record's name
	// on every assertion that provides it.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubjectAssertion) RawJSON() string { return r.JSON.raw }
func (r *SubjectAssertion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SubjectAssertion to a SubjectAssertionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SubjectAssertionParam.Overrides()
func (r SubjectAssertion) ToParam() SubjectAssertionParam {
	return param.Override[SubjectAssertionParam](json.RawMessage(r.RawJSON()))
}

// SubjectAssertion identifies a person within a tenant in the customer's own
// namespace — typically their user id. Asserting a subject upserts the subject
// record under the asserted tenant and associates the created resource with it. A
// subject assertion is only valid alongside a tenant assertion: subject
// identifiers are scoped to their tenant.
//
// The property ID is required.
type SubjectAssertionParam struct {
	// The subject identifier in the customer's namespace (e.g. their user id). Stored
	// as the subject record's external_id; unique within the tenant.
	ID string `json:"id" api:"required"`
	// Optional human-readable name for the subject. Updates the subject record's name
	// on every assertion that provides it.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r SubjectAssertionParam) MarshalJSON() (data []byte, err error) {
	type shadow SubjectAssertionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubjectAssertionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SubjectReference is the read-only echo of a resource's subject association,
// carrying both Cadenya's canonical id and the customer's own key.
type SubjectReference struct {
	// Cadenya's canonical subject id.
	ID string `json:"id" api:"required"`
	// The subject identifier in the customer's namespace, as asserted. Unique within
	// the subject's tenant.
	ExternalID string `json:"externalId" api:"required"`
	// Human-readable name of the subject, when one has been asserted.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubjectReference) RawJSON() string { return r.JSON.raw }
func (r *SubjectReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TenantAssertion identifies a tenant in the customer's own namespace — their org,
// company, or team identifier for an end user. Asserting a tenant upserts the
// tenant record in the workspace (keyed on `id` as the tenant's external_id) and
// associates the created resource with it.
type TenantAssertion struct {
	// The tenant identifier in the customer's namespace (e.g. "acme-corp"). Stored as
	// the tenant record's external_id; stable across requests.
	ID string `json:"id" api:"required"`
	// Optional human-readable name for the tenant. Updates the tenant record's name on
	// every assertion that provides it.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantAssertion) RawJSON() string { return r.JSON.raw }
func (r *TenantAssertion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TenantAssertion to a TenantAssertionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TenantAssertionParam.Overrides()
func (r TenantAssertion) ToParam() TenantAssertionParam {
	return param.Override[TenantAssertionParam](json.RawMessage(r.RawJSON()))
}

// TenantAssertion identifies a tenant in the customer's own namespace — their org,
// company, or team identifier for an end user. Asserting a tenant upserts the
// tenant record in the workspace (keyed on `id` as the tenant's external_id) and
// associates the created resource with it.
//
// The property ID is required.
type TenantAssertionParam struct {
	// The tenant identifier in the customer's namespace (e.g. "acme-corp"). Stored as
	// the tenant record's external_id; stable across requests.
	ID string `json:"id" api:"required"`
	// Optional human-readable name for the tenant. Updates the tenant record's name on
	// every assertion that provides it.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r TenantAssertionParam) MarshalJSON() (data []byte, err error) {
	type shadow TenantAssertionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenantAssertionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TenantReference is the read-only echo of a resource's tenant association,
// carrying both Cadenya's canonical id and the customer's own key.
type TenantReference struct {
	// Cadenya's canonical tenant id.
	ID string `json:"id" api:"required"`
	// The tenant identifier in the customer's namespace, as asserted.
	ExternalID string `json:"externalId" api:"required"`
	// Human-readable name of the tenant, when one has been asserted.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExternalID  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantReference) RawJSON() string { return r.JSON.raw }
func (r *TenantReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WidgetSession is a delegated, narrowed credential for one visitor's use of a
// widget, minted server-to-server by the customer's backend. The session carries
// all customer-asserted context — tenant, subject, labels, secrets — and every
// conversation (objective) created through the widget inherits it. The bearer
// token returned at mint is short-lived and refreshed at the widget host; the
// session row is what makes revocation possible.
type WidgetSession struct {
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// WidgetSessionSpec is the configuration of a session, fixed at mint.
	Spec WidgetSessionSpec `json:"spec" api:"required"`
	// The current lifecycle state of the session. Output only. Sessions are created
	// STATE_ACTIVE; use :revoke to end one early.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_EXPIRED", "STATE_REVOKED",
	// "STATE_EXHAUSTED".
	State WidgetSessionState `json:"state" api:"required"`
	// WidgetSessionInfo provides read-only server-derived data about a session.
	Info WidgetSessionInfo `json:"info"`
	// Names of the secrets attached to the session. Values are write-only: provided at
	// creation, encrypted at rest, and interpolated into tool-call headers server-side
	// — never returned by any API.
	Secrets []WidgetSessionSecret `json:"secrets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		State       respjson.Field
		Info        respjson.Field
		Secrets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WidgetSession) RawJSON() string { return r.JSON.raw }
func (r *WidgetSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current lifecycle state of the session. Output only. Sessions are created
// STATE_ACTIVE; use :revoke to end one early.
type WidgetSessionState string

const (
	WidgetSessionStateStateUnspecified WidgetSessionState = "STATE_UNSPECIFIED"
	WidgetSessionStateStateActive      WidgetSessionState = "STATE_ACTIVE"
	WidgetSessionStateStateExpired     WidgetSessionState = "STATE_EXPIRED"
	WidgetSessionStateStateRevoked     WidgetSessionState = "STATE_REVOKED"
	WidgetSessionStateStateExhausted   WidgetSessionState = "STATE_EXHAUSTED"
)

// Secret is the name-only echo of a secret attached to the session. Values are
// never returned.
type WidgetSessionSecret struct {
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WidgetSessionSecret) RawJSON() string { return r.JSON.raw }
func (r *WidgetSessionSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WidgetSessionInfo provides read-only server-derived data about a session.
type WidgetSessionInfo struct {
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Agent shared.BareMetadata `json:"agent"`
	// The widget hostname this session's tokens are bound to. Authoritative — clients
	// must use this value rather than constructing the hostname.
	Host string `json:"host"`
	// When the session last created a conversation, sent a message, or refreshed a
	// token.
	LastActiveAt time.Time `json:"lastActiveAt" format:"date-time"`
	// Number of conversation messages created through this session, counted against
	// the session's message cap.
	MessageCount int64 `json:"messageCount"`
	// SubjectReference is the read-only echo of a resource's subject association,
	// carrying both Cadenya's canonical id and the customer's own key.
	Subject SubjectReference `json:"subject"`
	// TenantReference is the read-only echo of a resource's tenant association,
	// carrying both Cadenya's canonical id and the customer's own key.
	Tenant TenantReference `json:"tenant"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Widget shared.BareMetadata `json:"widget"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent        respjson.Field
		Host         respjson.Field
		LastActiveAt respjson.Field
		MessageCount respjson.Field
		Subject      respjson.Field
		Tenant       respjson.Field
		Widget       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WidgetSessionInfo) RawJSON() string { return r.JSON.raw }
func (r *WidgetSessionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WidgetSessionSpec is the configuration of a session, fixed at mint.
type WidgetSessionSpec struct {
	// Widget this session is minted against. Accepts the canonical `wgt_…` form or the
	// `external_id:<value>` form.
	WidgetID string `json:"widgetId" api:"required"`
	// The session bearer token. Returned only on creation — subsequent reads omit it.
	// The token is short-lived; the widget refreshes it at the widget host without
	// involving the customer's backend.
	Token string `json:"token"`
	// Hard session expiry. Tokens never outlive it; after it passes the session
	// transitions to STATE_EXPIRED. Defaults to a server-chosen horizon when unset.
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// Parameters forced onto tool calls made by this session's conversations. A pinned
	// parameter is an overlay on a tool's JSON schema: the parameter is removed from
	// what the LLM sees, and its value is always overwritten server-side with the
	// pinned value — so the model cannot be tricked into calling a tool with a
	// different id than the one the session was minted for (e.g. pin "workspaceId" for
	// an OpenAPI tool with a /workspaces/{workspaceId} path). Flows to every objective
	// the session creates.
	PinnedParameters map[string]string `json:"pinnedParameters"`
	// SubjectAssertion identifies a person within a tenant in the customer's own
	// namespace — typically their user id. Asserting a subject upserts the subject
	// record under the asserted tenant and associates the created resource with it. A
	// subject assertion is only valid alongside a tenant assertion: subject
	// identifiers are scoped to their tenant.
	Subject SubjectAssertion `json:"subject"`
	// TenantAssertion identifies a tenant in the customer's own namespace — their org,
	// company, or team identifier for an end user. Asserting a tenant upserts the
	// tenant record in the workspace (keyed on `id` as the tenant's external_id) and
	// associates the created resource with it.
	Tenant TenantAssertion `json:"tenant"`
	// Expiry of the token returned in `token`. Distinct from `expires_at`, which
	// bounds the session itself.
	TokenExpiresAt time.Time `json:"tokenExpiresAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WidgetID         respjson.Field
		Token            respjson.Field
		ExpiresAt        respjson.Field
		PinnedParameters respjson.Field
		Subject          respjson.Field
		Tenant           respjson.Field
		TokenExpiresAt   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WidgetSessionSpec) RawJSON() string { return r.JSON.raw }
func (r *WidgetSessionSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WidgetSessionSpec to a WidgetSessionSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WidgetSessionSpecParam.Overrides()
func (r WidgetSessionSpec) ToParam() WidgetSessionSpecParam {
	return param.Override[WidgetSessionSpecParam](json.RawMessage(r.RawJSON()))
}

// WidgetSessionSpec is the configuration of a session, fixed at mint.
//
// The property WidgetID is required.
type WidgetSessionSpecParam struct {
	// Widget this session is minted against. Accepts the canonical `wgt_…` form or the
	// `external_id:<value>` form.
	WidgetID string `json:"widgetId" api:"required"`
	// Hard session expiry. Tokens never outlive it; after it passes the session
	// transitions to STATE_EXPIRED. Defaults to a server-chosen horizon when unset.
	ExpiresAt param.Opt[time.Time] `json:"expiresAt,omitzero" format:"date-time"`
	// Parameters forced onto tool calls made by this session's conversations. A pinned
	// parameter is an overlay on a tool's JSON schema: the parameter is removed from
	// what the LLM sees, and its value is always overwritten server-side with the
	// pinned value — so the model cannot be tricked into calling a tool with a
	// different id than the one the session was minted for (e.g. pin "workspaceId" for
	// an OpenAPI tool with a /workspaces/{workspaceId} path). Flows to every objective
	// the session creates.
	PinnedParameters map[string]string `json:"pinnedParameters,omitzero"`
	// SubjectAssertion identifies a person within a tenant in the customer's own
	// namespace — typically their user id. Asserting a subject upserts the subject
	// record under the asserted tenant and associates the created resource with it. A
	// subject assertion is only valid alongside a tenant assertion: subject
	// identifiers are scoped to their tenant.
	Subject SubjectAssertionParam `json:"subject,omitzero"`
	// TenantAssertion identifies a tenant in the customer's own namespace — their org,
	// company, or team identifier for an end user. Asserting a tenant upserts the
	// tenant record in the workspace (keyed on `id` as the tenant's external_id) and
	// associates the created resource with it.
	Tenant TenantAssertionParam `json:"tenant,omitzero"`
	paramObj
}

func (r WidgetSessionSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow WidgetSessionSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WidgetSessionSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delete tenant widget sessions response.
type WidgetSessionDeleteTenantResponse struct {
	// Number of conversations (objectives) deleted along with the sessions.
	ObjectivesDeleted int64 `json:"objectivesDeleted"`
	// Number of sessions deleted.
	SessionsDeleted int64 `json:"sessionsDeleted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectivesDeleted respjson.Field
		SessionsDeleted   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WidgetSessionDeleteTenantResponse) RawJSON() string { return r.JSON.raw }
func (r *WidgetSessionDeleteTenantResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WidgetSessionNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// WidgetSessionSpec is the configuration of a session, fixed at mint.
	Spec WidgetSessionSpecParam `json:"spec,omitzero" api:"required"`
	// CreateOperationMetadata contains the user-provided fields for creating an
	// operation. Read-only fields (id, account_id, workspace_id, created_at,
	// profile_id) are excluded since they are set by the server.
	Metadata shared.CreateOperationMetadataParam `json:"metadata,omitzero"`
	// Secrets to attach to the session.
	Secrets []WidgetSessionNewParamsSecret `json:"secrets,omitzero"`
	paramObj
}

func (r WidgetSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WidgetSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WidgetSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Secret is a named credential attached to the session — typically a token the
// customer's backend minted for the visitor, so the agent acts against their API
// as that subject. Values are captured at the boundary, encrypted at rest,
// appended to every conversation the session creates (re-synced on each turn), and
// never returned by any API. Session secrets take precedence over workspace and
// tool-set secrets of the same name.
type WidgetSessionNewParamsSecret struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r WidgetSessionNewParamsSecret) MarshalJSON() (data []byte, err error) {
	type shadow WidgetSessionNewParamsSecret
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WidgetSessionNewParamsSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WidgetSessionGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type WidgetSessionListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, the `info` field on each returned session is populated. Requests with
	// this flag count more against your rate limit.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time).
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Filter to sessions asserted for a subject. Accepts the canonical `subj_…` form
	// or the `external_id:<value>` form; the external_id form is scoped within a
	// tenant and requires `tenant_id` to also be set.
	SubjectID param.Opt[string] `query:"subjectId,omitzero" json:"-"`
	// Filter to sessions belonging to a tenant. Accepts the canonical `tenant_…` form
	// or the `external_id:<value>` form.
	TenantID param.Opt[string] `query:"tenantId,omitzero" json:"-"`
	// Filter to sessions on a specific widget. Accepts the canonical `wgt_…` form or
	// the `external_id:<value>` form.
	WidgetID param.Opt[string] `query:"widgetId,omitzero" json:"-"`
	// Filter by state.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_EXPIRED", "STATE_REVOKED",
	// "STATE_EXHAUSTED".
	State WidgetSessionListParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WidgetSessionListParams]'s query parameters as
// `url.Values`.
func (r WidgetSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by state.
type WidgetSessionListParamsState string

const (
	WidgetSessionListParamsStateStateUnspecified WidgetSessionListParamsState = "STATE_UNSPECIFIED"
	WidgetSessionListParamsStateStateActive      WidgetSessionListParamsState = "STATE_ACTIVE"
	WidgetSessionListParamsStateStateExpired     WidgetSessionListParamsState = "STATE_EXPIRED"
	WidgetSessionListParamsStateStateRevoked     WidgetSessionListParamsState = "STATE_REVOKED"
	WidgetSessionListParamsStateStateExhausted   WidgetSessionListParamsState = "STATE_EXHAUSTED"
)

type WidgetSessionDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type WidgetSessionDeleteTenantParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Tenant whose sessions to delete. Required — an empty value is rejected rather
	// than matching everything. Accepts the canonical `tenant_…` form or the
	// `external_id:<value>` form.
	TenantID param.Opt[string] `query:"tenantId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WidgetSessionDeleteTenantParams]'s query parameters as
// `url.Values`.
func (r WidgetSessionDeleteTenantParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WidgetSessionRevokeParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r WidgetSessionRevokeParams) MarshalJSON() (data []byte, err error) {
	type shadow WidgetSessionRevokeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WidgetSessionRevokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
