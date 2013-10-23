package Timeline

import (
	"time"
	"sort"
	"testing"
)

// t.Error for errors
// t.Fail for failures
func TestTimelineSorting(t *testing.T) {
	timestamp := time.Now()
	older := Event{
		AuthorId: "older",
		Content:  "This event is in the past",
		Date:     timestamp.Add(-1000*time.Second),
	}
	newer := Event{
		AuthorId: "newer",
		Content:  "This event is now",
		Date:     timestamp.Add(1000*time.Second),
	}
	events := []Event{newer, older}
	sort.Sort(ByDate(events))
	if !events[0].Date.After(events[1].Date) {
		t.Error("Events were sorted in the wrong order")
	}
}
