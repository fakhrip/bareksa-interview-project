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
		UpdateTopics(ctx context.Context, newTopicsModel *domain.Topics) (*domain.Topics, error)
		DeleteTopics(ctx context.Context, id int64) (*domain.Topics, error)
	}
	topicsService struct {
		Repository repositories.ITopicsRepository
	}
)

func CreateTopicsService(repository repositories.ITopicsRepository) TopicsService {
	return &topicsService{Repository: repository}
}

func (service *topicsService) GetTopicsById(ctx context.Context, id int64) (*domain.Topics, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *topicsService) GetAllTopics(ctx context.Context) ([]domain.Topics, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *topicsService) InsertTopics(ctx context.Context, newTopicsModel *domain.Topics) (*domain.Topics, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *topicsService) UpdateTopics(ctx context.Context, newTopicsModel *domain.Topics) (*domain.Topics, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *topicsService) DeleteTopics(ctx context.Context, id int64) (*domain.Topics, error) {
	// TODO: implement this function
	return nil, nil
}
