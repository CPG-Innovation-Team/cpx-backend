package data

import (
	userV1 "cpx/api/user/v1"
	"cpx/app/mall/portal/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewDiscovery,
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
	c := nacos()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
