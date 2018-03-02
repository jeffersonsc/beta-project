package repo

import (
	"time"

	"github.com/jeffersonsc/beta-project/conf"
	"github.com/jeffersonsc/beta-project/model"
	"gopkg.in/mgo.v2/bson"
)

func CreateTask(projectID string, boardID string, task *model.Task) error {
	col, err := conf.GetMongoCollection("projects")
	if err != nil {
		return err
	}

	task.ID = bson.NewObjectId()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	query := bson.M{
		"_id":        bson.ObjectIdHex(projectID),
		"boards._id": bson.ObjectIdHex(boardID),
	}

	update := bson.M{
		"$set": bson.M{"updated_at": time.Now()},
		"$push": bson.M{
			"board.$.tasks": task,
		},
	}

	return col.Update(query, update)

}

func UpdateTask(projectID string, taskID string, task *model.Task) error {
	col, err := conf.GetMongoCollection("projects")
	if err != nil {
		return err
	}

	query := bson.M{
		"_id":              bson.ObjectIdHex(projectID),
		"boards.tasks._id": bson.ObjectIdHex(taskID),
	}

	update := bson.M{
		"$set": bson.M{
			"name":        task.Name,
			"description": task.Description,
			"task_type":   task.TaskType,
			"cost":        task.Cost,
		},
	}

	return col.Update(query, update)
}
