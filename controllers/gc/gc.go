package gc

import (
	"appengine"
	"appengine/datastore"
	"appengine/taskqueue"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var defaultGcAge = time.Hour * 24 * 30 * 9 // default to 9 months
var defaultLimit = 100
func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.Timeout(appengine.NewContext(r), 60*time.Second)
	c.Debugf("Cleaning up old tasks")
	age, err := GetHours(r.FormValue("age"))
	limit := defaultLimit
	if err != nil {
		c.Warningf("error parsing age param. Using default of 1 year")
	}
	c.Debugf("Generating Query")
	q := GetQuery(age, limit)
	c.Debugf("Gathering Events")
	events, err := q.GetAll(c, nil)
	if err != nil {
		c.Errorf("error gathering events: " + err.Error())
		return
	}
	c.Debugf("Deleting Events")
	err = datastore.DeleteMulti(c, events)
	if err != nil {
		c.Errorf("error deleting events: " + err.Error())
		return
	}

	if len(events) == limit {
		c.Infof("Creating a task to delete more items.")
		t := taskqueue.NewPOSTTask("/gc", url.Values{"age": {strconv.FormatFloat(time.Duration.Hours(age), 'f', 0, 64)}})
		_, err = taskqueue.Add(c, t, "gc")
		if err != nil {
			c.Errorf("error creating task: " + err.Error())
		}
	} else {
		c.Infof("Finished deleting all items!")
	}
}

func GetHours(ageStr string) (age time.Duration, err error) {
	age, err = time.ParseDuration(ageStr + "h")
	if "" == ageStr || err != nil {
		age = defaultGcAge
	}
	return age, err
}

func GetQuery(hours time.Duration, limit int) (q *datastore.Query) {
	// limit 10000 because I don't see how to catch deadline exceptions
	q = datastore.NewQuery("TimelineEvent").KeysOnly().Filter("Date <", time.Now().Add(-1*hours)).Limit(limit)
	return q
}