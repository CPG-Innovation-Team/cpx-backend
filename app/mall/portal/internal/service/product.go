package service

import (
	v1 "cpx-backend/api/mall/portal/v1"
	"cpx-backend/app/mall/portal/internal/biz"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2022/8/11
  @description:
**/

func (s *PortalService) AddProduct(ctx context.Context, req *v1.AddProductRequest) (*v1.AddProductReply, error) {
	x, err := s.pc.AddProduct(ctx, &biz.Product{
		Sid:   req.Sid,
		Name:  req.Name,
		Price: req.Price,
		Cover: req.Cover,
		Type:  req.Type,
		State: req.State,
	})
	reply := &v1.AddProductReply{}
	_ = copier.Copy(reply, x)
	return reply, err
}

func (s *PortalService) UpdateProduct(ctx context.Context, req *v1.UpdateProductRequest) (*v1.UpdateProductReply, error) {
	x, err := s.pc.UpdateProduct(ctx, &biz.Product{
		Sku:   req.Sku,
		Sid:   req.Sid,
		Name:  req.Name,
		Price: req.Price,
		Cover: req.Cover,
		Type:  req.Type,
		State: req.State,
	})
	reply := &v1.UpdateProductReply{}
	_ = copier.Copy(reply, x)
	return reply, err
}

func (s *PortalService) GetProductDetail(ctx context.Context, req *v1.GetProductDetailRequest) (*v1.GetProductDetailReply, error) {
	product, err := s.pc.GetProductDetail(ctx, req.Sku)
	reply := v1.GetProductDetailReply{}
	_ = copier.Copy(&reply, product)
	return &reply, err
}

func (s *PortalService) GetProductList(ctx context.Context, req *v1.GetProductListRequest) (*v1.GetProductListReply, error) {
	p, err := s.pc.GetProductList(ctx, req.Sku)
	reply := v1.GetProductListReply{}
	_ = copier.Copy(&reply, p)
	return &reply, err
}
