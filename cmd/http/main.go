package main

import (
	"context"
	"log"
	"main/config"
	_ "main/docs"
	authServices "main/internal/auth/v1/application/service"
	authHandlers "main/internal/auth/v1/handler/http"
	authRepos "main/internal/auth/v1/infrastructure/repository"
	commonServices "main/internal/common/v1/application/service"
	commonHandlers "main/internal/common/v1/handler/http"
	commonRepos "main/internal/common/v1/infrastructure/repository"
	imsServices "main/internal/inventory/v1/application/service"
	imsHandlers "main/internal/inventory/v1/handler/http"
	imsRepos "main/internal/inventory/v1/infrastructure/repository"
	"main/pkg/databases/postgresql"
	"main/pkg/logger"
	"main/pkg/server"
	"main/pkg/utils/graceful_exit"
)

func main() {

	//os.Setenv("APP_ENV", "dev")

	cfg, errConfig := config.ParseConfig()
	if errConfig != nil {
		log.Fatal(errConfig)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.APP_VERSION, cfg.Logger.LEVEL, cfg.Server.APP_ENV)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init Clients
	postgresqlDB, err := postgresql.NewPostgresqlDB(cfg)
	if err != nil {
		appLogger.Fatal("Error when tyring to connect to Postgresql")
	} else {
		appLogger.Info("Postgresql connected")
	}

	// Init repositories
	pgRepoIms := imsRepos.NewPostgresqlRepository(ctx, postgresqlDB)
	pgRepoAuth := authRepos.NewPostgresqlRepository(ctx, postgresqlDB)
	pgRepoCommon := commonRepos.NewPostgresqlRepository(ctx, postgresqlDB)

	// Init services
	imsService := imsServices.NewService(ctx, cfg, pgRepoIms, appLogger)
	authService := authServices.NewService(ctx, cfg, pgRepoAuth, appLogger)
	commonService := commonServices.NewService(ctx, cfg, pgRepoCommon, appLogger)

	servers := server.NewServer(cfg, &ctx, appLogger)

	httpServer, errHttpServer := servers.NewHttpServer()
	if errHttpServer != nil {
		println(errHttpServer.Error())
	}
	versioning := httpServer.Group("/v1")

	// Init handlers for HTTP Server

	imsHandler := imsHandlers.NewHttpHandler(ctx, cfg, imsService, appLogger)
	authHandler := authHandlers.NewHttpHandler(ctx, cfg, authService, appLogger)
	commonHandler := commonHandlers.NewHttpHandler(ctx, cfg, commonService, appLogger)

	// Init routes for HTTP Server
	imsHandlers.MapRoutes(imsHandler, versioning)
	authHandlers.MapRoutes(authHandler, versioning)
	commonHandlers.MapRoutes(commonHandler, versioning)

	// Exit from application gracefully
	graceful_exit.TerminateApp(ctx)

	appLogger.Info("Server Exited Properly")
}
