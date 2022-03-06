package application

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	utilModel "bareksa-interview-project/util/model"
	"context"
	"errors"
)

type (
	NewsService interface {
		GetNewsById(ctx context.Context, id int64) (*domain.News, error)
		GetNewsByStatus(ctx context.Context, status int) ([]domain.News, error)
		GetNewsByTopics(ctx context.Context, topics string) ([]domain.News, error)
		GetAllNews(ctx context.Context) ([]domain.News, error)
		InsertNews(ctx context.Context, newNewsModel *domain.News) (*domain.News, error)
		UpdateNews(ctx context.Context, newNewsModel *domain.News, id int64) (*domain.News, error)
		DeleteNews(ctx context.Context, id int64) (interface{}, error)
	}
	newsService struct {
		Repository repositories.INewsRepository
	}
)

func CreateNewsService(repository repositories.INewsRepository) NewsService {
	return &newsService{Repository: repository}
}

func (service *newsService) GetNewsById(ctx context.Context, id int64) (*domain.News, error) {
	var (
		news *domain.News
		err  error
	)

	if news, err = service.Repository.FindOneByColumn(ctx, "id", id); err != nil {
		return nil, err
	}

	return news, nil
}

func (service *newsService) GetNewsByStatus(ctx context.Context, status int) ([]domain.News, error) {
	var (
		allNews []domain.News
		err     error
	)

	if allNews, err = service.Repository.FindAllByColumn(ctx, "status", utilModel.StatusString(status)); err != nil {
		return nil, err
	}

	return allNews, nil
}

func (service *newsService) GetNewsByTopics(ctx context.Context, topics string) ([]domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *newsService) GetAllNews(ctx context.Context) ([]domain.News, error) {
	var (
		allNews []domain.News
		err     error
	)

	if allNews, err = service.Repository.GetAll(ctx); err != nil {
		return nil, err
	}

	return allNews, nil
}

func (service *newsService) InsertNews(ctx context.Context, newNewsModel *domain.News) (*domain.News, error) {
	if _, ok := utilModel.StatusDict()[newNewsModel.Status]; !ok {
		return nil, errors.New("Status should be either 'draft', 'deleted', or 'publish'")
	}

	if err := service.Repository.Insert(ctx, newNewsModel); err != nil {
		return nil, err
	}

	return newNewsModel, nil
}

func (service *newsService) UpdateNews(ctx context.Context, newNewsModel *domain.News, id int64) (*domain.News, error) {
	if _, ok := utilModel.StatusDict()[newNewsModel.Status]; !ok {
		return nil, errors.New("Status should be either 'draft', 'deleted', or 'publish'")
	}

	if err := service.Repository.Update(ctx, newNewsModel, id); err != nil {
		return nil, err
	}

	return newNewsModel, nil
}

func (service *newsService) DeleteNews(ctx context.Context, id int64) (interface{}, error) {
	if err := service.Repository.Delete(ctx, id); err != nil {
		return nil, err
	}

	return id, nil
}
