package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"kratos_learn/api/user"
	"kratos_learn/internal/conf"
	"kratos_learn/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
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
	// auth middleware
	testKey := "mytest"
	opts = append(opts, http.Middleware(jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
		return []byte(testKey), nil
	})))
	// rate limit middleware
	opts = append(opts, http.Middleware(
		// 默认 bbr limiter
		ratelimit.Server(),
	))
	srv := http.NewServer(opts...)
	user.RegisterUserHTTPServer(srv, greeter)
	return srv
}
