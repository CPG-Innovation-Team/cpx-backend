package biz

import (
	"cpx-backend/app/mall/portal/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2022/8/11
  @description:
**/

type Product struct {
	Id    int64
	Sid   int64
	Sku   string
	Name  string
	Price float64
	Cover string
	Type  int64
	State int64
}

// ProductRepo is a retrieve product information repository.
type ProductRepo interface {
	GetProductDetail(ctx context.Context, sku string) (product *Product, err error)
	GetProductList(ctx context.Context, sku []string) (product *[]Product, err error)

	AddProduct(ctx context.Context, req *Product) (product *Product, err error)
	UpdateProduct(ctx context.Context, req *Product) (product *Product, err error)
}

// ProductUseCase is a product use case.
type ProductUseCase struct {
	key string
	pr  ProductRepo
	log *log.Helper
}

// NewProductUseCase new a product use case.
func NewProductUseCase(pr ProductRepo, conf *conf.Auth, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{key: conf.ApiKey, pr: pr, log: log.NewHelper(logger)}
}

func (pc ProductUseCase) GetProductDetail(ctx context.Context, sku string) (product *Product, err error) {
	return pc.pr.GetProductDetail(ctx, sku)
}

func (pc ProductUseCase) GetProductList(ctx context.Context, sku []string) (product *[]Product, err error) {
	return pc.pr.GetProductList(ctx, sku)
}

func (pc ProductUseCase) AddProduct(ctx context.Context, req *Product) (product *Product, err error) {
	return pc.pr.AddProduct(ctx, req)
}

func (pc ProductUseCase) UpdateProduct(ctx context.Context, req *Product) (product *Product, err error) {
	return pc.pr.UpdateProduct(ctx, req)
}
