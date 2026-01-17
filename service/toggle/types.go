package toggle

import (
	"context"
	"github.com/ricejson/apollo-backend/domain"
)

// ToggleService 开关服务
type ToggleService interface {
	FindAll(ctx context.Context) ([]domain.Toggle, error)
}
