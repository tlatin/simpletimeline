package gc

import (
	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"github.com/tlatin/simpletimeline/timeline"
	"strconv"
	"testing"
	"time"
)

func TestGetHours(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if age, _ := GetHours(""); age != defaultGcAge {
		t.Fatalf("default result for getHours not returning default.")
	}
	if age, _ := GetHours("1"); age != time.Hour {
		t.Fatalf("not getting 1 back when passing the string '1 to getHours. getting " + strconv.FormatFloat(time.Duration.Hours(age), 'f', 0, 64))
	}
	if age, err := GetHours("4320"); age != time.Hour*4320 || err != nil {
		t.Fatalf("not getting 4320 back when passing the string '4320' to getHours.")
	}
	if age, err := GetHours("asdf"); age != defaultGcAge || err == nil {
		t.Fatalf("default result for getHours not returning default.")
	}
}

// I thought about breaking up the logic to a func to get the events, and another to delete them.
// This would have made testing easier, but used more memory during run time. Instead I parameterized
// the limit function, so I can test the task creation with a small number of events. The task
// creation is tested in the TestGcPOST func.
func TestSingleIncompleteRunQuery(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	eventDates := []time.Duration{defaultGcAge - 10*time.Hour, defaultGcAge, defaultGcAge}

	for _, hoursAgo := range eventDates {
		if _, err := addDateTestEvent(c, hoursAgo); err != nil {
			t.Fatal(err.Error())
			return
		}
	}

	age := 0 * time.Hour
	q := GetQuery(age, 3)
	count, err := q.Count(c)
	if err != nil {
		t.Fatalf("Error getting count")
	}
	if count != 3 {
		t.Fatalf("Count for RunQuery test isn't 3. Instead it is " + strconv.Itoa(count))
	}

	complete, err := RunQuery(c, q, 2)
	if err != nil {
		t.Fatalf("Error testing RunQuery")
	}
	if complete {
		t.Fatalf("RunQuery returned complete = true incorrectly")
	}

}

// creation is tested in the TestGcPOST func.
func TestSingleCompleteRunQuery(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	eventDates := []time.Duration{defaultGcAge - 10*time.Hour, defaultGcAge, defaultGcAge}

	for _, hoursAgo := range eventDates {
		if _, err := addDateTestEvent(c, hoursAgo); err != nil {
			t.Fatal(err.Error())
			return
		}
	}

	age := 0 * time.Hour
	q := GetQuery(age, 3)
	count, err := q.Count(c)
	if err != nil {
		t.Fatalf("Error getting count")
	}
	if count != 3 {
		t.Fatalf("Count for RunQuery test isn't 3. Instead it is " + strconv.Itoa(count))
	}

	complete, err := RunQuery(c, q, 3)
	if err != nil {
		t.Fatalf("Error testing RunQuery")
	}
	if complete != false {
		t.Fatalf("RunQuery returned complete = false when everything should be deleted")
	}

	q = GetQuery(age, 10)
	count, err = q.Count(c)
	if err != nil {
		t.Fatalf("Error getting count")
	}
	if count != 0 {
		t.Fatalf("Count for RunQuery test isn't 0. Instead it is " + strconv.Itoa(count))
	}

}

// creation is tested in the TestGcPOST func.
func TestMultiCompleteRunQuery(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	eventDates := []time.Duration{defaultGcAge - 10*time.Hour, defaultGcAge, defaultGcAge}

	for _, hoursAgo := range eventDates {
		if _, err := addDateTestEvent(c, hoursAgo); err != nil {
			t.Fatal(err.Error())
			return
		}
	}

	age := 0 * time.Hour
	q := GetQuery(age, 10)
	count, err := q.Count(c)
	if err != nil {
		t.Fatalf("Error getting count")
	}
	if count != 3 {
		t.Fatalf("Count for RunQuery test isn't 3. Instead it is " + strconv.Itoa(count))
	}

	complete, err := RunQuery(c, q, 2)
	if err != nil {
		t.Fatalf("Error testing RunQuery")
	}
	if complete {
		t.Fatalf("RunQuery returned complete = true incorrectly")
	}

	q = GetQuery(age, 2)
	count, err = q.Count(c)
	if err != nil {
		t.Fatalf("Error getting count")
	}
	if count != 1 {
		t.Fatalf("Count for RunQuery test isn't 1. Instead it is " + strconv.Itoa(count))
	}

	complete, err = RunQuery(c, q, 2)
	if err != nil {
		t.Fatalf("Error testing RunQuery")
	}
	if complete == false {
		t.Fatalf("RunQuery returned complete = false when everything should be deleted")
	}

}

func TestGcPOST(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// t.Errorf("Implement the test for the Post method")

}

func addDateTestEvent(c appengine.Context, hoursAgo time.Duration) (key *datastore.Key, err error) {
	te := Timeline.Event{
		AuthorId: "gc-testauthor",
		Content:  "gc-test content",
		Date:     time.Now().Add(-1 * hoursAgo),
	}

	if key, err = datastore.Put(c, datastore.NewIncompleteKey(c, "TimelineEvent", nil), &te); err != nil {
		return key, err
	}

	// force consistency by looking up the event by key
	err = datastore.Get(c, key, new(Timeline.Event))
	return key, err
}
