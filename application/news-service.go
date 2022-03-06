package application

import (
	domain "bareksa-interview-project/domain"
	repositories "bareksa-interview-project/domain/repositories"
	"context"
)

type (
	NewsService interface {
		GetNewsById(ctx context.Context, id int64) (*domain.News, error)
		GetNewsByStatus(ctx context.Context, status int) ([]domain.News, error)
		GetAllNews(ctx context.Context) ([]domain.News, error)
		InsertNews(ctx context.Context, newNewsModel *domain.News) (*domain.News, error)
		UpdateNews(ctx context.Context, newNewsModel *domain.News) (*domain.News, error)
		DeleteNews(ctx context.Context, id int64) (*domain.News, error)
	}
	newsService struct {
		Repository repositories.INewsRepository
	}
)

func CreateNewsService(repository repositories.INewsRepository) NewsService {
	return &newsService{Repository: repository}
}

func (service *newsService) GetNewsById(ctx context.Context, id int64) (*domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *newsService) GetNewsByStatus(ctx context.Context, status int) ([]domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *newsService) GetAllNews(ctx context.Context) ([]domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *newsService) InsertNews(ctx context.Context, newNewsModel *domain.News) (*domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *newsService) UpdateNews(ctx context.Context, newNewsModel *domain.News) (*domain.News, error) {
	// TODO: implement this function
	return nil, nil
}

func (service *newsService) DeleteNews(ctx context.Context, id int64) (*domain.News, error) {
	// TODO: implement this function
	return nil, nil
}
