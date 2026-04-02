// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"net/http"
	"slices"

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
func (r *AccountService) Get(ctx context.Context, opts ...option.RequestOption) (res *shared.Account, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
