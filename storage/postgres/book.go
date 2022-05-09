package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/husanmusa/pro-book-service/config"
	pb "github.com/husanmusa/pro-book-service/genproto/book_service"
	"github.com/husanmusa/pro-book-service/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type bookServiceRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.BookServiceRepoI {
	return &bookServiceRepo{
		db: db,
	}
}

func (b bookServiceRepo) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	query := `INSERT INTO users (id, first_name, last_name, middle_name, phone, email) 
		VALUES ($1, $2, $3, $4, $5, $6) returning id`
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	fmt.Println("DOING")
	_, err = b.db.Exec(ctx, query,
		uuid,
		in.FirstName,
		in.LastName,
		in.MiddleName,
		in.Phone,
		in.Email)
	if err != nil {
		return nil, err
	}

	newUser, err := b.GetUser(ctx, &pb.ByIdReq{Id: uuid.String()})
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (b bookServiceRepo) GetUser(ctx context.Context, in *pb.ByIdReq) (*pb.User, error) {
	res := &pb.User{}
	query := `SELECT * FROM users WHERE id = $1`
	err := b.db.QueryRow(ctx, query, in.Id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.MiddleName,
		&res.Phone,
		&res.Email,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b bookServiceRepo) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error) {
	query := `UPDATE users SET first_name = $1, last_name = $2, middle_name = $3, 
		phone = $4, email = $5 WHERE id = $6 returning id`
	fmt.Println(in)
	result, err := b.db.Exec(ctx, query,
		in.FirstName,
		in.LastName,
		in.MiddleName,
		in.Phone,
		in.Email,
		in.Id)
	fmt.Println("UPDATING")
	if err != nil {
		return nil, err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	updatedUser, err := b.GetUser(ctx, &pb.ByIdReq{Id: in.Id})
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (b bookServiceRepo) DeleteUser(ctx context.Context, in *pb.ByIdReq) (*emptypb.Empty, error) {
	query := `DELETE FROM users WHERE id = $1`
	result, err := b.db.Exec(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	return &emptypb.Empty{}, nil
}

func (b bookServiceRepo) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	resp := &pb.ListUsersResponse{}
	fmt.Println("HEY1")
	query := `SELECT * FROM users limit $1 offset $2`
	rows, err := b.db.Query(ctx, query, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	var users []*pb.User

	for rows.Next() {
		var user pb.User
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Phone,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	fmt.Println("HEY21")
	resp.Users = users
	return resp, nil
}

func (b bookServiceRepo) FindBookFromUsers(ctx context.Context, in *pb.ByIdReq) (*pb.ListUsersResponse, error) {
	resp := &pb.ListUsersResponse{}

	query := `select u.id, u.first_name, u.last_name, u.middle_name, u.phone, u.email from users u join user_books ub on u.id = ub.user_id where ub.book_id=$1 group by u.id;`
	rows, err := b.db.Query(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}

	var users []*pb.User
	for rows.Next() {
		var user pb.User
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Phone,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	resp.Users = users
	return resp, nil
}

func (b bookServiceRepo) CreateBook(ctx context.Context, in *pb.CreateBookRequest) (*pb.Book, error) {
	query := `insert into books(id, name, count, author, published_date, created_at, updated_at) 
values ($1, $2, $3, $4, $5, $6, $7) returning id`
	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}
	_, err = b.db.Exec(ctx, query,
		uuid,
		in.Name,
		in.Count,
		in.Author,
		in.PublishedDate,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	newBook, err := b.GetBook(ctx, &pb.ByIdReq{Id: uuid.String()})
	if err != nil {
		return nil, err
	}
	return newBook, nil
}

func (b bookServiceRepo) GetBook(ctx context.Context, in *pb.ByIdReq) (*pb.Book, error) {
	var res pb.Book
	query := `SELECT id, name, count, author, 
	TO_CHAR(published_date, ` + config.DatabaseQueryTimeLayout + `) AS published_date,
	TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,	
	TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at

	 FROM books WHERE id = $1`
	err := b.db.QueryRow(ctx, query, in.Id).Scan(
		&res.Id,
		&res.Name,
		&res.Count,
		&res.Author,
		&res.PublishedDate,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		fmt.Println("YA")
		return nil, err
	}
	return &res, nil
}

func (b bookServiceRepo) UpdateBook(ctx context.Context, in *pb.UpdateBookRequest) (*pb.Book, error) {
	query := `UPDATE books SET name = $1, count = $2, author = $3, published_date = $4, 
		updated_at = $5 WHERE id = $6 returning id`
	result, err := b.db.Exec(ctx, query,
		in.Name,
		in.Count,
		in.Author,
		in.PublishedDate,
		time.Now(),
		in.Id,
	)
	if err != nil {
		return nil, err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	updatedBook, err := b.GetBook(ctx, &pb.ByIdReq{Id: in.Id})
	if err != nil {
		return nil, err
	}
	return updatedBook, nil
}

func (b bookServiceRepo) DeleteBook(ctx context.Context, in *pb.ByIdReq) (*emptypb.Empty, error) {
	query := `DELETE FROM books WHERE id = $1`
	result, err := b.db.Exec(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &emptypb.Empty{}, nil
}

func (b bookServiceRepo) ListBooks(ctx context.Context, in *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	resp := &pb.ListBooksResponse{}
	query := `SELECT id, name, count, author, 
TO_CHAR(published_date, ` + config.DatabaseQueryTimeLayout + `) AS published_date,
	TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,	
	TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM books limit $1 offset $2`
	rows, err := b.db.Query(ctx, query, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	var books []*pb.Book

	for rows.Next() {
		var book pb.Book
		err = rows.Scan(
			&book.Id,
			&book.Name,
			&book.Count,
			&book.Author,
			&book.PublishedDate,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	resp.Books = books

	return resp, nil
}

func (b bookServiceRepo) ListBooksByUserId(ctx context.Context, in *pb.ByIdReq) (*pb.ListBooksResponse, error) {
	resp := &pb.ListBooksResponse{}
	query := `SELECT b.id, b.name, b.count, b.author, 
	TO_CHAR(published_date, ` + config.DatabaseQueryTimeLayout + `) AS published_date,
	TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,	
	TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	from books b join user_books ub on b.id = ub.book_id 
	where ub.user_id=$1 group by b.id;`
	fmt.Println(in.Id, "HEE")
	rows, err := b.db.Query(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}

	var books []*pb.Book
	for rows.Next() {
		var book pb.Book
		err = rows.Scan(
			&book.Id,
			&book.Name,
			&book.Count,
			&book.Author,
			&book.PublishedDate,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}
	resp.Books = books
	return resp, nil
}

func (b bookServiceRepo) CreateSaleBook(ctx context.Context, in *pb.CreateSaleBookRequest) (*pb.SaleBook, error) {
	query := `insert into user_books(id, book_id, user_id)
values ($1, $2, $3) returning id`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	book, err := b.GetBook(ctx, &pb.ByIdReq{Id: in.BookId})
	if book.Count <= 0 {
		var (
			firstName, lastName string
		)
		queryU := `select first_name, last_name from users where id=(select user_id from user_books where book_id=$1 limit 1);`
		err = b.db.QueryRow(ctx, queryU, in.BookId).Scan(&firstName, &lastName)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("book %s is not available. Last book bought from %s %s", book.Name, firstName, lastName)

	}
	book.Count -= 1

	_, err = b.UpdateBook(ctx,
		&pb.UpdateBookRequest{
			Id:            book.Id,
			Name:          book.Name,
			Count:         book.Count,
			Author:        book.Author,
			PublishedDate: book.PublishedDate,
		})
	if err != nil {
		return nil, err
	}
	_, err = b.db.Exec(ctx, query,
		uuid,
		in.BookId,
		in.UserId,
	)
	if err != nil {
		return nil, err
	}
	newSaleBook, err := b.GetSaleBook(ctx, &pb.ByIdReq{Id: uuid.String()})
	if err != nil {
		return nil, err
	}

	return newSaleBook, nil
}

func (b bookServiceRepo) GetSaleBook(ctx context.Context, in *pb.ByIdReq) (*pb.SaleBook, error) {
	query := `SELECT * FROM user_books WHERE id = $1`
	res := &pb.SaleBook{}
	fmt.Println("DO")
	err := b.db.QueryRow(ctx, query, in.Id).Scan(
		&res.Id,
		&res.BookId,
		&res.UserId,
	)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b bookServiceRepo) UpdateSaleBook(ctx context.Context, in *pb.UpdateSaleBookRequest) (*pb.SaleBook, error) {
	query := `UPDATE user_books SET book_id = $1, user_id = $2 WHERE id = $3 returning id`
	result, err := b.db.Exec(ctx, query,
		in.BookId,
		in.UserId,
		in.Id,
	)
	if err != nil {
		return nil, err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	updatedSaleBook, err := b.GetSaleBook(ctx, &pb.ByIdReq{Id: in.Id})
	if err != nil {
		return nil, err
	}
	return updatedSaleBook, nil
}

func (b bookServiceRepo) DeleteSaleBook(ctx context.Context, in *pb.ByIdReq) (*emptypb.Empty, error) {
	query := `DELETE FROM user_books WHERE id = $1`
	result, err := b.db.Exec(ctx, query, in.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &emptypb.Empty{}, nil
}

func (b bookServiceRepo) ListSaleBooks(ctx context.Context, in *pb.ListSaleBooksRequest) (*pb.ListSaleBooksResponse, error) {
	resp := &pb.ListSaleBooksResponse{}
	query := `SELECT * FROM user_books limit $1 offset $2`
	rows, err := b.db.Query(ctx, query, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	var saleBooks []*pb.SaleBook

	for rows.Next() {
		var saleBook pb.SaleBook
		err = rows.Scan(
			&saleBook.Id,
			&saleBook.BookId,
			&saleBook.UserId,
		)
		if err != nil {
			return nil, err
		}
		saleBooks = append(saleBooks, &saleBook)
	}
	resp.SaleBooks = saleBooks

	return resp, nil
}
