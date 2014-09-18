package index

import (
	"appengine"
	"appengine/datastore"
	"timeline"
	"html/template"
	"net/http"
)

var timelineTemplate = template.Must(
	template.ParseFiles(
		"controllers/templates/index.html",
		"controllers/templates/new_application_form.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	applications := make([]Timeline.Application, 0, 10)
	// Having this code in the Timeline package isn't working.
	// So keeping it here for now.
	q := datastore.NewQuery("Application").Limit(10)
	keys, err := q.GetAll(c, &applications)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range applications {
		applications[i].WebKey = keys[i].Encode()
	}

	// if err := Timeline.GetAllApplications(c, applications, 10); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	if err := timelineTemplate.Execute(w, applications); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
