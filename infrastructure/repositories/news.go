package repositories

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"

	"github.com/uptrace/bun"
)

type newsRepository struct {
	db *bun.DB
}

func newNewsRepository(db *bun.DB) repositories.INewsRepository {
	return &newsRepository{db: db}
}

func (repository *newsRepository) FindOneById(ctx context.Context, id int64) (*domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (repository *newsRepository) GetAll(ctx context.Context) ([]domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (repository *newsRepository) Insert(ctx context.Context, news *domain.News) error {
	// TODO: implement this function
	return nil
}

func (repository *newsRepository) Update(ctx context.Context, news *domain.News) error {
	// TODO: implement this function
	return nil
}

func (repository *newsRepository) Delete(ctx context.Context, id int64) error {
	// TODO: implement this function
	return nil
}
