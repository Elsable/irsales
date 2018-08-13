package model

import "gopkg.in/mgo.v2/bson"

// Quote have information for quotes sent to clients/companies
type Quote struct {
	Model
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Amount int           `bson:"amount" json:"amount"`
}
