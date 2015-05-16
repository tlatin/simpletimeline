package cron

import (
	"appengine"
	"appengine/datastore"
	"github.com/tlatin/simpletimeline/utils"
	"html/template"
	"net/http"
	"time"
)

var cronTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath() + "cron.html"))
var timeWindow = time.Hour * 24 * 30 * 3 // Look at all events more than 90 days old

type CronTemplateValues struct {
	Keys []*datastore.Key
	TimeWindow time.Duration
}

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	oldTimelineEvents, err := getEventsToDelete(c, timeWindow);
	if utils.CheckHandlerError(c, err, w, "failed to load all timeline events.") {
		return
	}

	templateValues := CronTemplateValues{Keys: oldTimelineEvents, TimeWindow:timeWindow}
	if utils.CheckHandlerError(c, cronTemplate.Execute(w, templateValues), w, "cron template failed to load.") {
		return
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	eventsArray, err := getEventsToDelete(c, timeWindow);
	if utils.CheckHandlerError(c, err, w, "failed to load all timeline events.") {
		return
	}

	for _, eventKey := range eventsArray {
		if utils.CheckHandlerError(c, datastore.Delete(c, eventKey), w, "event datastore failed to delete.") {
			return
		}
	}
	
	http.Redirect(w, r, "/cron", http.StatusFound)
}

func getEventsToDelete(c appengine.Context, hours time.Duration) (eventKeys []*datastore.Key, err error){
	q := datastore.NewQuery("TimelineEvent").KeysOnly().Filter("Date <", time.Now().Add(-1 * hours))
	eventKeys, err = q.GetAll(c, nil)
	return eventKeys, err
}