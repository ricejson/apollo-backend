package dao

import (
	"context"
)

type ToggleDAO interface {
	List(context.Context) ([]Toggle, error)
	InsertOne(context.Context, Toggle) (bool, error)
}
