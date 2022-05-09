package service

import (
	"context"
	"fmt"
	"github.com/husanmusa/pro-book-service/config"
	pb "github.com/husanmusa/pro-book-service/genproto/book_service"
	"github.com/husanmusa/pro-book-service/grpc/client"
	"github.com/husanmusa/pro-book-service/storage"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
	"regexp"
)

type bookService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedBookServiceServer
}

func NewBookService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *bookService {
	return &bookService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *bookService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	s.log.Info("---CreateUser--->", logger.Any("req", req))

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	email := emailRegex.MatchString(req.Email)
	if !email {
		err := fmt.Errorf("email is not valid")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}
	phoneRegex := regexp.MustCompile(`^[+]?(\d{1,2})?[\s.-]?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`)
	phone := phoneRegex.MatchString(req.Phone)
	if !phone {
		err := fmt.Errorf("phone number is not valid")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	return s.strg.BookService().CreateUser(ctx, req)
}

func (s *bookService) GetUser(ctx context.Context, req *pb.ByIdReq) (*pb.User, error) {
	s.log.Info("---GetUser--->", logger.Any("req", req))

	res, err := s.strg.BookService().GetUser(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUser--->", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *bookService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	s.log.Info("---UpdateUser--->", logger.Any("req", req))

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	email := emailRegex.MatchString(req.Email)
	if !email {
		err := fmt.Errorf("email is not valid")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}

	phoneRegex := regexp.MustCompile(`^[+]?(\d{1,2})?[\s.-]?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`)
	phone := phoneRegex.MatchString(req.Phone)
	if !phone {
		err := fmt.Errorf("phone number is not valid")
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, err
	}
	fmt.Println("BOOKING", req.Id)

	user, err := s.strg.BookService().UpdateUser(ctx, req)

	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *bookService) DeleteUser(ctx context.Context, req *pb.ByIdReq) (*emptypb.Empty, error) {
	s.log.Info("---DeleteUser--->", logger.Any("req", req))

	res, err := s.strg.BookService().DeleteUser(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteUser--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.log.Info("---ListUsers--->", logger.Any("req", req))

	res, err := s.strg.BookService().ListUsers(ctx, req)
	if err != nil {
		s.log.Error("!!!ListUsers--->", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *bookService) FindBookFromUsers(ctx context.Context, req *pb.ByIdReq) (*pb.ListUsersResponse, error) {
	s.log.Info("---ListUsers--->", logger.Any("req", req))
	fmt.Println("fsdfadsf", req)
	res, err := s.strg.BookService().FindBookFromUsers(ctx, req)
	if err != nil {
		s.log.Error("!!!ListUsers--->", logger.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *bookService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Book, error) {
	s.log.Info("---CreateBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().CreateBook(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateBook--->", logger.Error(err))
	}

	return res, err
}

func (s *bookService) GetBook(ctx context.Context, req *pb.ByIdReq) (*pb.Book, error) {
	s.log.Info("---GetBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().GetBook(ctx, req)
	if err != nil {
		s.log.Error("!!!GetBook--->", logger.Error(err))
	}

	return res, err
}

func (s *bookService) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.Book, error) {
	s.log.Info("---UpdateBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().UpdateBook(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateBook--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) DeleteBook(ctx context.Context, req *pb.ByIdReq) (*emptypb.Empty, error) {
	s.log.Info("---DeleteBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().DeleteBook(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteBook--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	s.log.Info("---ListBooks--->", logger.Any("req", req))

	res, err := s.strg.BookService().ListBooks(ctx, req)
	if err != nil {
		s.log.Error("!!!ListBooks--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) ListBooksByUserId(ctx context.Context, req *pb.ByIdReq) (*pb.ListBooksResponse, error) {
	s.log.Info("---ListBooks--->", logger.Any("req", req))
	res, err := s.strg.BookService().ListBooksByUserId(ctx, req)
	if err != nil {
		s.log.Error("!!!ListBooks--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) CreateSaleBook(ctx context.Context, req *pb.CreateSaleBookRequest) (*pb.SaleBook, error) {
	s.log.Info("---CreateSaleBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().CreateSaleBook(ctx, req)
	if err != nil {
		s.log.Error("!!!CreateSaleBook--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) GetSaleBook(ctx context.Context, req *pb.ByIdReq) (*pb.SaleBook, error) {
	s.log.Info("---GetSaleBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().GetSaleBook(ctx, req)
	if err != nil {
		s.log.Error("!!!GetSaleBook--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) UpdateSaleBook(ctx context.Context, req *pb.UpdateSaleBookRequest) (*pb.SaleBook, error) {
	s.log.Info("---UpdateSaleBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().UpdateSaleBook(ctx, req)
	if err != nil {
		s.log.Error("!!!UpdateSaleBook--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) DeleteSaleBook(ctx context.Context, req *pb.ByIdReq) (*emptypb.Empty, error) {
	s.log.Info("---DeleteSaleBook--->", logger.Any("req", req))

	res, err := s.strg.BookService().DeleteSaleBook(ctx, req)
	if err != nil {
		s.log.Error("!!!DeleteSaleBook--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *bookService) ListSaleBooks(ctx context.Context, req *pb.ListSaleBooksRequest) (*pb.ListSaleBooksResponse, error) {
	s.log.Info("---ListSaleBooks--->", logger.Any("req", req))

	res, err := s.strg.BookService().ListSaleBooks(ctx, req)
	if err != nil {
		s.log.Error("!!!ListSaleBooks--->", logger.Error(err))
		return nil, err
	}

	return res, nil
}
