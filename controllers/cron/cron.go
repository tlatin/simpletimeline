package cron

import (
	"appengine"
	"appengine/datastore"
	"github.com/tlatin/simpletimeline/utils"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

var cronTemplate = template.Must(template.ParseFiles(utils.GetTemplatePath() + "cron.html"))
var timeWindow = time.Hour * 24 * 30 * 3 // Look at all events more than 90 days old

type CronTemplateValues struct {
	Keys       []*datastore.Key
	TimeWindow time.Duration
	Limit      int
}

func Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.Timeout(appengine.NewContext(r), 60*time.Second)

	limitStr := r.FormValue("limit")
	if "" == limitStr {
		limitStr = "-1"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		if utils.CheckHandlerError(c, err, w, "failed to parse limit param.") {
			return
		}
	}

	oldTimelineEvents, err := getEventsToDelete(c, timeWindow, limit)
	if utils.CheckHandlerError(c, err, w, "failed to load all timeline events.") {
		return
	}

	templateValues := CronTemplateValues{Keys: oldTimelineEvents, TimeWindow: timeWindow, Limit: limit}
	if utils.CheckHandlerError(c, cronTemplate.Execute(w, templateValues), w, "cron template failed to load.") {
		return
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.Timeout(appengine.NewContext(r), 60*time.Second)

	limitStr := r.FormValue("limit")
	if "" == limitStr {
		limitStr = "-1"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		if utils.CheckHandlerError(c, err, w, "failed to parse limit param.") {
			return
		}
	}

	eventsArray, err := getEventsToDelete(c, timeWindow, limit)
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

func getEventsToDelete(c appengine.Context, hours time.Duration, limit int) (eventKeys []*datastore.Key, err error) {
	q := datastore.NewQuery("TimelineEvent").KeysOnly().Filter("Date <", time.Now().Add(-1*hours)).Limit(limit)

	for i := q.Run(c); ; {
		key, err := i.Next(nil)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		eventKeys = append(eventKeys, key)
	}
	eventKeys, err = q.GetAll(c, nil)
	return eventKeys, err
}
