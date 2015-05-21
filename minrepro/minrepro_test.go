package Minrepro

import (
	"appengine"
	"appengine/aetest"
	"appengine/datastore"
	"testing"
	"time"
)

// looking up the newly created child by query fails.
func TestDatastoreQuery(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	childKey, parentKey, err := CreateChildAndParentExample(c, t)
	// try looking the object up by key to block on the datastore write
	child := new(Child)
	if err = datastore.Get(c, childKey, child); err != nil {
		t.Error("Error getting child object")
	}

	q := datastore.NewQuery("Child").Ancestor(parentKey)
	var children []Child
	_, err = q.GetAll(c, &children)
	if err != nil {
		t.Error("Error in the query: " + err.Error())
	}

	if len(children) != 1 {
		t.Errorf("TestDatastoreQuery expected 1 child, found %d", len(children))
	}
}

func TestDatastoreQueryWithSleep(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	childKey, parentKey, err := CreateChildAndParentExample(c, t)
	// try looking the object up by key to block on the datastore write
	child := new(Child)
	if err = datastore.Get(c, childKey, child); err != nil {
		t.Error("Error getting child object")
	}

	time.Sleep(time.Second * 10)
	q := datastore.NewQuery("Child").Ancestor(parentKey)
	var children []Child
	_, err = q.GetAll(c, &children)
	if err != nil {
		t.Error("Error in the query: " + err.Error())
	}

	if len(children) != 1 {
		t.Errorf("TestDatastoreQuery expected 1 child, found %d", len(children))
	}
}

// Looking up the newly created child by key works.
func TestGetChildByKey(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	childKey, _, err := CreateChildAndParentExample(c, t)
	// try looking the object up by key to block on the datastore write
	childObj := new(Child)
	if err = datastore.Get(c, childKey, childObj); err != nil {
		t.Error("Error getting child object")
	}

}

func CreateChildAndParentExample(c appengine.Context, t *testing.T) (childKey *datastore.Key, parentKey *datastore.Key, err error) {
	parentKey, err = NewParent(c, "Alice")
	if err != nil {
		t.Error("Error Creating a new parent datstore object")
		return childKey, parentKey, err
	}

	// try looking the object up by key to block on the datastore write
	parentObj := new(Parent)
	if err = datastore.Get(c, parentKey, parentObj); err != nil {
		t.Error("Error getting parent object")
		return childKey, parentKey, err
	}

	// create the child with a parent ancestor
	childKey, err = NewChild(c, parentKey, "Bob")
	if err != nil {
		t.Error("Error Creating a new child datstore object")
		return childKey, parentKey, err
	}
	return childKey, parentKey, err
}
