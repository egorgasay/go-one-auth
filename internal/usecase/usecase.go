package usecase

import (
	"context"
	"fmt"
	"go-one-auth/internal/storage"
)

type UseCase struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *UseCase {
	return &UseCase{
		storage: storage,
	}
}

func (uc *UseCase) CreatePair(context context.Context, username, password string) error {
	return uc.storage.CreatePair(context, username, password)
}

func (uc *UseCase) VerifyPair(context context.Context, username, password string) (bool, error) {
	value, err := uc.storage.GetPair(context, username)
	if err != nil {
		return false, fmt.Errorf("usecase: %w", err)
	}

	return value == password, nil
}
