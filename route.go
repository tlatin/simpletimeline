package route

import (
	"controllers/application"
	"controllers/event"
	"controllers/index"
	"controllers/search"
	"net/http"
)

func init() {
	http.HandleFunc("/", index.Get)
	http.HandleFunc("/search", search.Get)
	http.HandleFunc("/search/new", search.Post)
	http.HandleFunc("/event/new", event.Post)
	http.HandleFunc("/application/", application.Get)
	http.HandleFunc("/application/new", application.Post)
}
