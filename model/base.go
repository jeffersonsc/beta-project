package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Base struct
type Base struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
