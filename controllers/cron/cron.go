package cron

import (
	"appengine"
	"appengine/datastore"
	"github.com/tlatin/simpletimeline/utils"
	// "github.com/tlatin/simpletimeline/timeline"
	"html/template"
	"net/http"
	"time"
)

var cronTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath() + "cron.html"))
var timeWindow = time.Hour * 24 * 30 * 3 

type CronTemplateValues struct {
	Keys []*datastore.Key
	TimeWindow time.Duration
}

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// List all of the events older than month.
	// q := datastore.NewQuery("TimelineEvent").Filter("Date <", time.Now().Add(-1 * time.Hour * 24 * 30) )
	// var oldTimelineEvents []Timeline.Event
	// var err error
	// if _, err := q.GetAll(c, &oldTimelineEvents); err != nil {
	oldTimelineEvents, err := getEventsToDelete(c, timeWindow);
	if err != nil {
		c.Errorf("failed to load all timeline events.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateValues := CronTemplateValues{Keys: oldTimelineEvents, TimeWindow:timeWindow}
	if err := cronTemplate.Execute(w, templateValues); err != nil {
		c.Errorf("cron template failed to load.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	eventsArray, err := getEventsToDelete(c, timeWindow);
	if err != nil {
		c.Errorf("failed to load all timeline events.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, eventKey := range eventsArray {
		if err := datastore.Delete(c, eventKey); err != nil {
			c.Errorf("event datastore failed to delete.")
			http.Error(w, err.Error(), http.StatusInternalServerError)
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