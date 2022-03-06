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

func createNewsRepository(db *bun.DB) repositories.INewsRepository {
	return &newsRepository{db: db}
}

func (repository *newsRepository) FindOneByColumn(ctx context.Context, col string, query interface{}) (*domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (repository *newsRepository) FindAllByColumn(ctx context.Context, col string, query interface{}) ([]domain.News, error) {
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
