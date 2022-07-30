package data

import (
	userV1 "cpx/api/user/v1"
	"cpx/app/mall/portal/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewDiscovery,
	NewRegistrar,
)

// Data .
type Data struct {
	log *log.Helper
	uc  userV1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc userV1.UserClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc}, nil
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
