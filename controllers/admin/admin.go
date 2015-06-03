package admin

// Not currently being used.

import (
	"appengine"
	"github.com/tlatin/simpletimeline/utils"
	"html/template"
	"net/http"
)

var adminTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath()+"admin.html", utils.GetTemplatePath()+"new_application_form.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if utils.CheckHandlerError(c, adminTemplate.Execute(w, nil), w, "admin template failed to load.") {
		return
	}
}
