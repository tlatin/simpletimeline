package Timeline

import (
	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"sort"
	"testing"
	"time"
)

// t.Error for errors
// t.Fail for failures
func TestTimelineSorting(t *testing.T) {
	timestamp := time.Now()
	older := Event{
		AuthorId: "older",
		Content:  "This event is in the past",
		Date:     timestamp.Add(-1000 * time.Second),
	}
	newer := Event{
		AuthorId: "newer",
		Content:  "This event is now",
		Date:     timestamp.Add(1000 * time.Second),
	}
	events := []Event{newer, older}
	sort.Sort(ByDate(events))
	if !events[0].Date.After(events[1].Date) {
		t.Error("Events were sorted in the wrong order")
	}
}

func TestNewEvent(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	authorId := "This is an Author ID"
	content := "this is the content"
	key := CreateTestEvent(t, c, nil, authorId, content)

	event := new(Event)
	if err := datastore.Get(c, key, event); err != nil {
		t.Error("Error getting Event: " + err.Error())
	} else if event.AuthorId != authorId {
		t.Error("Returned event has the wrong authorId`")
	} else if event.Content != content {
		t.Error("Returned event has the wrong content")
	}
}

func TestQueryEvent(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	key := CreateExampleEvent(t, c)
	event := new(Event)
	if err := datastore.Get(c, key, event); err != nil {
		t.Error("Error getting Event: " + err.Error())
	}

	app, _ := CreateExampleEventWithApplication(t, c)
	q := datastore.NewQuery("TimelineEvent").Ancestor(app).KeysOnly()
	eventKeys, err := q.GetAll(c, nil)
	if err != nil {
		t.Error("Error in the query: " + err.Error())
	}

	if len(eventKeys) != 1 {
		t.Errorf("TestQueryEvent expected 1 event, found %d", len(eventKeys))
	}

}

func TestNewApplication(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	name := "Example Name"
	url := "http://example.com"
	key := CreateTestApplication(t, c, name, url)
	app := new(Application)
	if err := datastore.Get(c, key, app); err != nil {
		t.Error("Error getting application: " + err.Error())
	} else if app.Name != name {
		t.Error("Returned application has the wrong name")
	} else if app.Url != url {
		t.Error("Returned application has the url name")
	}
}

func TestGetApplicationByEncodedKey(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	name := "Example Name"
	url := "http://example.com"
	key := CreateTestApplication(t, c, name, url)
	keystr := key.Encode()
	if "" == keystr {
		t.Error("The key string is empty.")
		return
	}
	app, err := GetApplicationByEncodedKey(c, keystr)
	if err != nil {
		t.Error("Error Retrieving an application: " + err.Error())
	} else if app.Name != name {
		t.Error("Returned application has the wrong name")
	} else if app.Url != url {
		t.Error("Returned application has the url name")
	}
}

func TestGetApplicationKeyByString(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	key := CreateExampleApplication(t, c)
	keystr := key.Encode()
	if "" == keystr {
		t.Error("The key string is empty.")
		return
	}
	appkey, err := GetApplicationKeyByString(c, keystr)
	if err != nil {
		t.Error("Error retrieving the key for the application: " + err.Error())
	} else if !key.Equal(appkey) {
		t.Error("GetApplicationKeyByString doesn't match the Put() key")
	}
}

//
// Helper Functions
//

func CreateExampleApplication(t *testing.T, c appengine.Context) (key *datastore.Key) {
	name := "Example Name"
	url := "http://example.com"
	return CreateTestApplication(t, c, name, url)
}

func CreateTestApplication(t *testing.T, c appengine.Context, name string, url string) (key *datastore.Key) {
	key, err := NewApplication(c, name, url)
	if err != nil {
		t.Error("Error Creating a new application: " + err.Error())
		return
	}

	// Doing a get to force consistency. Lame I know.
	// http://stackoverflow.com/questions/24159413/gae-go-tests-do-datastore-queries-work-in-test-environment
	app := new(Application)
	if err := datastore.Get(c, key, app); err != nil {
		t.Error("Error getting application: " + err.Error())
	}
	return key
}

func CreateExampleEventWithApplication(t *testing.T, c appengine.Context) (app *datastore.Key, event *datastore.Key) {
	app = CreateExampleApplication(t, c)
	authorId := "This is an Author ID"
	content := "this is the content"
	event = CreateTestEvent(t, c, app, authorId, content)
	return app, event

}

func CreateExampleEvent(t *testing.T, c appengine.Context) (key *datastore.Key) {
	authorId := "This is an Author ID"
	content := "this is the content"
	return CreateTestEvent(t, c, nil, authorId, content)
}

func CreateTestEvent(t *testing.T, c appengine.Context, application *datastore.Key, authorId string, content string) (key *datastore.Key) {
	key, err := NewEvent(c, application, authorId, content)
	if err != nil {
		t.Error("Error Creating a new Event Object: " + err.Error())
		return
	}

	// Doing a get to force consistency. Lame I know.
	// http://stackoverflow.com/questions/24159413/gae-go-tests-do-datastore-queries-work-in-test-environment
	event := new(Event)
	if err := datastore.Get(c, key, event); err != nil {
		t.Error("Error getting Event: " + err.Error())
	}
	return key
}
