package application

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
	"timeline"
)


var newApplicationTemplate = template.Must(template.ParseFiles("controllers/templates/results.html"))
var applicationsTemplate = template.Must(template.ParseFiles("controllers/templates/applications.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Application").Limit(10)
	applications := make([]Timeline.Application, 0, 10)
	if _, err := q.GetAll(c, &applications); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := applicationsTemplate.Execute(w, applications); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	name := r.FormValue("name")
	url := r.FormValue("url")
	// Validate URL
	// Make sure it is absolute
	if _, err := newApplication(name, url, c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := newApplicationTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newApplication(name string, url string, c appengine.Context) (key datastore.Key, err error) {
	app := Timeline.Application{
		Name: name,
		Url: url,
	}
	_, err = datastore.Put(c, datastore.NewIncompleteKey(c, "Application", nil), &app)
	if err != nil {
		return key, err
	}
	return
}
