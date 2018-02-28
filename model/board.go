package model

// Board struct
type Board struct {
	Base
	Name   string `json:"name" bson:"name"`
	Status string `json:"status" bson:"status"`
	Tasks  []Task `json:"tasks" bson:"tasks"`
}
