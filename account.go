package cadenya

import (
	"context"
	"go.cadenya.com/cadenya-go/internal/apijson"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/respjson"
	"go.cadenya.com/cadenya-go/shared"
	"net/http"
	"slices"
)

// Manage the authenticated account. Accounts are the top-level organizational unit
// and contain one or more workspaces.
//
// AccountService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.options = opts
	return
}

// Retrieves the current account for the token accessing the API. Useful to check
// if the credentials are valid.
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rotates the challenge token sent in the X-Cadenya-Challenge-Token header on MCP
// tools/list requests. Returns only the new token.
func (r *AccountService) RotateChallengeToken(ctx context.Context, opts ...option.RequestOption) (res *RotateChallengeTokenResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/account:rotateChallengeToken"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Rotates the webhook signing key for the account. Returns only the new key.
func (r *AccountService) RotateWebhookSigningKey(ctx context.Context, opts ...option.RequestOption) (res *RotateWebhookSigningKeyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/account:rotateWebhookSigningKey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// An account, the top-level organizational unit. Contains workspaces and
// account-wide settings such as the webhook signing secret.
type Account struct {
	// Server-populated information about the account.
	Info AccountInfo `json:"info" api:"required"`
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata shared.AccountResourceMetadata `json:"metadata" api:"required"`
	// Configuration for an account.
	Spec AccountSpec `json:"spec" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Info        respjson.Field
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Server-populated information about the account.
type AccountInfo struct {
	// The challenge token Cadenya sends in the X-Cadenya-Challenge-Token header on
	// every MCP tools/list request. Server implementations can accept a valid
	// challenge token in place of per-user auth when listing tools, while still
	// requiring real auth on tools/call. Rotate with RotateChallengeToken; update any
	// servers validating the token before rotating.
	ChallengeToken string `json:"challengeToken"`
	// The generated secret that will sign all webhooks that are sent to your
	// configured Webhook URL. Formatted as "wh_asdf1234" per the
	// https://www.standardwebhooks.com/ format.
	WebhookEventsHMACSecret string `json:"webhookEventsHmacSecret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChallengeToken          respjson.Field
		WebhookEventsHMACSecret respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountInfo) RawJSON() string { return r.JSON.raw }
func (r *AccountInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an account.
type AccountSpec struct {
	BillingEmail string      `json:"billingEmail"`
	Description  string      `json:"description"`
	Domain       string      `json:"domain"`
	Workspaces   []Workspace `json:"workspaces"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingEmail respjson.Field
		Description  respjson.Field
		Domain       respjson.Field
		Workspaces   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountSpec) RawJSON() string { return r.JSON.raw }
func (r *AccountSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A profile identifies a user or non-human principal (such as an API key) at the
// account level. Profiles are account-scoped and can be granted access to multiple
// workspaces.
type Profile struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata shared.AccountResourceMetadata `json:"metadata" api:"required"`
	// Configuration for a profile.
	Spec ProfileSpec `json:"spec" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Profile) RawJSON() string { return r.JSON.raw }
func (r *Profile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a profile.
type ProfileSpec struct {
	// Whether this profile represents a human user, an API key, or a system principal.
	//
	// Any of "PROFILE_TYPE_UNSPECIFIED", "PROFILE_TYPE_USER", "PROFILE_TYPE_API_KEY",
	// "PROFILE_TYPE_SYSTEM".
	Type ProfileSpecType `json:"type" api:"required"`
	// Email address of the profile. Required and unique within an account for user
	// profiles.
	Email string `json:"email"`
	// Display name (e.g., "Bobby Tables").
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileSpec) RawJSON() string { return r.JSON.raw }
func (r *ProfileSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this profile represents a human user, an API key, or a system principal.
type ProfileSpecType string

const (
	ProfileSpecTypeProfileTypeUnspecified ProfileSpecType = "PROFILE_TYPE_UNSPECIFIED"
	ProfileSpecTypeProfileTypeUser        ProfileSpecType = "PROFILE_TYPE_USER"
	ProfileSpecTypeProfileTypeAPIKey      ProfileSpecType = "PROFILE_TYPE_API_KEY"
	ProfileSpecTypeProfileTypeSystem      ProfileSpecType = "PROFILE_TYPE_SYSTEM"
)

// Response containing the newly generated challenge token.
type RotateChallengeTokenResponse struct {
	ChallengeToken string `json:"challengeToken"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChallengeToken respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RotateChallengeTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *RotateChallengeTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing the newly generated webhook signing secret.
type RotateWebhookSigningKeyResponse struct {
	WebhookEventsHMACSecret string `json:"webhookEventsHmacSecret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WebhookEventsHMACSecret respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RotateWebhookSigningKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *RotateWebhookSigningKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
