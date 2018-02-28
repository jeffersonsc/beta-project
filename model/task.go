package model

// Task struct
type Task struct {
	Base
	Name            string    `json:"name" bson:"name"`
	Description     string    `json:"description" bson:"description"`
	TaskType        string    `json:"task_type" bson:"task_type"`
	Cost            float32   `json:"cost" bson:"cost"`
	Evaluation      string    `json:"evaluation" bson:"evaluation"`
	StatusEvaluated string    `json:"status_evaluated" bson:"status_evaluated"`
	StatusComplited string    `json:"status_completed" bson:"status_completed"`
	Color           string    `json:"color" bson:"color"`
	SubTasks        []SubTask `json:"subtasks" bson:"subtasks"`
}
