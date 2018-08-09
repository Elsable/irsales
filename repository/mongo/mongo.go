/**
*	This file creates MongoDB interfaces and structs so the can be more testable
*	See: http://thylong.com/golang/2016/mocking-mongo-in-golang/
**/

package mongo

import (
	mgo "gopkg.in/mgo.v2"
)

// Session is an interface to access to the Session struct.
type Session interface {
	DB(name string) DataLayer
	Close()
}

// MgoSession is currently a Mongo session.
type MgoSession struct {
	Session *mgo.Session
}

// DataLayer is an interface to access to the database struct.
type DataLayer interface {
	C(name string) Collection
}

// MgoCollection wraps a mgo.Collection to embed methods in models.
type MgoCollection struct {
	Collection *mgo.Collection
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Find(query interface{}) *mgo.Query
	Count() (n int, err error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
}

// MgoDatabase wraps a mgo.Database to embed methods in models.
type MgoDatabase struct {
	Database *mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d MgoDatabase) C(name string) Collection {
	// return &MgoCollection{ Collection: d.Database.C(name)}
	return MgoCollection{Collection: d.Database.C(name)}.Collection
}
