package search

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"sort"
	"timeline"
)

var searchTemplate = template.Must(
	template.ParseFiles(
		"controllers/templates/search_query.html",
		"controllers/templates/search_query_form.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	if err := searchTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var newSearchQueryTemplate = template.Must(template.ParseFiles("controllers/templates/results.html"))

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	applicationId := r.FormValue("applicationId")
	applicationKey, err := datastore.DecodeKey(applicationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	query := r.FormValue("query")
	AuthorIds := make([]string, 0, 10)
	if err := json.Unmarshal([]byte(query), &AuthorIds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	TimelineEvents := make([]Timeline.Event, 0, 10)
	for _, authorId := range AuthorIds {
		log.Println("Searching for posts by " + authorId)
		q := datastore.NewQuery("TimelineEvent").Ancestor(applicationKey).Filter("AuthorId =", authorId).Limit(25)
		if _, err := q.GetAll(c, &TimelineEvents); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	sort.Sort(Timeline.ByDate(TimelineEvents))

	if err := newSearchQueryTemplate.Execute(w, TimelineEvents); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
