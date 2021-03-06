package application

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"
)

type (
	TopicsService interface {
		GetTopicsById(ctx context.Context, id int64) (*domain.Topics, error)
		GetNewsByTopics(ctx context.Context, topics string) ([]domain.News, error)
		GetAllTopics(ctx context.Context) ([]domain.Topics, error)
		InsertTopics(ctx context.Context, newTopicsModel *domain.Topics) (*domain.Topics, error)
		UpdateTopics(ctx context.Context, newTopicsModel *domain.Topics, id int64) (*domain.Topics, error)
		DeleteTopics(ctx context.Context, id int64) (interface{}, error)
	}
	topicsService struct {
		Repository repositories.ITopicsRepository
	}
)

func CreateTopicsService(repository repositories.ITopicsRepository) TopicsService {
	return &topicsService{Repository: repository}
}

func (service *topicsService) GetTopicsById(ctx context.Context, id int64) (*domain.Topics, error) {
	var (
		topics *domain.Topics
		err    error
	)

	if topics, err = service.Repository.FindOneByColumn(ctx, "id", id); err != nil {
		return nil, err
	}

	return topics, nil
}

func (service *topicsService) GetNewsByTopics(ctx context.Context, topics string) ([]domain.News, error) {
	var (
		allNews   []domain.News
		allTopics []domain.Topics
		err       error
	)

	if allTopics, err = service.Repository.FindAllByColumn(ctx, "name", topics); err != nil {
		return nil, err
	}

	for _, value := range allTopics {
		allNews = append(allNews, value.News)
	}

	return allNews, nil
}

func (service *topicsService) GetAllTopics(ctx context.Context) ([]domain.Topics, error) {
	var (
		allTopics []domain.Topics
		err       error
	)

	if allTopics, err = service.Repository.GetAll(ctx); err != nil {
		return nil, err
	}

	return allTopics, nil
}

func (service *topicsService) InsertTopics(ctx context.Context, newTopicsModel *domain.Topics) (*domain.Topics, error) {
	res, err := service.Repository.Insert(ctx, newTopicsModel)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *topicsService) UpdateTopics(ctx context.Context, newTopicsModel *domain.Topics, id int64) (*domain.Topics, error) {
	res, err := service.Repository.Update(ctx, newTopicsModel, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *topicsService) DeleteTopics(ctx context.Context, id int64) (interface{}, error) {
	if err := service.Repository.Delete(ctx, id); err != nil {
		return nil, err
	}

	return id, nil
}
