package service

import (
	"context"
	"cpx-backend/api/user/v1"
)

func (s *UserService) Register(ctx context.Context, req *v1.CreateRequest) (*v1.CreateReply, error) {
	return &v1.CreateReply{}, nil
}
func (s *UserService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{}, nil
}
func (s *UserService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}
func (s *UserService) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateReply, error) {
	return &v1.UpdateReply{}, nil
}
func (s *UserService) GetUserProfile(ctx context.Context, req *v1.GetUserProfileRequest) (*v1.UserProfileReply, error) {
	return &v1.UserProfileReply{}, nil
}
func (s *UserService) GetUserList(ctx context.Context, req *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	return &v1.GetUserListReply{}, nil
}
