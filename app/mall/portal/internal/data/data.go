package data

import (
	productV1 "cpx-backend/api/product/service/v1"
	userV1 "cpx-backend/api/user/service/v1"
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
	NewProductRepo,
	NewDiscovery,
	NewUserServiceClient,
	NewProductServiceClient,
)

// Data .
type Data struct {
	log *log.Helper
	pc  productV1.ProductClient
	uc  userV1.UserClient
}

// NewData .
func NewData(
	c *conf.Data,
	logger log.Logger,
	pc productV1.ProductClient,
	uc userV1.UserClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, pc: pc, uc: uc}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
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
			ClientConfig: cc,
		},
	)

	if err != nil {
		panic(err)
	}

	return nacos.New(cli)
}

func NewUserServiceClient(r registry.Discovery) userV1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userV1.NewUserClient(conn)
	return c
}

func NewProductServiceClient(r registry.Discovery) productV1.ProductClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mall.product.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := productV1.NewProductClient(conn)
	return c
}
