package repositories

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"

	"github.com/uptrace/bun"
)

type topicsRepository struct {
	db *bun.DB
}

func newTopicsRepository(db *bun.DB) repositories.ITopicsRepository {
	return &topicsRepository{db: db}
}

func (repository *topicsRepository) FindOneById(ctx context.Context, id int64) (*domain.Topics, error) {
	// TODO: implement this function
	return nil, nil
}

func (repository *topicsRepository) GetAll(ctx context.Context) ([]domain.Topics, error) {
	// TODO: implement this function
	return nil, nil
}

func (repository *topicsRepository) Insert(ctx context.Context, topic *domain.Topics) error {
	// TODO: implement this function
	return nil
}

func (repository *topicsRepository) Update(ctx context.Context, topic *domain.Topics) error {
	// TODO: implement this function
	return nil
}

func (repository *topicsRepository) Delete(ctx context.Context, id int64) error {
	// TODO: implement this function
	return nil
}
