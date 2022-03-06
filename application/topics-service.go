package application

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"
)

type (
	TopicsService interface {
		GetTopicsById(ctx context.Context, id int64) (*domain.Topics, error)
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
	if err := service.Repository.Insert(ctx, newTopicsModel); err != nil {
		return nil, err
	}

	return newTopicsModel, nil
}

func (service *topicsService) UpdateTopics(ctx context.Context, newTopicsModel *domain.Topics, id int64) (*domain.Topics, error) {
	if err := service.Repository.Update(ctx, newTopicsModel, id); err != nil {
		return nil, err
	}

	return newTopicsModel, nil
}

func (service *topicsService) DeleteTopics(ctx context.Context, id int64) (interface{}, error) {
	if err := service.Repository.Delete(ctx, id); err != nil {
		return nil, err
	}

	return id, nil
}
