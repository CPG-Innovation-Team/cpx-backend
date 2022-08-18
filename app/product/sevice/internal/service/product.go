package service

import (
	"context"

	v1 "cpx-backend/api/product/v1"
)

func (s *ProductService) GetProductDetailBySku(ctx context.Context, req *v1.GetProductDetailRequest) (*v1.GetProductDetailReply, error) {
	return &v1.GetProductDetailReply{}, nil
}
func (s *ProductService) GetProductList(ctx context.Context, req *v1.GetProductListRequest) (*v1.GetProductListReply, error) {
	return &v1.GetProductListReply{}, nil
}
func (s *ProductService) AddProduct(ctx context.Context, req *v1.AddProductRequest) (*v1.AddProductReply, error) {
	return &v1.AddProductReply{}, nil
}
func (s *ProductService) UpdateProduct(ctx context.Context, req *v1.UpdateProductRequest) (*v1.UpdateProductReply, error) {
	return &v1.UpdateProductReply{}, nil
}
