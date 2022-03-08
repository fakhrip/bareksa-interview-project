package repositories

import (
	domain "bareksa-interview-project/domain"
	"context"
)

type ITopicsRepository interface {
	FindOneByColumn(ctx context.Context, col string, query interface{}) (*domain.Topics, error)
	FindAllByColumn(ctx context.Context, col string, query interface{}) ([]domain.Topics, error)
	GetAll(ctx context.Context) ([]domain.Topics, error)
	Insert(ctx context.Context, topic *domain.Topics) (*domain.Topics, error)
	Update(ctx context.Context, topic *domain.Topics, id int64) (*domain.Topics, error)
	Delete(ctx context.Context, id int64) error
}
