package data

import (
	v1 "cpx-backend/api/mall/portal/v1"
	"cpx-backend/app/mall/portal/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"golang.org/x/net/context"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewDiscovery,
	NewRegistrar,
	NewUserServiceClient,
)

// Data .
type Data struct {
	log *log.Helper
	pc  v1.PortalClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, pc v1.PortalClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, pc: pc}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Address, 8848),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "info",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	return nacos.New(cli)
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Address, 8848),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "info",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	return nacos.New(cli)
}

func NewUserServiceClient(r registry.Discovery) v1.PortalClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.portal.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewPortalClient(conn)
	return c
}
