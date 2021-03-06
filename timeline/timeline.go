package Timeline

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type Application struct {
	Name string
	Url  string
	// This needs to be filled out after a query is done.
	// Don't store it in the datastore.
	// Also, there should be a better way to do this.
	WebKey string `datastore:"-"`
}

func GetApplicationKeyByString(c appengine.Context, applicationKeyStr string) (key *datastore.Key, err error) {
	applicationKey, err := datastore.DecodeKey(applicationKeyStr)
	if err != nil {
		return key, err
	}
	return applicationKey, err
}

// When called from index.go it doesn't appear to be filling the slice.
func GetAllApplications(c appengine.Context, dst []Application, limit int) (err error) {
	q := datastore.NewQuery("Application").Limit(limit)
	keys, err := q.GetAll(c, &dst)
	if err != nil {
		return err
	}
	for i := range dst {
		dst[i].WebKey = keys[i].Encode()
	}
	return err
}

func GetApplicationByEncodedKey(c appengine.Context, applicationKey string) (app *Application, err error) {
	key, err := datastore.DecodeKey(applicationKey)
	if err != nil {
		return app, err
	}
	application := new(Application)
	if err := datastore.Get(c, key, application); err != nil {
		return app, err
	}
	application.WebKey = key.Encode()
	return application, err
}

func NewApplication(c appengine.Context, name string, url string) (key *datastore.Key, err error) {
	app := Application{
		Name: name,
		Url:  url,
	}
	return datastore.Put(c, datastore.NewIncompleteKey(c, "Application", nil), &app)
}

type Event struct {
	// An Application is used as the ancestor for each event.
	AuthorId string
	Content  string
	Date     time.Time
}

func NewEvent(c appengine.Context, application *datastore.Key, authorId string, content string) (key *datastore.Key, err error) {
	te := Event{
		AuthorId: authorId,
		Content:  content,
		Date:     time.Now(),
	}
	return datastore.Put(c, datastore.NewIncompleteKey(c, "TimelineEvent", application), &te)
}

// Methods required by sort.Interface.
type ByDate []Event

func (s ByDate) Len() int {
	return len(s)
}
func (s ByDate) Less(i, j int) bool {
	return s[i].Date.After(s[j].Date)
}
func (s ByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
