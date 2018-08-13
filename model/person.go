package model

import "gopkg.in/mgo.v2/bson"

// Person holds information for people
type Person struct {
	Model
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Phone string        `bson:"phone" json:"phone"`
}
