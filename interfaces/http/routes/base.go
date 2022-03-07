package routes

import (
	request "bareksa-interview-project/util/requesttype"
	"html/template"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func BaseRoutes(apiRequests *[]request.Request) func(g *bunrouter.Group) {
	indexTmpl := `
		<html>
		<head>
			<title>
				Bareksa Interview Project (auto-generated documentation)
			</title>
		</head>
		<body>
			<h1>Bareksa Interview Project (auto-generated documentation)</h1>
			<h3>Api endpoint list :</h3>
			<ul>
				{{range .Apis}}
					<li>/api/v1{{.Path}} [{{.TypeString}}]: {{.Description}}</li>
				{{end}}
			</ul>
		</body>
		</html>
	`

	return func(group *bunrouter.Group) {
		group.GET("", func(w http.ResponseWriter, req bunrouter.Request) error {
			m := map[string]interface{}{
				"Apis": apiRequests,
			}
			return indexTemplate(indexTmpl).Execute(w, m)
		})
	}
}

func indexTemplate(templateString string) *template.Template {
	return template.Must(template.New("index").Parse(templateString))
}
