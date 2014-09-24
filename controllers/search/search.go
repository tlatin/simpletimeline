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
	searchTemplate := template.Must(
		template.ParseFiles(
			utils.GetTemplatePath() + "search_query.html",
			utils.GetTemplatePath() + "search_query_form.html"))

	if err := searchTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func Post(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	c.Infof("Call to search to log")

	applicationId := r.FormValue("applicationId")
	applicationKey, err := datastore.DecodeKey(applicationId)
	if err != nil {
		c.Infof("Searching for posts by failed to decode applicationId: " + applicationId)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := r.FormValue("query")
	AuthorIds := make([]string, 0, 10)
	if err := json.Unmarshal([]byte(query), &AuthorIds); err != nil {
		c.Infof("Searching for posts by failed to unmarshall json:" + query)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	TimelineEvents := make([]Timeline.Event, 0, 10)
	for _, authorId := range AuthorIds {
		c.Infof("Searching for posts by " + authorId)
		q := datastore.NewQuery("TimelineEvent").Ancestor(applicationKey).Filter("AuthorId =", authorId).Limit(25)
		if _, err := q.GetAll(c, &TimelineEvents); err != nil {
			c.Infof("GetAll query failed.")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	sort.Sort(Timeline.ByDate(TimelineEvents))

	var newSearchQueryTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath() + "results.json"))
	if err := newSearchQueryTemplate.Execute(w, TimelineEvents); err != nil {
		c.Infof("failed to render search template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
