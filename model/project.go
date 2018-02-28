package model

// Project struct
type Project struct {
	Base
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Status      string  `json:"status" bson:"status"`
	Boards      []Board `json:"boards" bson:"boards"`
}
