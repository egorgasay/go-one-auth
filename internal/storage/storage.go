package storage

import (
	"context"
	"github.com/egorgasay/grpcis-go-sdk"
	"go-one-auth/internal/storage/sqlite3"
)

type db interface {
	SetOne(ctx context.Context, key string, value string) error
	GetOne(ctx context.Context, key string) (string, error)
}

type Config struct {
	Type           string
	DataSourceCred string
}

const (
	Sqlite3 = "sqlite3"
	GRPCis  = "grpcis"
)

type Storage struct {
	conn db
}

func New(cfg *Config) *Storage {
	switch cfg.Type {
	case Sqlite3:
		return &Storage{
			conn: sqlite3.New(cfg.DataSourceCred),
		}
	case GRPCis:
		cl, err := grpcis.New(cfg.DataSourceCred)
		if err != nil {
			panic(err)
		}

		return &Storage{
			conn: cl,
		}
	}
	panic("unknown storage type")
	return nil
}

func (s *Storage) CreatePair(context context.Context, username, password string) error {
	return s.conn.SetOne(context, username, password)
}

func (s *Storage) GetPair(context context.Context, username string) (string, error) {
	return s.conn.GetOne(context, username)
}
