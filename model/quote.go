package model

// Quote have information for quotes sent to clients/companies
type Quote struct {
	Model
	Amount int `bson:"amount" json:"amount"`
}
