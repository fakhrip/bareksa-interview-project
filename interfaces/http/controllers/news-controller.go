package controllers

import (
	application "bareksa-interview-project/application"
	domain "bareksa-interview-project/domain"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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

func ReadAllNews(newsService application.NewsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			news []domain.News
			err  error
		)

		if news, err = newsService.GetAllNews(context.Background()); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "News successfully read",
			"news":    news,
		})
	}
}

func ReadNewsById(newsService application.NewsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			news *domain.News
			err  error
		)

		id, err := strconv.Atoi(req.Params().ByName("id"))
		if err != nil {
			return err
		}

		if news, err = newsService.GetNewsById(context.Background(), int64(id)); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "News successfully read",
			"news":    news,
		})
	}
}

func UpdateNewsById(newsService application.NewsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			news *domain.News
			err  error
		)

		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&news); err != nil {
			return err
		}

		id, err := strconv.Atoi(req.Params().ByName("id"))
		if err != nil {
			return err
		}

		if news, err = newsService.UpdateNews(context.Background(), news, int64(id)); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "News successfully updated",
			"news":    news,
		})
	}
}

func DeleteNewsById(newsService application.NewsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			id  interface{}
			err error
		)

		idReq, err := strconv.Atoi(req.Params().ByName("id"))
		if err != nil {
			return err
		}

		id, err = newsService.DeleteNews(context.Background(), int64(idReq))
		if err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "News successfully read",
			"newsId":  id,
		})
	}
}
