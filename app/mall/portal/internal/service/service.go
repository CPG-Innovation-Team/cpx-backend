package service

import (
	v1 "cpx/api/mall/portal/v1"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}