package routes

import (
	repositories "bareksa-interview-project/infrastructure/repositories"
	controllers "bareksa-interview-project/interfaces/http/controllers"
	request "bareksa-interview-project/util/requesttype"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func Initialize(dbPass string) []request.Request {
	allRequests := make([]request.Request, 0, 20)
	newsService := repositories.CreateNewsServiceResolve(dbPass)
	topicsService := repositories.CreateTopicsServiceResolve(dbPass)

	allRequests = append(allRequests,
		request.AddRequest(request.GET, "/health_check", func(w http.ResponseWriter, req bunrouter.Request) error {
			return bunrouter.JSON(w, bunrouter.H{
				"message": "🤖: Ayy sir, service is currently healthy, you may want to continue enjoy your life now",
			})
		}, "Check backend service health (basically a status check)"),

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
			"Delete topics by given id"))

	return allRequests
}

func ApiRoutes(dbPass string) (func(g *bunrouter.Group), *[]request.Request) {
	allRequests := Initialize(dbPass)

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
