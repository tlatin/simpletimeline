package Timeline

import (
	"time"
)

type Application struct {
	Name string
	Url  string
}

type Event struct {
	Application   Application
	AuthorId      string
	Content       string
	Date          time.Time
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
