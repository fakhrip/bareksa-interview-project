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

func CreateTopics(topicsService application.TopicsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			topics *domain.Topics
			err    error
		)

		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&topics); err != nil {
			return err
		}

		if topics, err = topicsService.InsertTopics(context.Background(), topics); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "Topics successfully created",
			"topics":  topics,
		})
	}
}

func ReadAllTopics(topicsService application.TopicsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			topics []domain.Topics
			err    error
		)

		if topics, err = topicsService.GetAllTopics(context.Background()); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "Topics successfully read",
			"topics":  topics,
		})
	}
}

func ReadTopicsById(topicsService application.TopicsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			topics *domain.Topics
			err    error
		)

		id, err := strconv.Atoi(req.Params().ByName("id"))
		if err != nil {
			return err
		}

		if topics, err = topicsService.GetTopicsById(context.Background(), int64(id)); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "Topics successfully read",
			"topics":  topics,
		})
	}
}

func UpdateTopicsById(topicsService application.TopicsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			topics *domain.Topics
			err    error
		)

		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&topics); err != nil {
			return err
		}

		id, err := strconv.Atoi(req.Params().ByName("id"))
		if err != nil {
			return err
		}

		if topics, err = topicsService.UpdateTopics(context.Background(), topics, int64(id)); err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message": "Topics successfully updated",
			"topics":  topics,
		})
	}
}

func DeleteTopicsById(topicsService application.TopicsService) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		var (
			id  interface{}
			err error
		)

		idReq, err := strconv.Atoi(req.Params().ByName("id"))
		if err != nil {
			return err
		}

		id, err = topicsService.DeleteTopics(context.Background(), int64(idReq))
		if err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"message":  "Topics successfully read",
			"topicsId": id,
		})
	}
}
