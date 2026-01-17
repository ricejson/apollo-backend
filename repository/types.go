package repository

import (
	"context"
	"github.com/ricejson/apollo-backend/domain"
)

type ToggleRepository interface {
	FindAll(ctx context.Context) ([]domain.Toggle, error)
}
