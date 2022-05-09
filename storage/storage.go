package storage

import (
	"context"
	pb "github.com/husanmusa/pro-book-service/genproto/book_service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StorageI interface {
	CloseDB()
	BookService() BookServiceRepoI
}

type BookServiceRepoI interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error)
	GetUser(ctx context.Context, in *pb.ByIdReq) (*pb.User, error)
	UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error)
	DeleteUser(ctx context.Context, in *pb.ByIdReq) (*emptypb.Empty, error)
	ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
	FindBookFromUsers(ctx context.Context, in *pb.ByIdReq) (*pb.ListUsersResponse, error)

	CreateBook(ctx context.Context, in *pb.CreateBookRequest) (*pb.Book, error)
	GetBook(ctx context.Context, in *pb.ByIdReq) (*pb.Book, error)
	UpdateBook(ctx context.Context, in *pb.UpdateBookRequest) (*pb.Book, error)
	DeleteBook(ctx context.Context, in *pb.ByIdReq) (*emptypb.Empty, error)
	ListBooks(ctx context.Context, in *pb.ListBooksRequest) (*pb.ListBooksResponse, error)
	ListBooksByUserId(ctx context.Context, in *pb.ByIdReq) (*pb.ListBooksResponse, error)

	CreateSaleBook(ctx context.Context, in *pb.CreateSaleBookRequest) (*pb.SaleBook, error)
	GetSaleBook(ctx context.Context, in *pb.ByIdReq) (*pb.SaleBook, error)
	UpdateSaleBook(ctx context.Context, in *pb.UpdateSaleBookRequest) (*pb.SaleBook, error)
	DeleteSaleBook(ctx context.Context, in *pb.ByIdReq) (*emptypb.Empty, error)
	ListSaleBooks(ctx context.Context, in *pb.ListSaleBooksRequest) (*pb.ListSaleBooksResponse, error)
}
