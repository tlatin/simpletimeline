package cron

import (
	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"github.com/tlatin/simpletimeline/timeline"
	"testing"
	"time"
)

// This is failing. Created min repro case and emailed the goappengine group
func TestGetEventsToDelete(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if timeWindow != time.Hour*24*30*3 {
		t.Fatalf("timeWindow in cron is not what was expected.")
	}

	// timeWindow := timeWindow //time.Hour * 24 * 30 * 3
	eventDates := []time.Duration{timeWindow - 10*time.Hour, 2 * timeWindow, timeWindow + 20*time.Hour}
	if len(eventDates) != 3 {
		t.Fatalf("failed to create array of len 3 for TestGetEventsToDelete")
	}

	for _, hoursAgo := range eventDates {
		if _, err := addOldEvent(c, hoursAgo); err != nil {
			t.Fatal(err.Error())
			return
		}
	}

	// Should return all events
	eventKeys, err := getEventsToDelete(c, 0*time.Hour, 1000)
	if err != nil {
		t.Fatalf("error getting events to delete", err)
	}
	if len(eventKeys) != 3 {
		t.Fatalf("expected all 3 events, found %d", len(eventKeys))
	}

	// Should return 2 of the 3 events created
	eventKeys, err = getEventsToDelete(c, timeWindow, 1000)
	if err != nil {
		t.Fatalf("error getting events to delete", err)
	}
	if len(eventKeys) != 2 {
		t.Fatalf("expected 2, found %d", len(eventKeys))
	}
}

func addOldEvent(c appengine.Context, hoursAgo time.Duration) (key *datastore.Key, err error) {
	te := Timeline.Event{
		AuthorId: "cron-testauthor",
		Content:  "cront-test content",
		Date:     time.Now().Add(-1 * hoursAgo),
	}

	if key, err = datastore.Put(c, datastore.NewIncompleteKey(c, "TimelineEvent", nil), &te); err != nil {
		return key, err
	}

	// force consistency by looking up the event by key
	err = datastore.Get(c, key, new(Timeline.Event))
	return key, err
}
