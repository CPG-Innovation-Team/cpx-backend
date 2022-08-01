package service

import (
	v1 "cpx-backend/api/mall/portal/v1"
	"cpx-backend/app/mall/portal/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPortalService)

type PortalService struct {
	v1.UnimplementedPortalServer
	uc *biz.UserUseCase
}

func NewPortalService(uc *biz.UserUseCase) *PortalService {
	return &PortalService{
		uc: uc,
	}
}
