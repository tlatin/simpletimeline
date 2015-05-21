package Minrepro

import (
	"appengine"
	"appengine/datastore"
)

type Parent struct {
	Name string
}

type Child struct {
	Name string
}

func NewParent(c appengine.Context, name string) (key *datastore.Key, err error) {
	parent := Parent{
		Name: name,
	}
	return datastore.Put(c, datastore.NewIncompleteKey(c, "Parent", nil), &parent)
}

func NewChild(c appengine.Context, parent *datastore.Key, name string) (key *datastore.Key, err error) {
	child := Child{
		Name: name,
	}
	return datastore.Put(c, datastore.NewIncompleteKey(c, "Child", parent), &child)
}
