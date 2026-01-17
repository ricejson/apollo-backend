package toggle

import (
	"context"
	"github.com/ricejson/apollo-backend/domain"
	"github.com/ricejson/apollo-backend/repository"
)

// HTTPToggleService HTTP实现
type HTTPToggleService struct {
	toggleRepo repository.ToggleRepository
}

func NewHTTPToggleService(toggleRepo repository.ToggleRepository) *HTTPToggleService {
	return &HTTPToggleService{toggleRepo: toggleRepo}
}

func (s *HTTPToggleService) FindAll(ctx context.Context) ([]domain.Toggle, error) {
	return s.toggleRepo.FindAll(ctx)
}
