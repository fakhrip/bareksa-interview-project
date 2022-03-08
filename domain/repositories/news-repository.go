package repositories

import (
	domain "bareksa-interview-project/domain"
	"context"
)

type INewsRepository interface {
	FindOneByColumn(ctx context.Context, col string, query interface{}) (*domain.News, error)
	FindAllByColumn(ctx context.Context, col string, query interface{}) ([]domain.News, error)
	GetAll(ctx context.Context) ([]domain.News, error)
	Insert(ctx context.Context, news *domain.News) (*domain.News, error)
	Update(ctx context.Context, news *domain.News, id int64) (*domain.News, error)
	Delete(ctx context.Context, id int64) error
}
