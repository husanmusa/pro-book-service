package client

import (
	"github.com/husanmusa/pro-book-service/config"
	pb "github.com/husanmusa/pro-book-service/genproto/book_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	BookService() pb.BookServiceClient
}

type grpcClients struct {
	bookService pb.BookServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connAuthService, err := grpc.Dial(
		cfg.AuthServiceHost+cfg.AuthGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService: pb.NewBookServiceClient(connAuthService),
	}, nil
}

func (g grpcClients) BookService() pb.BookServiceClient {
	return g.bookService
}
