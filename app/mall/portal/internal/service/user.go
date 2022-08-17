package service

import (
	"context"
	"github.com/jinzhu/copier"

	v1 "cpx-backend/api/mall/portal/v1"
)

func (s *PortalService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	uid, token, err := s.uc.RegisterUser(ctx, req)
	return &v1.RegisterReply{
		Uid:   uid,
		Token: token,
	}, err
}

func (s *PortalService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	uid, token, err := s.uc.Login(ctx, req.Email, req.Password)
	return &v1.LoginReply{
		Uid:   uid,
		Token: token,
	}, err
}

func (s *PortalService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}

func (s *PortalService) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateReply, error) {
	uid, token, err := s.uc.Update(ctx, req)
	return &v1.UpdateReply{
		Uid:   uid,
		Token: token,
	}, err
}

func (s *PortalService) GetUserProfile(ctx context.Context, req *v1.GetUserProfileRequest) (*v1.UserProfileReply, error) {
	user, err := s.uc.GetUserProfile(ctx, req.Uid)
	reply := v1.UserProfileReply{}
	_ = copier.Copy(&reply, user)
	return &reply, err
}

func (s *PortalService) GetUserList(ctx context.Context, req *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	users, err := s.uc.GetUserList(ctx, req.Uid, req.Type)
	reply := v1.GetUserListReply{}
	_ = copier.Copy(&reply, users)
	return &reply, err
}
