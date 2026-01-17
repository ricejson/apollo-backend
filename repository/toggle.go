package repository

import (
	"context"
	"github.com/ricejson/apollo-backend/domain"
	"github.com/ricejson/apollo-backend/repository/dao"
)

type DefaultToggleRepository struct {
	toggleDAO dao.ToggleDAO
}

func NewDefaultToggleRepository(toggleDAO dao.ToggleDAO) *DefaultToggleRepository {
	return &DefaultToggleRepository{
		toggleDAO: toggleDAO,
	}
}

func (repo *DefaultToggleRepository) FindAll(ctx context.Context) ([]domain.Toggle, error) {
	toggles, err := repo.toggleDAO.List(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]domain.Toggle, len(toggles))
	for i, toggle := range toggles {
		results[i] = domain.ToggleDao2Domain(toggle)
	}
	return results, nil

}
