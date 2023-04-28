package grpcis

import (
	"context"
	"errors"
	"github.com/egorgasay/grpcis-go-sdk"
	"go-one-auth/internal/storage/service"
)

type GRPCis struct {
	*grpcis.Client
}

const path = "file://migrations/sqlite3"

func New(cred string) GRPCis {
	cl, err := grpcis.New(cred)
	if err != nil {
		panic(err)
	}

	return GRPCis{Client: cl}
}

func (g GRPCis) Get(ctx context.Context, key string) (string, error) {
	return g.GetOne(ctx, key)
}

func (g GRPCis) Set(ctx context.Context, key, value string) error {
	if err := g.SetOne(ctx, key, value, true); err != nil {
		if errors.Is(err, grpcis.ErrUniqueConstraint) {
			return service.ErrPairExists
		}
		return err
	}

	return nil
}
