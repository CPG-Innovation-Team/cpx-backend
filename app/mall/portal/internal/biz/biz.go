package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)

// NewUserUseCase new a user use case.
func NewUserUseCase(rr RetrieveRepo, ur UpdateRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{rr: rr, ur: ur, log: log.NewHelper(logger)}
}
