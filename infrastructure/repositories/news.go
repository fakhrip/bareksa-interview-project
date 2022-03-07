package repositories

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"
	"strconv"

	"github.com/go-redis/cache/v8"
	"github.com/uptrace/bun"
)

type newsRepository struct {
	db    *bun.DB
	cache *cache.Cache
}

func createNewsRepository(db *bun.DB, cache *cache.Cache) repositories.INewsRepository {
	return &newsRepository{db: db, cache: cache}
}

func (repository *newsRepository) FindOneByColumn(ctx context.Context, col string, query interface{}) (*domain.News, error) {
	news := new(domain.News)

	err := repository.db.NewSelect().Model(news).
		Where("? = ?", bun.Ident(col), query).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (repository *newsRepository) FindAllByColumn(ctx context.Context, col string, query interface{}) ([]domain.News, error) {
	someNews := make([]domain.News, 0)

	err := repository.db.NewSelect().Model(&someNews).
		Where("? = ?", bun.Ident(col), query).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return someNews, nil
}

func (repository *newsRepository) GetAll(ctx context.Context) ([]domain.News, error) {
	allNews := make([]domain.News, 0)

	err := repository.db.NewSelect().Model(&allNews).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return allNews, nil
}

func (repository *newsRepository) Insert(ctx context.Context, news *domain.News) error {
	_, err := repository.db.NewInsert().Model(news).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repository *newsRepository) Update(ctx context.Context, news *domain.News, id int64) error {
	_, err := repository.db.NewUpdate().Model(news).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repository *newsRepository) Delete(ctx context.Context, id int64) error {
	_, err := repository.db.NewDelete().Model((*domain.News)(nil)).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
