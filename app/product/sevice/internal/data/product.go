package data

import (
	"context"
	"cpx-backend/app/product/sevice/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

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

func (r *productRepo) Save(ctx context.Context, g *biz.Product) (*biz.Product, error) {
	return g, nil
}
