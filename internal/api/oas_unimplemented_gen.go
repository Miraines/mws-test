// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateCat implements createCat operation.
//
// Creates a new cat and returns its complete representation.
//
// POST /cats
func (UnimplementedHandler) CreateCat(ctx context.Context, req *NewCat) (r *Cat, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteCat implements deleteCat operation.
//
// Deletes a cat by ID. Returns `204 No Content' if deletion was successful.
//
// DELETE /cats/{id}
func (UnimplementedHandler) DeleteCat(ctx context.Context, params DeleteCatParams) (r DeleteCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetCat implements getCat operation.
//
// Finds a cat by ID and returns it.
//
// GET /cats/{id}
func (UnimplementedHandler) GetCat(ctx context.Context, params GetCatParams) (r GetCatRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCats implements listCats operation.
//
// Returns the full list of cats stored in the service's memory.
//
// GET /cats
func (UnimplementedHandler) ListCats(ctx context.Context) (r []Cat, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateCat implements updateCat operation.
//
// Updates the data of the cat by ID and returns the updated record.
//
// PUT /cats/{id}
func (UnimplementedHandler) UpdateCat(ctx context.Context, req *UpdateCat, params UpdateCatParams) (r UpdateCatRes, _ error) {
	return r, ht.ErrNotImplemented
}
