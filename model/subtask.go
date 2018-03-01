package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SubTask struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Evaluate    string        `json:"evaluate" bson:"evaluate"`
	Cost        float32       `json:"cost" bson:"cost"`
	StatusBonus string        `json:"status_bonus" bson:"status_bonus"`
	Finished    bool          `json:"finished" bson:"finished"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}
