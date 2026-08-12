package cadenya

import (
	"context"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"net/http"
	"slices"
)

// Operations on profiles, the account-level principals (users, API keys, system)
// that authenticate against the API.
//
// ProfileService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.options = opts
	return
}

// Retrieves the profile of the authenticated caller. Useful to check which
// principal a token belongs to.
func (r *ProfileService) Whoami(ctx context.Context, opts ...option.RequestOption) (res *Profile, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/whoami"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
