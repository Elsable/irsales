package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Company holds details for clients
type Company struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Location string        `bson:"location" json:"location"`
}
