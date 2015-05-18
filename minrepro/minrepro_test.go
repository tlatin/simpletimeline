package Minrepro

import (
	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"testing"
)

func TestDatastoreQuery(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	parentKey := CreateAndRetrieveParent(c, t)
	CreateAndRetrieveChild(c, t, parentKey)

	q := datastore.NewQuery("Child").Ancestor(parentKey)
	children, err := q.GetAll(c, nil)
	if len(children) != 1 {
		t.Errorf("TestDatastoreQuery expected 1 child, found %d", len(children))
	}
}

func CreateAndRetrieveParent(c appengine.Context, t *testing.T) (key *datastore.Key) {
	key, err := NewParent(c, "Alice")
	if err != nil {
		t.Error("Error Creating a new parent datstore object")
		return nil
	}

	// try looking the object up by key to block on the datastore write
	parent := new(Parent)
	if err = datastore.Get(c, key, parent); err != nil {
		t.Error("Error getting parent object")
	}

	if parent.Name != "Alice" {
		t.Error("Wrong name on retrieved parent object")
	}
	return key
}

func CreateAndRetrieveChild(c appengine.Context, t *testing.T, parent *datastore.Key) (key *datastore.Key) {
	key, err := NewChild(c, parent, "Bob")
	if err != nil {
		t.Error("Error Creating a new child datstore object")
		return
	}

	// try looking the object up by key to block on the datastore write
	child := new(Child)
	if err = datastore.Get(c, key, child); err != nil {
		t.Error("Error getting child object")
	}

	if child.Name != "Bob" {
		t.Error("Wrong name on retrieved child object")
	}

	return key
}
