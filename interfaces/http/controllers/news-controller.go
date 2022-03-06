package controllers

import (
	application "bareksa-interview-project/application"
	domain "bareksa-interview-project/domain"
	"context"
	"encoding/json"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func CreateNews(newsService application.NewsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			news *domain.News
			err  error
		)

		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&news); err != nil {
			return err
		}

		if news, err = newsService.InsertNews(context.Background(), news); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "News successfully created",
			"news":    news,
		})
	}
}
