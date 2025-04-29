package server

import (
	v1 "explorer/api/explorer/v1"
	"explorer/internal/conf"
	"explorer/internal/middleware"
	"explorer/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, basicService *service.BasicService, userService *service.UserService, chainService *service.ChainService, logger log.Logger) *http.Server {
	authMiddleware := middleware.NewAuthMiddleware(userService.UserManager, logger)

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			middleware.AuthMiddlewareWrap(authMiddleware),
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
	v1.RegisterBasicHTTPServer(srv, basicService)
	v1.RegisterUserHTTPServer(srv, userService)
	v1.RegisterChainHTTPServer(srv, chainService)
	return srv
}
