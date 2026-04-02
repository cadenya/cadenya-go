// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"net/http"
	"slices"

	"github.com/cadenya/cadenya-sdk-go/internal/apijson"
	"github.com/cadenya/cadenya-sdk-go/internal/requestconfig"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/cadenya/cadenya-sdk-go/shared"
)

// AccountService manages account-level operations. Accounts are the top-level
// organizational unit in the system. All operations are scoped to the
// authenticated account determined by the JWT token.
//
// Authentication: Bearer token (JWT) Scope: Account-level operations
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

type Account struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata shared.AccountResourceMetadata `json:"metadata" api:"required"`
	Spec     AccountSpec                    `json:"spec" api:"required"`
	JSON     accountJSON                    `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
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

type AccountSpec struct {
	BillingEmail string             `json:"billingEmail"`
	Description  string             `json:"description"`
	Domain       string             `json:"domain"`
	Workspaces   []shared.Workspace `json:"workspaces"`
	JSON         accountSpecJSON    `json:"-"`
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
