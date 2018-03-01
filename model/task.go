package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Task struct
type Task struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name            string        `json:"name" bson:"name"`
	Description     string        `json:"description" bson:"description"`
	TaskType        string        `json:"task_type" bson:"task_type"`
	Cost            float32       `json:"cost" bson:"cost"`
	Evaluation      string        `json:"evaluation" bson:"evaluation"`
	StatusEvaluated string        `json:"status_evaluated" bson:"status_evaluated"`
	StatusComplited string        `json:"status_completed" bson:"status_completed"`
	Color           string        `json:"color" bson:"color"`
	SubTasks        []SubTask     `json:"subtasks" bson:"subtasks"`
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" bson:"updated_at"`
}
