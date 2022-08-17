package data

import (
	productV1 "cpx-backend/api/product/v1"
	"cpx-backend/app/mall/portal/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2022/8/11
  @description:
**/

type productRepo struct {
	data *Data
	log  *log.Helper
}

// NewProductRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (pr productRepo) GetProductDetail(ctx context.Context, sku string) (product *biz.Product, err error) {
	reply, err := pr.data.pc.GetProductDetailBySku(ctx, &productV1.GetProductDetailRequest{
		Sku: sku,
	})
	_ = copier.Copy(product, reply)
	return
}

func (pr productRepo) GetProductList(ctx context.Context, sku []string) (product *[]biz.Product, err error) {
	reply, err := pr.data.pc.GetProductList(ctx, &productV1.GetProductListRequest{
		Sku: sku,
	})
	_ = copier.Copy(product, reply)
	return
}

func (pr productRepo) AddProduct(ctx context.Context, req *biz.Product) (product *biz.Product, err error) {
	request := &productV1.AddProductRequest{}
	_ = copier.Copy(request, req)
	reply, err := pr.data.pc.AddProduct(ctx, request)
	_ = copier.Copy(product, reply)
	return product, err
}

func (pr productRepo) UpdateProduct(ctx context.Context, req *biz.Product) (product *biz.Product, err error) {
	request := &productV1.UpdateProductRequest{}
	_ = copier.Copy(request, req)
	reply, err := pr.data.pc.UpdateProduct(ctx, request)
	_ = copier.Copy(product, reply)
	return product, err
}
