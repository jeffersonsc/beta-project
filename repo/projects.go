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

	for _, v := range []string{"Backlog", "Doing", "Done"} {
		board := model.Board{}
		board.ID = bson.NewObjectId()
		board.Name = v
		board.CreatedAt = time.Now()
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
