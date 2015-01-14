package index

import (
	// "appengine"
	// "github.com/tlatin/simpletimeline/timeline"
	// "net/http"
	"appengine/aetest"
  "net/http/httptest"
  "net/http"
	"testing"
)

func TestEventPost(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	// req1, err := c.NewRequest("GET", "/", nil)
	// if err != nil {
  //   t.Fatalf("Failed to create req1: %v", err)
	// }

  // r, _ := http.NewRequest("GET", "/", nil)
  // w := httptest.NewRecorder()
  // Get(w, r)
  // if 200 != w.Code {
  //     t.Fail()
  // }

  // req, _ := http.NewRequest("GET", "", nil)
  // w := httptest.NewRecorder()
  // if w.Code != http.StatusOK {
  //     t.Errorf("Home page didn't return %v", http.StatusOK)
  // }

}