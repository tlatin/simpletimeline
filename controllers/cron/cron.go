package cron

import (
	// "appengine"
	// "appengine/datastore"
	"html/template"
	"github.com/tlatin/simpletimeline/utils"
	"net/http"
)

var cronTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath() + "cron.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	// List all of the events older than month.
	if err := cronTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
		// q := datastore.NewQuery("TimelineEvent").Filter("Date <"time.Now() )
		// TimelineEvents := make([]Timeline.Event, 0, 10)
}
