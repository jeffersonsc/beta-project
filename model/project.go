package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Project struct
type Project struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Status      string        `json:"status" bson:"status"`
	Boards      []Board       `json:"boards" bson:"boards"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}
