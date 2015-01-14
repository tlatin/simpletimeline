package event

import (
	"appengine"
	"github.com/tlatin/simpletimeline/timeline"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	applicationKeyStr := r.FormValue("applicationKey")
	applicationKey, err := Timeline.GetApplicationKeyByString(c, applicationKeyStr)
	if err != nil {
		c.Errorf("failed to find the application key.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := Timeline.NewEvent(c, applicationKey, r.FormValue("authorId"), r.FormValue("content")); err != nil {
		c.Errorf("failed to create a new event.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: Don't redirect.
	http.Redirect(w, r, "/application/"+applicationKeyStr, http.StatusFound)
}
