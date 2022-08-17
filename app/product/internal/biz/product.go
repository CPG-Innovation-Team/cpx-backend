package biz

import (
	"context"

	v1 "cpx-backend/api/product/v1"
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Product is a product model.
type Product struct {
}

// ProductRepo is a Greater repo.
type ProductRepo interface {
	Save(context.Context, *Product) (*Product, error)
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ProductUseCase) CreateGreeter(ctx context.Context, g *Product) (*Product, error) {
	return nil, nil
}
