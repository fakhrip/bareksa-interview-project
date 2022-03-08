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

type topicsRepository struct {
	db    *bun.DB
	cache *cache.Cache
}

func createTopicsRepository(db *bun.DB, cache *cache.Cache) repositories.ITopicsRepository {
	return &topicsRepository{db: db, cache: cache}
}

func (repository *topicsRepository) FindOneByColumn(ctx context.Context, col string, query interface{}) (*domain.Topics, error) {
	topics := new(domain.Topics)

	cacheKey := fmt.Sprintf("topics-%s-%v", col, query)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &topics); err == nil {
			return topics, nil
		} else {
			return nil, err
		}
	}

	err := repository.db.NewSelect().Model(topics).Relation("News").
		Where("? = ?", bun.Ident(col), query).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if err = repository.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: topics,
	}); err != nil {
		return nil, err
	}

	return topics, nil
}

func (repository *topicsRepository) FindAllByColumn(ctx context.Context, col string, query interface{}) ([]domain.Topics, error) {
	someTopics := make([]domain.Topics, 0)

	cacheKey := fmt.Sprintf("topics-%s-%v-all", col, query)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &someTopics); err == nil {
			return someTopics, nil
		} else {
			return nil, err
		}
	}

	err := repository.db.NewSelect().Model(&someTopics).Relation("News").
		Where("? = ?", bun.Ident(col), query).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if err = repository.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: someTopics,
	}); err != nil {
		return nil, err
	}

	return someTopics, nil
}

func (repository *topicsRepository) GetAll(ctx context.Context) ([]domain.Topics, error) {
	allTopics := make([]domain.Topics, 0)

	cacheKey := "topics-all"
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &allTopics); err == nil {
			return allTopics, nil
		} else {
			return nil, err
		}
	}

	err := repository.db.NewSelect().Model(&allTopics).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if err = repository.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: allTopics,
	}); err != nil {
		return nil, err
	}

	return allTopics, nil
}

func (repository *topicsRepository) Insert(ctx context.Context, topic *domain.Topics) (*domain.Topics, error) {
	allTopics := make([]domain.Topics, 0)

	res, err := repository.db.NewInsert().Model(topic).Exec(ctx)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	topic.ID = id

	cacheKey := "topics-all"
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Get(ctx, cacheKey, &allTopics); err != nil {
			return nil, err
		}

		allTopics = append(allTopics, *topic)
		if err := repository.cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   cacheKey,
			Value: allTopics,
		}); err != nil {
			return nil, err
		}
	}

	return topic, err
}

func (repository *topicsRepository) Update(ctx context.Context, topic *domain.Topics, id int64) (*domain.Topics, error) {
	res, err := repository.db.NewUpdate().Model(topic).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return nil, err
	}

	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	topic.ID = id

	cacheKey := fmt.Sprintf("topics-id-%v", id)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   cacheKey,
			Value: topic,
		}); err != nil {
			return nil, err
		}
	}

	return topic, nil
}

func (repository *topicsRepository) Delete(ctx context.Context, id int64) error {
	cacheKey := fmt.Sprintf("topics-id-%v", id)
	if repository.cache.Exists(ctx, cacheKey) {
		if err := repository.cache.Delete(ctx, cacheKey); err != nil {
			return err
		}
	}

	_, err := repository.db.NewDelete().Model((*domain.Topics)(nil)).
		Where("? = ?", bun.Ident("id"), strconv.Itoa(int(id))).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
