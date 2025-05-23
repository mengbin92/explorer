package main

import (
	"context"
	"explorer/internal/conf"
	"explorer/internal/models/block"
	"explorer/internal/models/users"
	"explorer/internal/server"
	"explorer/internal/service"
	"explorer/provider/cache"
	"explorer/provider/chain"
	"explorer/provider/db"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(ctx context.Context, bc *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	// init db
	if err := db.Init(ctx, bc.Database, logger); err != nil {
		logger.Log(log.LevelFatal, "msg", "init db failed")
		return nil, nil, errors.Wrap(err, "init db failed")
	}
	db.Get().AutoMigrate(&users.User{}, &block.Block{})

	// init redis
	if err := cache.InitRedis(ctx, bc.Redis, logger); err != nil {
		logger.Log(log.LevelFatal, "msg", "init redis failed")
		return nil, nil, errors.Wrap(err, "init redis failed")
	}

	// init chain client
	if bc.ChainConfig.HttpEndpoint != "" {
		if err := chain.InitEthereumHttpClient(ctx, bc.ChainConfig, logger); err != nil {
			logger.Log(log.LevelFatal, "msg", "init ethereum http client failed")
			return nil, nil, errors.Wrap(err, "init ethereum http client failed")
		}
	}
	if bc.ChainConfig.WsEndpoint != "" {
		if err := chain.InitEthereumWSClient(ctx, bc.ChainConfig, logger); err != nil {
			logger.Log(log.LevelFatal, "msg", "init ethereum ws client failed")
			return nil, nil, errors.Wrap(err, "init ethereum ws client failed")
		}
	}

	cleanup := func(ctx context.Context) {
		logger.Log(log.LevelInfo, "msg", "close the data resources")

		// close redis
		cache.GetRedisClient().Close()

		// release ethereum client
		if chain.GetEthereumHttpClient() != nil {
			chain.GetEthereumHttpClient().Close()
		}
		if chain.GetEthereumWSClient() != nil {
			chain.GetEthereumWSClient().Close()
		}
	}

	// init service
	basicClient := service.NewBasicService(logger)
	blockManager,err := block.NewBlockManager(logger)
	if err != nil {
		logger.Log(log.LevelFatal, "msg", "init block manager failed")
		return nil, nil, errors.Wrap(err, "init block manager failed")
	}
	blockManager.WatcherBlock(ctx)

	userManager := users.NewUserManager(bc.AuthConfig, logger)
	userService := service.NewUserService(userManager, logger)
	chainService := service.NewChainService(logger)

	httpServer := server.NewHTTPServer(bc.Server, basicClient, userService, chainService, logger)
	grpcServer := server.NewGRPCServer(bc.Server, basicClient, userService, chainService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup(ctx)
	}, nil
}
