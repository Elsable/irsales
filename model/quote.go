package model

// Quote have information for quotes sent to clients/companies
type Quote struct {
	ID     int `bson:"_id" json:"id"`
	Amount int `bson:"amount" json:"amount"`
}
