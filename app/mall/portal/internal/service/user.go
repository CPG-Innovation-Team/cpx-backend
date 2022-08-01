package service

import (
	"context"

	v1 "cpx-backend/api/mall/portal/v1"
)

func (s *PortalService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	return &v1.RegisterReply{}, nil
}
func (s *PortalService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{}, nil
}
func (s *PortalService) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutReply, error) {
	return &v1.LogoutReply{}, nil
}
func (s *PortalService) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateReply, error) {
	return &v1.UpdateReply{}, nil
}
func (s *PortalService) GetUserProfile(ctx context.Context, req *v1.GetUserProfileRequest) (*v1.UserProfileReply, error) {
	return &v1.UserProfileReply{}, nil
}
func (s *PortalService) GetUserList(ctx context.Context, req *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	return &v1.GetUserListReply{}, nil
}
