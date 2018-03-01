package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Board struct
type Board struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Status    string        `json:"status" bson:"status"`
	Tasks     []Task        `json:"tasks" bson:"tasks"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
