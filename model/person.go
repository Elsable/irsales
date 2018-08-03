package model

// Person holds information for people
type Person struct {
	ID    int    `bson:"_id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Phone string `bson:"phone" json:"phone"`
}
