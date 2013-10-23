package event

import (
	"appengine"
	"appengine/datastore"
	"net/http"
	"time"
	"timeline"
)

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	applicationId := r.FormValue("applicationKey")
	// This looks like the wrong style. But inline app is scoped to the if block
	app, err := getApplication(applicationId, c)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return		
	}
	if err := newEvent(app, r.FormValue("authorId"), r.FormValue("content"), c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusFound)	
}

func getApplication(keyStr string, c appengine.Context) (app Timeline.Application, err error) {
	key, err := datastore.DecodeKey(keyStr)
	if err != nil {
		return
	}
	q := datastore.NewQuery("Application").Filter("__key__ = ", key).Limit(1)
	apps := make([]Timeline.Application, 0 , 1)
	if _, err := q.GetAll(c, &apps); err != nil {
		return app, err
	}
	app = apps[0]
	return
}

func newEvent(app Timeline.Application, authorId string, content string, c appengine.Context) error {
	te := Timeline.Event{
		Application:   app,
		AuthorId:      authorId,
		Content:       content,
		Date:          time.Now(),
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "TimelineEvent", nil), &te)
	if err != nil {		
		return err
	}
	return nil
}
