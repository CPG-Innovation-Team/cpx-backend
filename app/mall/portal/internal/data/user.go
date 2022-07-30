package data

import (
	"context"

	"cpx/app/mall/portal/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) GetUserProfile(ctx context.Context) {

}

func (ur *userRepo) GetUserList(ctx context.Context) {}

func (ur *userRepo) Create(ctx context.Context) {}
