package sqlite3

import (
	"context"
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"go-one-auth/internal/storage/service"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Sqlite3 struct {
	*sql.DB
}

const path = "file://migrations/sqlite3"

func New(cred string) Sqlite3 {
	open, err := sql.Open("sqlite3", cred)
	if err != nil {
		panic(err)
	}

	driver, err := sqlite.WithInstance(open, &sqlite.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		path,
		"sqlite", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	return Sqlite3{DB: open}
}

func (s Sqlite3) Close() error {
	return s.DB.Close()
}

func (s Sqlite3) Ping() error {
	return s.DB.Ping()
}

func (s Sqlite3) GetOne(ctx context.Context, key string) (string, error) {
	var value string
	err := s.QueryRowContext(ctx, "SELECT value FROM key_value WHERE key = ?", key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s Sqlite3) SetOne(ctx context.Context, key, value string) error {
	_, err := s.ExecContext(ctx, "INSERT INTO key_value (key, value) VALUES (?, ?)", key, value)
	if err != nil {
		if c, ok := err.(sqlite3.Error); ok && c.Code == sqlite3.ErrConstraint {
			return service.ErrPairExists
		}
		return err
	}

	return nil
}
