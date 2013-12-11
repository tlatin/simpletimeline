package event

import (
	"appengine"
	"net/http"
	"timeline"
)

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	applicationKeyStr := r.FormValue("applicationKey")
	applicationKey, err := Timeline.GetApplicationKeyByString(c, applicationKeyStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := Timeline.NewEvent(c, applicationKey, r.FormValue("authorId"), r.FormValue("content")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/application/"+applicationKeyStr, http.StatusFound)
}
