package repositories

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"
	"fmt"
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

	cacheKey := fmt.Sprintf("news-%s-%v", col, query)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &news); err == nil {
			return news, nil
		} else {
			return nil, err
		}
	}

	err := repository.db.NewSelect().Model(news).
		Where("? = ?", bun.Ident(col), query).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if err = repository.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: news,
	}); err != nil {
		return nil, err
	}

	return news, nil
}

func (repository *newsRepository) FindAllByColumn(ctx context.Context, col string, query interface{}) ([]domain.News, error) {
	someNews := make([]domain.News, 0)

	cacheKey := fmt.Sprintf("news-%s-%v-all", col, query)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &someNews); err == nil {
			return someNews, nil
		} else {
			return nil, err
		}
	}

	err := repository.db.NewSelect().Model(&someNews).
		Where("? = ?", bun.Ident(col), query).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if err = repository.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: someNews,
	}); err != nil {
		return nil, err
	}

	return someNews, nil
}

func (repository *newsRepository) GetAll(ctx context.Context) ([]domain.News, error) {
	allNews := make([]domain.News, 0)

	cacheKey := "news-all"
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &allNews); err == nil {
			return allNews, nil
		} else {
			return nil, err
		}
	}

	err := repository.db.NewSelect().Model(&allNews).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if err = repository.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: allNews,
	}); err != nil {
		return nil, err
	}

	return allNews, nil
}

func (repository *newsRepository) Insert(ctx context.Context, news *domain.News) (*domain.News, error) {
	allNews := make([]domain.News, 0)

	res, err := repository.db.NewInsert().Model(news).Exec(ctx)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	news.ID = id

	cacheKey := "news-all"
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &allNews); err != nil {
			return nil, err
		}

		allNews = append(allNews, *news)
		if err := repository.cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   cacheKey,
			Value: allNews,
		}); err != nil {
			return nil, err
		}
	}

	return news, nil
}

func (repository *newsRepository) Update(ctx context.Context, news *domain.News, id int64) (*domain.News, error) {
	res, err := repository.db.NewUpdate().Model(news).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return nil, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	news.ID = id

	cacheKey := fmt.Sprintf("news-id-%v", id)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   cacheKey,
			Value: news,
		}); err != nil {
			return nil, err
		}
	}

	return news, nil
}

func (repository *newsRepository) Delete(ctx context.Context, id int64) error {
	cacheKey := fmt.Sprintf("news-id-%v", id)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Delete(ctx, cacheKey); err != nil {
			return err
		}
	}

	_, err := repository.db.NewDelete().Model((*domain.News)(nil)).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
