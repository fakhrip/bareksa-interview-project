package routes

import (
	request "bareksa-interview-project/util/requesttype"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func Initialize() []request.Request {
	allRequests := make([]request.Request, 0, 20)

	allRequests = append(allRequests,
		request.AddRequest(request.GET, "/health_check", func(w http.ResponseWriter, req bunrouter.Request) error {
			return bunrouter.JSON(w, bunrouter.H{
				"message": "🤖: Ayy sir, service is currently healthy, you may want to continue enjoy your life now",
			})
		}, "Check backend service health (basically a status check)"))

	return allRequests
}

func ApiRoutes() (func(g *bunrouter.Group), *[]request.Request) {
	allRequests := Initialize()

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
