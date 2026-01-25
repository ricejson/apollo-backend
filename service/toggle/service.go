package toggle

import (
	"context"
	"github.com/ricejson/apollo-backend/domain"
	"github.com/ricejson/apollo-backend/repository"
)

// DefaultToggleService 默认实现
type DefaultToggleService struct {
	toggleRepo repository.ToggleRepository
}

func NewDefaultToggleService(toggleRepo repository.ToggleRepository) *DefaultToggleService {
	return &DefaultToggleService{toggleRepo: toggleRepo}
}

func (s *DefaultToggleService) FindAll(ctx context.Context) ([]domain.Toggle, error) {
	return s.toggleRepo.FindAll(ctx)
}

func (s *DefaultToggleService) InsertToggle(ctx context.Context, toggle domain.Toggle) (bool, error) {
	return s.toggleRepo.InsertOne(ctx, toggle)
}
