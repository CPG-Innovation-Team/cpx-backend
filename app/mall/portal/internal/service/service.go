package service

import (
	v1 "cpx-backend/api/mall/portal/v1"
	"cpx-backend/app/mall/portal/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPortalService)

type PortalService struct {
	v1.UnimplementedPortalServer
	log *log.Helper

	uc *biz.UserUseCase
	pc *biz.ProductUseCase
}

func NewPortalService(
	uc *biz.UserUseCase,
	pc *biz.ProductUseCase,
	logger log.Logger,
) *PortalService {
	return &PortalService{
		log: log.NewHelper(log.With(logger, "module", "service/portal")),
		uc:  uc,
		pc:  pc,
	}
}
