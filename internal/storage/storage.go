package storage

import (
	"context"
	"go-one-auth/internal/storage/grpcis"
	"go-one-auth/internal/storage/sqlite3"
)

type db interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
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

		return &Storage{
			conn: grpcis.New(cfg.DataSourceCred),
		}
	}
	panic("unknown storage type")
	return nil
}

func (s *Storage) CreatePair(ctx context.Context, username, password string) error {
	return s.conn.Set(ctx, username, password)
}

func (s *Storage) GetPair(ctx context.Context, username string) (string, error) {
	return s.conn.Get(ctx, username)
}
