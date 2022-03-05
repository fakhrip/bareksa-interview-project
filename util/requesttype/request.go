package requesttype

import (
	"github.com/uptrace/bunrouter"
)

const (
	GET int = iota
	POST
	PUT
	DELETE
)

type Request struct {
	Type        int
	TypeString  string
	Path        string
	Handler     bunrouter.HandlerFunc
	Description string
}

func AddRequest(requestType int, path string, handler bunrouter.HandlerFunc, description string) Request {
	return Request{
		requestType,
		enumTypeToString(requestType),
		path,
		handler,
		description,
	}
}

func enumTypeToString(requestType int) string {
	switch requestType {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	}

	return ""
}
