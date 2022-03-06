package routes

import (
	domain "bareksa-interview-project/domain"
	persistence "bareksa-interview-project/infrastructure/persistence"
	repositories "bareksa-interview-project/infrastructure/repositories"
	controllers "bareksa-interview-project/interfaces/http/controllers"
	request "bareksa-interview-project/util/requesttype"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func Initialize(dbPass string, migrationPass string) []request.Request {
	allRequests := make([]request.Request, 0, 20)
	newsService := repositories.CreateNewsServiceResolve(dbPass)
	topicsService := repositories.CreateTopicsServiceResolve(dbPass)

	allRequests = append(allRequests,
		request.AddRequest(request.GET, "/health_check", func(w http.ResponseWriter, req bunrouter.Request) error {
			return bunrouter.JSON(w, bunrouter.H{
				"message": "🤖: Ayy sir, service is currently healthy, you may want to continue enjoy your life now",
			})
		}, "Check backend service health (basically a status check)"),

		request.AddRequest(request.POST, "/refresh_migration", func(w http.ResponseWriter, req bunrouter.Request) error {
			var (
				secretStruct struct {
					secret string
				}
				err error
			)

			decoder := json.NewDecoder(req.Body)
			if err := decoder.Decode(&secretStruct); err != nil {
				return err
			}

			if secretStruct.secret != migrationPass {
				return errors.New("The secret is wrong, dont try any harder if you are not the admin")
			}

			db := persistence.CreateDatabase(dbPass)
			err = db.ResetModel(req.Context(), (*domain.News)(nil), (*domain.Topics)(nil))
			if err != nil {
				panic(err)
			}

			return bunrouter.JSON(w, bunrouter.H{
				"message": "🤖: Ayy sir, database migration has been refreshed successfully",
			})
		}, "Refresh database migration (dropping and recreating all tables)"),

		request.AddRequest(request.POST, "/news", controllers.CreateNews(newsService),
			"Create given news"),

		request.AddRequest(request.GET, "/news", controllers.ReadAllNews(newsService),
			"Get list of all news"),

		request.AddRequest(request.GET, "/news/:id", controllers.ReadNewsById(newsService),
			"Get news from given id"),

		request.AddRequest(request.PUT, "/news/:id", controllers.UpdateNewsById(newsService),
			"Update news by given id"),

		request.AddRequest(request.DELETE, "/news/:id", controllers.DeleteNewsById(newsService),
			"Delete news by given id"),

		request.AddRequest(request.POST, "/topics", controllers.CreateTopics(topicsService),
			"Create given topics"),

		request.AddRequest(request.GET, "/topics", controllers.ReadAllTopics(topicsService),
			"Get list of all topics"),

		request.AddRequest(request.GET, "/topics/:id", controllers.ReadTopicsById(topicsService),
			"Get topics from given id"),

		request.AddRequest(request.PUT, "/topics/:id", controllers.UpdateTopicsById(topicsService),
			"Update topics by given id"),

		request.AddRequest(request.DELETE, "/topics/:id", controllers.DeleteTopicsById(topicsService),
			"Delete topics by given id"),

		request.AddRequest(request.GET, "/news/status/:status", controllers.ReadNewsByStatus(newsService),
			"Get news from given status"),

		request.AddRequest(request.GET, "/news/topic/:topic", controllers.ReadNewsByTopic(topicsService),
			"Get news from given topic"))

	return allRequests
}

func ApiRoutes(dbPass string, migrationPass string) (func(g *bunrouter.Group), *[]request.Request) {
	allRequests := Initialize(dbPass, migrationPass)

	return func(group *bunrouter.Group) {
		for index := range allRequests {
			switch allRequests[index].Type {
			case request.GET:
				group.GET(allRequests[index].Path, allRequests[index].Handler)
			case request.POST:
				group.POST(allRequests[index].Path, allRequests[index].Handler)
			case request.PUT:
				group.PUT(allRequests[index].Path, allRequests[index].Handler)
			case request.DELETE:
				group.DELETE(allRequests[index].Path, allRequests[index].Handler)
			}
		}
	}, &allRequests
}
