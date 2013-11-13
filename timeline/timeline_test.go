package Timeline

import (
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
	nilkey := new(datastore.Key)
	key, err := NewEvent(c, nilkey, authorId, content)
	if err != nil {
		t.Error("Error Creating a new application: " + err.Error())
	}

	event := new(Event)
	if err := datastore.Get(c, key, event); err != nil {
		t.Error("Error getting Event: " + err.Error())
	}
	if event.AuthorId != authorId {
		t.Error("Returned event has the wrong authorId`")
	}
	if event.Content != content {
		t.Error("Returned event has the wrong content")
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
	key, err := NewApplication(c, name, url)
	if err != nil {
		t.Error("Error Creating a new application: " + err.Error())
	}

	app := new(Application)
	if err := datastore.Get(c, key, app); err != nil {
		t.Error("Error getting application: " + err.Error())
	}
	if app.Name != name {
		t.Error("Returned application has the wrong name")
	}
	if app.Url != url {
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
	key, err := NewApplication(c, name, url)
	if err != nil {
		t.Error("Error Creating a new application: " + err.Error())
	}

	keystr := key.Encode()
	if "" == keystr {
		t.Error("The key string is empty.")
		return
	}
	app, err := GetApplicationByEncodedKey(c, keystr)
	if err != nil {
		t.Error("Error Retrieving an application: " + err.Error())
	}	
	if app.Name != name {
		t.Error("Returned application has the wrong name")
	}
	if app.Url != url {
		t.Error("Returned application has the url name")
	}
}

func TestGetApplicationKeyByString(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	name := "Example Name"
	url := "http://example.com"
	key, err := NewApplication(c, name, url)
	if err != nil {
		t.Error("Error Creating a new application: " + err.Error())
	}

	keystr := key.Encode()
	if "" == keystr {
		t.Error("The key string is empty.")
		return
	}
	appkey, err := GetApplicationKeyByString(c, keystr)
	if err != nil {
		t.Error("Error retrieving the key for the application: " + err.Error())
	}	
	if !key.Equal(appkey) {
		t.Error("GetApplicationKeyByString doesn't match the Put() key")
	}
}