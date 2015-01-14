package event

import (
	// "appengine"
	// "github.com/tlatin/simpletimeline/timeline"
	// "net/http"
	"appengine/aetest"
	// "httptest"
	"testing"
)

func TestEventPost(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

  // r, _ := http.NewRequest("GET", "/", nil)
  // w := httptest.NewRecorder()
  // myNewHandler(c, w, r)
  // if 200 != w.Code {
  //     t.Fail()
  // }


    // req, _ := http.NewRequest("GET", "", nil)
    // w := httptest.NewRecorder()
    // if w.Code != http.StatusOK {
    //     t.Errorf("Home page didn't return %v", http.StatusOK)
    // }

}