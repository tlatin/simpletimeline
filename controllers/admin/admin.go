package admin

// Not currently being used.

import (
	"github.com/tlatin/simpletimeline/utils"
	"html/template"
	"net/http"
)

var adminTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath() + "admin.html", utils.GetTemplatePath() + "new_application_form.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	if err := adminTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
