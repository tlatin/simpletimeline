package application

import (
	"appengine"
	"html/template"
	"net/http"
	"regexp"
	"timeline"
)

var newApplicationTemplate = template.Must(template.ParseFiles("controllers/templates/results.html"))
var applicationsTemplate = template.Must(
	template.ParseFiles(
		"controllers/templates/applications.html",
		"controllers/templates/new_event_form.html",
		"controllers/templates/search_query.html",
		"controllers/templates/search_query_form.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Is there a better way to declare a string, not an empty string pointer?
	applicationKey := new(string)
	re := regexp.MustCompile("^/application/([^/]+)$")
	if matches := re.FindStringSubmatch(r.URL.Path); len(matches) != 2 {
		http.Error(w, "Invalid application key", http.StatusInternalServerError)
		return
	} else {
		applicationKey = &matches[1]
	}
	application, err := Timeline.GetApplicationById(c, *applicationKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := applicationsTemplate.Execute(w, application); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	name := r.FormValue("name")
	url := r.FormValue("url")
	// Validate URL
	// Make sure it is absolute
	if _, err := Timeline.NewApplication(name, url, c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := newApplicationTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
