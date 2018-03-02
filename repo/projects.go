package repo

import (
	"time"

	"github.com/jeffersonsc/beta-project/conf"
	"github.com/jeffersonsc/beta-project/model"
	"gopkg.in/mgo.v2/bson"
)

// CreateProject save my project in database
func CreateProject(project *model.Project) (err error) {
	col, err := conf.GetMongoCollection("projects")
	if err != nil {
		return
	}

	project.ID = bson.NewObjectId()
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	subtask := model.SubTask{
		ID:        bson.NewObjectId(),
		Name:      "Teste de subtask",
		CreatedAt: time.Now(),
	}

	task := model.Task{
		ID:          bson.NewObjectId(),
		Name:        "Task Default",
		Description: "Description task",
		Color:       "#C0C0C0",
		Cost:        100,
		CreatedAt:   time.Now(),
	}

	for _, v := range []string{"Backlog", "Doing", "Done"} {
		board := model.Board{}
		board.ID = bson.NewObjectId()
		board.Name = v
		board.CreatedAt = time.Now()

		task.SubTasks = append(task.SubTasks, subtask)

		board.Tasks = append(board.Tasks, task)

		project.Boards = append(project.Boards, board)
	}

	err = col.Insert(project)
	return
}

// FindAllProjects return all project onjects
func FindAllProjects() (projects []model.Project, err error) {
	col, err := conf.GetMongoCollection("projects")
	if err != nil {
		return
	}

	err = col.Find(bson.M{}).All(&projects)
	return
}

// FindProject return project by id
func FindProject(id string) (project model.Project, err error) {
	col, err := conf.GetMongoCollection("projects")
	if err != nil {
		return
	}

	err = col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&project)
	return
}

// UpdateProject change project
func UpdateProject(id string, project *model.Project) error {
	col, err := conf.GetMongoCollection("projects")
	if err != nil {
		return err
	}

	return col.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"name": project.Name, "description": project.Description, "status": project.Status, "updated_at": time.Now()}})
}
