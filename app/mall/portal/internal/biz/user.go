package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	v1 "cpx-backend/api/mall/portal/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// UserRepo is a retrieve user information repository.
type UserRepo interface {
	GetUserProfile(ctx context.Context)
	GetUserList(ctx context.Context)

	Create(ctx context.Context)
}

// UserUseCase is a user use case.
type UserUseCase struct {
	ur  UserRepo
	log *log.Helper
}

// NewUserUseCase new a user use case.
func NewUserUseCase(ur UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{ur: ur, log: log.NewHelper(logger)}
}

// RegisterUser creates a User, and returns the new User.
func (uc *UserUseCase) RegisterUser(ctx context.Context) error {
	return nil
}
