// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/cadenya/cadenya-sdk-go/internal/apijson"
	"github.com/cadenya/cadenya-sdk-go/internal/requestconfig"
	"github.com/cadenya/cadenya-sdk-go/option"
)

type CursorPaginationPagination struct {
	NextCursor string                         `json:"next_cursor"`
	JSON       cursorPaginationPaginationJSON `json:"-"`
}

// cursorPaginationPaginationJSON contains the JSON metadata for the struct
// [CursorPaginationPagination]
type cursorPaginationPaginationJSON struct {
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorPaginationPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorPaginationPaginationJSON) RawJSON() string {
	return r.raw
}

type CursorPagination[T any] struct {
	Items      []T                        `json:"items"`
	Pagination CursorPaginationPagination `json:"pagination"`
	JSON       cursorPaginationJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// cursorPaginationJSON contains the JSON metadata for the struct
// [CursorPagination[T]]
type cursorPaginationJSON struct {
	Items       apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CursorPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPagination[T]) GetNextPage() (res *CursorPagination[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPaginationAutoPager[T any] struct {
	page *CursorPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewCursorPaginationAutoPager[T any](page *CursorPagination[T], err error) *CursorPaginationAutoPager[T] {
	return &CursorPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPaginationAutoPager[T]) Index() int {
	return r.run
}
