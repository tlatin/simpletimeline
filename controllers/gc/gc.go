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

var defaultGcAge = time.Hour * 24 * 30 * 12 // default to 1 year
var defaultLimit = 10000

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.Timeout(appengine.NewContext(r), 60*time.Second)

	age, err := GetHours(r.FormValue("age"))
	if err != nil {
		c.Errorf("error parsing age param. Using default of 1 year")
	}
	q := GetQuery(age, defaultLimit)
	complete, err := RunQuery(c, q, defaultLimit)
	if err != nil {
		// log something
	}
	if complete != true {
		t := taskqueue.NewPOSTTask("/gc", url.Values{"age": {strconv.FormatFloat(time.Duration.Hours(age), 'f', 0, 64)}})
		taskqueue.Add(c, t, "gc")
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

// To determine if the clean up is complete, check to see if the expected number events was deleted.
func RunQuery(c appengine.Context, q *datastore.Query, finishedCount int) (complete bool, err error) {
	deleted := 0
	for i := q.Run(c); ; {
		key, err := i.Next(nil)
		if err != nil {
			// If you've finished the query and deleted LESS than the total number expected, you've deleted all the items to gc.
			if err == datastore.Done && deleted < finishedCount {
				return true, nil
			}
			return false, nil
		}

		datastore.Delete(c, key)
		deleted++
		time.Sleep(time.Second * 5)
	}
	return false, nil
}
