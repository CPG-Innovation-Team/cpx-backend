//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"cpx-backend/app/mall/portal/internal/biz"
	"cpx-backend/app/mall/portal/internal/conf"
	"cpx-backend/app/mall/portal/internal/data"
	"cpx-backend/app/mall/portal/internal/server"
	"cpx-backend/app/mall/portal/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	traceSdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, *conf.Auth, log.Logger, *traceSdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
