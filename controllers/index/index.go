package index

import (
	"appengine"
	"appengine/datastore"
	"github.com/tlatin/simpletimeline/timeline"
	"github.com/tlatin/simpletimeline/utils"
	"html/template"
	"net/http"
)


func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	applications := make([]Timeline.Application, 0, 10)
	// Having this code in the Timeline package isn't working.
	// So keeping it here for now.
	q := datastore.NewQuery("Application").Limit(10)
	keys, err := q.GetAll(c, &applications)
	if utils.CheckHandlerError(c, err, w, "Query to get all applications failed.") {
		return
	}

	for i := range applications {
		applications[i].WebKey = keys[i].Encode()
	}

	timelineTemplate := template.Must(
		template.ParseFiles(
			utils.GetTemplatePath() + "index.html",
			utils.GetTemplatePath() + "new_application_form.html"))
	
	if utils.CheckHandlerError(c, timelineTemplate.Execute(w, applications), w, "failed to render index template.") {
		return
	}

}
