package repositories

import (
	domain "bareksa-interview-project/domain"
	"context"
)

type ITopicsRepository interface {
	FindOneByColumn(ctx context.Context, col string, query interface{}) (*domain.Topics, error)
	GetAll(ctx context.Context) ([]domain.Topics, error)
	Insert(ctx context.Context, topic *domain.Topics) error
	Update(ctx context.Context, topic *domain.Topics) error
	Delete(ctx context.Context, id int64) error
}
