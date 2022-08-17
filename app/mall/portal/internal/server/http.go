package server

import (
	v1 "cpx-backend/api/mall/portal/v1"
	"cpx-backend/app/mall/portal/internal/conf"
	"cpx-backend/app/mall/portal/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"golang.org/x/net/context"

	jwtV4 "github.com/golang-jwt/jwt/v4"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/mall.portal.v1.PortalService/Login"] = struct{}{}
	whiteList["/mall.portal.v1.PortalService/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, portal *service.PortalService, tp *traceSdk.TracerProvider, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			selector.Server(
				jwt.Server(func(token *jwtV4.Token) (interface{}, error) {
					return []byte(ac.ApiKey), nil
				},
					jwt.WithSigningMethod(jwtV4.SigningMethodHS256),
					jwt.WithClaims(func() jwtV4.Claims {
						return &jwtV4.MapClaims{}
					})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterPortalHTTPServer(srv, portal)
	return srv
}
