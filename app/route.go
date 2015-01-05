package route

import (
	"github.com/tlatin/simpletimeline/controllers/application"
	"github.com/tlatin/simpletimeline/controllers/event"
	"github.com/tlatin/simpletimeline/controllers/index"
	"github.com/tlatin/simpletimeline/controllers/search"
	"github.com/tlatin/simpletimeline/controllers/cron"
	"net/http"
)

func init() {
	http.HandleFunc("/", index.Get)
	http.HandleFunc("/application/", application.Get)
	http.HandleFunc("/application/new", application.Post)
	http.HandleFunc("/cron", cron.Get)
	http.HandleFunc("/search", search.Get)
	http.HandleFunc("/search/new", search.Post)
	http.HandleFunc("/event/new", event.Post)
}
