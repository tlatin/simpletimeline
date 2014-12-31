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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := Timeline.NewEvent(c, applicationKey, r.FormValue("authorId"), r.FormValue("content")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/application/"+applicationKeyStr, http.StatusFound)
}
