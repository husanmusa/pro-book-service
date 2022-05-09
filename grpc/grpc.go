package grpc

import (
	"github.com/husanmusa/pro-book-service/config"
	pb "github.com/husanmusa/pro-book-service/genproto/book_service"
	"github.com/husanmusa/pro-book-service/grpc/client"
	"github.com/husanmusa/pro-book-service/grpc/service"
	"github.com/husanmusa/pro-book-service/storage"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	pb.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
