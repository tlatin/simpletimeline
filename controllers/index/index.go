package index

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
	"timeline"
)

var timelineTemplate = template.Must(
	template.ParseFiles(
		"controllers/templates/index.html", 
		"controllers/templates/search_query_form.html", 
		"controllers/templates/new_event_form.html",
		"controllers/templates/new_application_form.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("TimelineEvent").Order("-Date").Limit(10)
	TimelineEvents := make([]Timeline.Event, 0, 10)
	if _, err := q.GetAll(c, &TimelineEvents); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := timelineTemplate.Execute(w, TimelineEvents); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

