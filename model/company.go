package model

// Company holds details for clients
type Company struct {
	Model
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`
}
