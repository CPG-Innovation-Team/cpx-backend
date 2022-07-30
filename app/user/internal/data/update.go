package data

import (
	"context"
	"cpx/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2022/7/28
  @description:
**/

type updateRepo struct {
	data *Data
	log  *log.Helper
}

// NewUpdateRepo .
func NewUpdateRepo(data *Data, logger log.Logger) biz.UpdateRepo {
	return &updateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *updateRepo) Create(ctx context.Context) {}
