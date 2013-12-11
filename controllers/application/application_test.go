package application

import (
	"appengine/aetest"
	"testing"
)

func TestApplicationPost(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
}

func TestparseApplicationKeyFromURL(t *testing.T) {
	test_url := "application/ahJkZXZ-c2ltcGxldGltZWxpbmVyGAsSC0FwcGxpY2F0aW9uGICAgICAgIAKDA"
	expected_key := "ahJkZXZ-c2ltcGxldGltZWxpbmVyGAsSC0FwcGxpY2F0aW9uGICAgICAgIAKDA"
	appKey, err := parseApplicationKeyFromURL(test_url)
	if err != nil {
		t.Error("Parse Application Key threw an unexpected error")
		return
	}
	if appKey != expected_key {
		t.Error("The parsed key is not the expected key")
		return
	}
}