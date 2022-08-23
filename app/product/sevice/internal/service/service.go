package service

import (
	"cpx-backend/api/product/service/v1"
	"cpx-backend/app/product/sevice/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewProductService)

// ProductService is a product service.
type ProductService struct {
	v1.UnimplementedProductServer

	uc *biz.ProductUseCase
}

// NewProductService new a product service.
func NewProductService(uc *biz.ProductUseCase) *ProductService {
	return &ProductService{uc: uc}
}
