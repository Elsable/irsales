package model

// Model holds generic information for models
type Model struct {
	CreatedAt string `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt string `bson:"updatedAt,omitempty" json:"updatedAt"`
}
