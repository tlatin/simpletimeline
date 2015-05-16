package search

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"github.com/tlatin/simpletimeline/timeline"
	"github.com/tlatin/simpletimeline/utils"
	"html/template"
	"net/http"
	"sort"
)

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	searchTemplate := template.Must(
		template.ParseFiles(
			utils.GetTemplatePath() + "search_query.html",
			utils.GetTemplatePath() + "search_query_form.html"))

	if utils.CheckHandlerError(c, searchTemplate.Execute(w, nil), w, "failed to render template.") {
		return
	}
}

func Post(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	c.Infof("Call to search to log")

	applicationKeyStr := r.FormValue("applicationKey")
	applicationKey, err := datastore.DecodeKey(applicationKeyStr)
	if utils.CheckHandlerError(c, err, w, "Searching for posts by failed to decode applicationKey: " + applicationKeyStr) {
		return
	}

	query := r.FormValue("query")
	AuthorIds := make([]string, 0, 10)
	if utils.CheckHandlerError(c, json.Unmarshal([]byte(query), &AuthorIds), w, "Searching for posts by failed to unmarshall json:" + query) {
		return
	}

	TimelineEvents := make([]Timeline.Event, 0, 10)
	for _, authorId := range AuthorIds {
		c.Infof("Searching for posts by " + authorId)
		q := datastore.NewQuery("TimelineEvent").Ancestor(applicationKey).Filter("AuthorId =", authorId).Limit(25)
		_, err := q.GetAll(c, &TimelineEvents)
		if utils.CheckHandlerError(c, err, w, "GetAll query failed.") {
			return
		}
	}
	sort.Sort(Timeline.ByDate(TimelineEvents))

	var newSearchQueryTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath() + "results.json"))
	if utils.CheckHandlerError(c, newSearchQueryTemplate.Execute(w, TimelineEvents), w, "failed to render search template") {
		return
	}
}
