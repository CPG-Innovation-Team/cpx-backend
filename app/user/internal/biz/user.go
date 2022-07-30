package biz

import (
	"context"

	v1 "cpx/api/user/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// RetrieveRepo is a retrieve user information repository.
type RetrieveRepo interface {
	GetUserProfile(ctx context.Context)
	GetUserList(ctx context.Context)
}

// UpdateRepo is an update user information repository.
type UpdateRepo interface {
	Create(ctx context.Context)
}

// UserUsecase is a user usecase.
type UserUsecase struct {
	rr  RetrieveRepo
	ur  UpdateRepo
	log *log.Helper
}

// RegisterUser creates a User, and returns the new User.
func (uc *UserUsecase) RegisterUser(ctx context.Context) error {
	return nil
}