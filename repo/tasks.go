package repo

import (
	"fmt"
	"log"
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
			"boards.$.tasks": task,
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
			"boards.$.tasks.$.name":             task.Name,
			"boards.$.tasks.$.description":      task.Description,
			"boards.$.tasks.$.task_type":        task.TaskType,
			"boards.$.tasks.$.cost":             task.Cost,
			"boards.$.tasks.$.evaluation":       task.Evaluation,
			"boards.$.tasks.$.status_evaluated": task.StatusEvaluated,
			"boards.$.tasks.$.status_completed": task.StatusComplited,
			"boards.$.tasks.$.color":            task.Color,
			"boards.$.tasks.$.updated_at":       time.Now(),
		},
	}

	return col.Update(query, update)
}

func MoveTaskBoard(projectID, fromBoard, toBoard, taskID string) (err error) {
	col, err := conf.GetMongoCollection("projects")
	if err != nil {
		return err
	}
	queryFind := bson.M{
		"_id":                bson.ObjectIdHex(projectID),
		"boards._id":         bson.ObjectIdHex(projectID),
		"boards.tasks.$._id": bson.ObjectIdHex(taskID),
	}

	queryRemove := bson.M{
		"_id":        bson.ObjectIdHex(projectID),
		"boards._id": bson.ObjectIdHex(fromBoard),
	}

	queryNew := bson.M{
		"_id":        bson.ObjectIdHex(projectID),
		"boards._id": bson.ObjectIdHex(toBoard),
	}

	// Find task on board
	task := model.Task{}
	err = col.Find(queryFind).One(&task)
	if err != nil {
		return fmt.Errorf("Error on get task. ERROR: %v", err)
	}

	log.Printf("%+v", task)
	return

	// Put task in new board
	update := bson.M{
		"$set": bson.M{"updated_at": time.Now()},
		"$push": bson.M{
			"boards.$.tasks": task,
		},
	}

	// Remove older board
	err = col.Update(queryRemove, bson.M{"$pull": bson.M{"boards.$.tasks": bson.M{"_id": task.ID}}})
	if err != nil {
		return fmt.Errorf("Error on remove task of old board. ERROR: %v", err)
	}

	// Add in new board
	err = col.Update(queryNew, update)
	if err != nil {
		return fmt.Errorf("Error on set task into new board. ERROR: %v", err)
	}

	return
}
