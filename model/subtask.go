package model

type SubTask struct {
	Base
	Name        string  `json:"name" bson:"name"`
	Evaluate    string  `json:"evaluate" bson:"evaluate"`
	Cost        float32 `json:"cost" bson:"cost"`
	StatusBonus string  `json:"status_bonus" bson:"status_bonus"`
}
