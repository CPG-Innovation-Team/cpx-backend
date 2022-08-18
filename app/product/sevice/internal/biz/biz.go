package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewProductUseCase)

// ProductUseCase is a product use case.
type ProductUseCase struct {
	repo ProductRepo
	log  *log.Helper
}

// NewProductUseCase new a product use case.
func NewProductUseCase(repo ProductRepo, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, log: log.NewHelper(logger)}
}
