// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"net/http"
	"slices"

	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
)

// Manage the account's system-provisioned global API key. The global key is the
// only key that spans every workspace; it is created by the system and cannot be
// deleted, so the surface is retrieve, rotate, and the disable/enable kill switch.
//
// GlobalAPIKeyService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalAPIKeyService] method instead.
type GlobalAPIKeyService struct {
	Options []option.RequestOption
}

// NewGlobalAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGlobalAPIKeyService(opts ...option.RequestOption) (r *GlobalAPIKeyService) {
	r = &GlobalAPIKeyService{}
	r.Options = opts
	return
}

// Retrieves the account's global API key. The token is included only when the
// caller's scopes dominate the key's.
func (r *GlobalAPIKeyService) Get(ctx context.Context, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account/global_api_key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disables the global API key. While disabled, presenting its token fails
// authentication on every endpoint; the key is retained. Idempotent.
func (r *GlobalAPIKeyService) Disable(ctx context.Context, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account/global_api_key:disable"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Re-enables the disabled global API key so its token authenticates again.
// Idempotent.
func (r *GlobalAPIKeyService) Enable(ctx context.Context, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account/global_api_key:enable"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Rotates the global API key and returns a new token. All previous tokens are
// invalidated.
func (r *GlobalAPIKeyService) Rotate(ctx context.Context, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account/global_api_key:rotate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}
