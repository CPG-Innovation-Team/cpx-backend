package data

import (
	"context"
	"github.com/jinzhu/copier"

	"cpx-backend/api/user/v1"
	"cpx-backend/app/mall/portal/internal/biz"
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

func (ur *userRepo) FindByEmail(ctx context.Context, email string) (user *biz.User, err error, uid int64) {
	u, err := ur.data.uc.GetUserByEmail(ctx, &v1.GetUserByEmailRequest{
		Email: email,
	})

	if err != nil {
		return nil, biz.ErrUserNotFound, 0
	}

	return &biz.User{
		Name:  u.Username,
		Email: email,
	}, err, u.Uid
}

func (ur *userRepo) VerifyPassword(ctx context.Context, email, password string) (uid int64, err error) {
	reply, err := ur.data.uc.VerifyPassword(ctx, &v1.VerifyPasswordRequest{Email: email, Password: password})
	if !reply.Ok && err != nil {
		return 0, err
	}
	return reply.Uid, nil
}

func (ur *userRepo) GetUserProfile(ctx context.Context, uid int64) (user *biz.User, err error) {
	reply, err := ur.data.uc.GetUserById(ctx, &v1.GetUserByIdRequest{
		Uid: uid,
	})
	_ = copier.Copy(user, reply.User)
	return user, err
}

func (ur userRepo) GetUserList(ctx context.Context, uid []int64, typ int64) (users *[]biz.User, err error) {
	reply, err := ur.data.uc.GetUserListByUid(ctx, &v1.GetUserListRequest{
		Uid:  uid,
		Type: typ,
	})
	_ = copier.Copy(users, reply.User)
	return
}

func (ur *userRepo) Create(ctx context.Context, user *biz.User) (uid int64, err error) {
	createParams := v1.CreateRequest{}
	_ = copier.Copy(createParams.GetUser, user)
	reply, err := ur.data.uc.Create(ctx, &createParams)
	if err != nil {
		return 0, err
	}
	return reply.Uid, nil
}

func (ur userRepo) Update(ctx context.Context, user *biz.User) error {
	updateParams := v1.UpdateRequest{}
	_ = copier.Copy(updateParams.User, user)
	reply, err := ur.data.uc.Update(ctx, &updateParams)
	if !reply.Ok {
		return biz.ErrUpdateUser
	}
	return err
}
