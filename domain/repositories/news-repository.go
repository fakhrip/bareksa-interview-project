package repositories

import (
	domain "bareksa-interview-project/domain"
	"context"
)

type INewsRepository interface {
	FindOneById(ctx context.Context, id int64) (*domain.News, error)
	GetAll(ctx context.Context) ([]domain.News, error)
	Insert(ctx context.Context, news *domain.News) error
	Update(ctx context.Context, news *domain.News) error
	Delete(ctx context.Context, id int64) error
}
