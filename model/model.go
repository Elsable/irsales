package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Model holds generic information for models
type Model struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
}
