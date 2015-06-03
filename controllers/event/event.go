package event

import (
	"appengine"
	"github.com/tlatin/simpletimeline/timeline"
	"github.com/tlatin/simpletimeline/utils"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	formRedirect := r.FormValue("formRedirect")
	applicationKeyStr := r.FormValue("applicationKey")
	applicationKey, err := Timeline.GetApplicationKeyByString(c, applicationKeyStr)
	if utils.CheckHandlerError(c, err, w, "failed to find the application key.") {
		return
	}
	_, err = Timeline.NewEvent(c, applicationKey, r.FormValue("authorId"), r.FormValue("content"))
	if utils.CheckHandlerError(c, err, w, "failed to create a new event.") {
		return
	}
	// TODO: Don't redirect.
	if formRedirect != "" {
		http.Redirect(w, r, "/application/"+applicationKeyStr, http.StatusFound)
	} else {
		w.WriteHeader(200)
	}

}
