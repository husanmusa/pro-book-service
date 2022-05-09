package main

import (
	"context"
	"github.com/husanmusa/pro-book-service/api"
	"github.com/husanmusa/pro-book-service/api/handlers"
	"github.com/husanmusa/pro-book-service/config"
	"github.com/husanmusa/pro-book-service/grpc"
	"github.com/husanmusa/pro-book-service/grpc/client"
	"github.com/husanmusa/pro-book-service/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
	"net"
)

func main() {
	var loggerLevel string
	cfg := config.Load()

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
		}
	}()

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	grpcServer := grpc.SetupServer(cfg, log, pgStore, svcs)
	go func() {
		lis, err := net.Listen("tcp", cfg.AuthGRPCPort)
		if err != nil {
			log.Panic("net.Listen", logger.Error(err))
		}

		log.Info("GRPC: Server being started...", logger.String("port", cfg.AuthGRPCPort))

		if err := grpcServer.Serve(lis); err != nil {
			log.Panic("grpcServer.Serve", logger.Error(err))
		}
	}()

	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetupRouter(h, cfg)

	if r.Listen(cfg.HTTPPort) != nil {
		panic(err)
	}
}
