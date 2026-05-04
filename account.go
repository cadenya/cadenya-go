// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"net/http"
	"slices"

	"github.com/cadenya/cadenya-go/internal/apijson"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/shared"
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
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Retrieves the current account for the token accessing the API. Useful to check
// if the credentials are valid.
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rotates the webhook signing key for the account. Returns only the new key.
func (r *AccountService) RotateWebhookSigningKey(ctx context.Context, opts ...option.RequestOption) (res *RotateWebhookSigningKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account/rotate_webhook_signing_key"
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
	JSON accountJSON `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	Info        apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountJSON) RawJSON() string {
	return r.raw
}

// Server-populated information about the account.
type AccountInfo struct {
	// The generated secret that will sign all webhooks that are sent to your
	// configured Webhook URL. Formatted as "wh_asdf1234" per the
	// https://www.standardwebhooks.com/ format.
	WebhookEventsHmacSecret string          `json:"webhookEventsHmacSecret"`
	JSON                    accountInfoJSON `json:"-"`
}

// accountInfoJSON contains the JSON metadata for the struct [AccountInfo]
type accountInfoJSON struct {
	WebhookEventsHmacSecret apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountInfoJSON) RawJSON() string {
	return r.raw
}

// Configuration for an account.
type AccountSpec struct {
	BillingEmail string          `json:"billingEmail"`
	Description  string          `json:"description"`
	Domain       string          `json:"domain"`
	Workspaces   []Workspace     `json:"workspaces"`
	JSON         accountSpecJSON `json:"-"`
}

// accountSpecJSON contains the JSON metadata for the struct [AccountSpec]
type accountSpecJSON struct {
	BillingEmail apijson.Field
	Description  apijson.Field
	Domain       apijson.Field
	Workspaces   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSpecJSON) RawJSON() string {
	return r.raw
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
	JSON profileJSON `json:"-"`
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

// A profile identifies a user or non-human principal (such as an API key) at the
// account level. Profiles are account-scoped and can be granted access to multiple
// workspaces.
type ProfileParam struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata param.Field[shared.AccountResourceMetadataParam] `json:"metadata" api:"required"`
	// Configuration for a profile.
	Spec param.Field[ProfileSpecParam] `json:"spec" api:"required"`
}

func (r ProfileParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a profile.
type ProfileSpec struct {
	// Whether this profile represents a human user, an API key, or a system principal.
	Type ProfileSpecType `json:"type" api:"required"`
	// Email address of the profile. Required and unique within an account for user
	// profiles.
	Email string `json:"email"`
	// Display name (e.g., "Bobby Tables").
	Name string          `json:"name"`
	JSON profileSpecJSON `json:"-"`
}

// profileSpecJSON contains the JSON metadata for the struct [ProfileSpec]
type profileSpecJSON struct {
	Type        apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileSpecJSON) RawJSON() string {
	return r.raw
}

// Whether this profile represents a human user, an API key, or a system principal.
type ProfileSpecType string

const (
	ProfileSpecTypeProfileTypeUser   ProfileSpecType = "PROFILE_TYPE_USER"
	ProfileSpecTypeProfileTypeAPIKey ProfileSpecType = "PROFILE_TYPE_API_KEY"
	ProfileSpecTypeProfileTypeSystem ProfileSpecType = "PROFILE_TYPE_SYSTEM"
)

func (r ProfileSpecType) IsKnown() bool {
	switch r {
	case ProfileSpecTypeProfileTypeUser, ProfileSpecTypeProfileTypeAPIKey, ProfileSpecTypeProfileTypeSystem:
		return true
	}
	return false
}

// Configuration for a profile.
type ProfileSpecParam struct {
	// Whether this profile represents a human user, an API key, or a system principal.
	Type param.Field[ProfileSpecType] `json:"type" api:"required"`
	// Email address of the profile. Required and unique within an account for user
	// profiles.
	Email param.Field[string] `json:"email"`
	// Display name (e.g., "Bobby Tables").
	Name param.Field[string] `json:"name"`
}

func (r ProfileSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response containing the newly generated webhook signing secret.
type RotateWebhookSigningKeyResponse struct {
	WebhookEventsHmacSecret string                              `json:"webhookEventsHmacSecret"`
	JSON                    rotateWebhookSigningKeyResponseJSON `json:"-"`
}

// rotateWebhookSigningKeyResponseJSON contains the JSON metadata for the struct
// [RotateWebhookSigningKeyResponse]
type rotateWebhookSigningKeyResponseJSON struct {
	WebhookEventsHmacSecret apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RotateWebhookSigningKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rotateWebhookSigningKeyResponseJSON) RawJSON() string {
	return r.raw
}
