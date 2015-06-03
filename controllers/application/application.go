package application

import (
	"appengine"
	"errors"
	"github.com/tlatin/simpletimeline/timeline"
	"github.com/tlatin/simpletimeline/utils"
	"html/template"
	"net/http"
	"regexp"
)

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Is there a better way to declare a string, not an empty string pointer?
	applicationKey, err := parseApplicationKeyFromURL(r.URL.Path)
	if utils.CheckHandlerError(c, err, w, "failed parse Application Key from URL.") {
		return
	}

	application, err := Timeline.GetApplicationByEncodedKey(c, applicationKey)
	if utils.CheckHandlerError(c, err, w, "The application key failed to return an Application object.") {
		return
	}
	applicationsTemplate := template.Must(
		template.ParseFiles(
			utils.GetTemplatePath()+"applications.html",
			utils.GetTemplatePath()+"new_event_form.html",
			utils.GetTemplatePath()+"search_query.html",
			utils.GetTemplatePath()+"search_query_form.html"))
	if utils.CheckHandlerError(c, applicationsTemplate.Execute(w, application), w, "failed to render template.") {
		return
	}
}

func parseApplicationKeyFromURL(s string) (applicationKey string, err error) {
	re := regexp.MustCompile("^/application/([^/]+)$")
	if matches := re.FindStringSubmatch(s); len(matches) != 2 {
		err = errors.New("Invalid application key")
	} else {
		applicationKey = matches[1]
	}
	return applicationKey, err
}

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	name := r.FormValue("name")
	url := r.FormValue("url")
	// Validate URL
	// Make sure it is absolute
	if _, err := Timeline.NewApplication(c, name, url); utils.CheckHandlerError(c, err, w, "failed to create a new appliation.") {
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
