package postgres

import (
	"context"
	"fmt"
	"github.com/husanmusa/pro-book-service/config"
	"github.com/husanmusa/pro-book-service/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db          *pgxpool.Pool
	bookService storage.BookServiceRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) BookService() storage.BookServiceRepoI {
	if s.bookService == nil {
		s.bookService = NewUserRepo(s.db)
	}

	return s.bookService
}

func (s *Store) ClientPlatform() storage.BookServiceRepoI {
	if s.bookService == nil {
		s.bookService = NewUserRepo(s.db)
	}

	return s.bookService
}
