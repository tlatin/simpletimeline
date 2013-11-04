package admin

// Not currently being used.

import (
	"html/template"
	"net/http"
)

var adminTemplate = template.Must(template.ParseFiles("controllers/templates/admin.html", "controllers/templates/new_application_form.html"))

func Get(w http.ResponseWriter, r *http.Request) {
	if err := adminTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
