package model

// Person holds information for people
type Person struct {
	Model
	Name  string `bson:"name" json:"name"`
	Phone string `bson:"phone" json:"phone"`
}
