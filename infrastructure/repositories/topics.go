package repositories

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"
	"strconv"

	"github.com/uptrace/bun"
)

type topicsRepository struct {
	db *bun.DB
}

func createTopicsRepository(db *bun.DB) repositories.ITopicsRepository {
	return &topicsRepository{db: db}
}

func (repository *topicsRepository) FindOneByColumn(ctx context.Context, col string, query interface{}) (*domain.Topics, error) {
	topics := new(domain.Topics)

	err := repository.db.NewSelect().Model(topics).
		Where("? = ?", bun.Ident(col), query).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (repository *topicsRepository) FindAllByColumn(ctx context.Context, col string, query interface{}) ([]domain.Topics, error) {
	someTopics := make([]domain.Topics, 0)

	err := repository.db.NewSelect().Model(&someTopics).
		Where("? = ?", bun.Ident(col), query).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return someTopics, nil
}

func (repository *topicsRepository) GetAll(ctx context.Context) ([]domain.Topics, error) {
	allTopics := make([]domain.Topics, 0)

	err := repository.db.NewSelect().Model(&allTopics).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return allTopics, nil
}

func (repository *topicsRepository) Insert(ctx context.Context, topic *domain.Topics) error {
	_, err := repository.db.NewInsert().Model(topic).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repository *topicsRepository) Update(ctx context.Context, topic *domain.Topics, id int64) error {
	_, err := repository.db.NewUpdate().Model(topic).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repository *topicsRepository) Delete(ctx context.Context, id int64) error {
	_, err := repository.db.NewDelete().Model((*domain.Topics)(nil)).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
