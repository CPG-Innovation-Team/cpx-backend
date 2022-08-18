package data

import (
	"context"
	"cpx-backend/app/user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2022/7/28
  @description:
**/

type retrieveRepo struct {
	data *Data
	log  *log.Helper
}

// NewRetrieveRepo .
func NewRetrieveRepo(data *Data, logger log.Logger) biz.RetrieveRepo {
	return &retrieveRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (rr *retrieveRepo) GetUserProfile(ctx context.Context) {

}

func (rr *retrieveRepo) GetUserList(ctx context.Context) {}
