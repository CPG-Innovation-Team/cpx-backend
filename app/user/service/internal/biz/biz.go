package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)

// UserUseCase is a user use case.
type UserUseCase struct {
	rr  RetrieveRepo
	ur  UpdateRepo
	log *log.Helper
}

// NewUserUseCase new a user use case.
func NewUserUseCase(rr RetrieveRepo, ur UpdateRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{rr: rr, ur: ur, log: log.NewHelper(logger)}
}
