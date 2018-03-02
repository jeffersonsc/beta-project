package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffersonsc/beta-project/lib/context"
	"github.com/jeffersonsc/beta-project/model"
	"github.com/jeffersonsc/beta-project/repo"
)

// CreateProjectHandler save and create structure on project
func CreateProjectHandler(ctx *context.Context) {

	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on read body"})
		return
	}
	defer ctx.Req.Body().ReadCloser()

	project := model.Project{}
	if err := json.Unmarshal(body, &project); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error parse json"})
		return
	}

	if err := repo.CreateProject(&project); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on save project. ERROR" + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, project)
}

// AllProjectsHandler select all projetcs
func AllProjectsHandler(ctx *context.Context) {
	projects, err := repo.FindAllProjects()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on get projects"})
		return
	}
	if projects == nil {
		ctx.JSON(http.StatusOK, map[string][]string{"projects": []string{}})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"projects": projects})
}

// FindProjectHandler find by id
func FindProjectHandler(ctx *context.Context) {
	log.Println("ID ", ctx.Params("id"))
	project, err := repo.FindProject(ctx.Params("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on get project. ERROR: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"project": project})
}

// UpdateProjectHandler change project
func UpdateProjectHandler(ctx *context.Context) {
	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on read body"})
		return
	}
	defer ctx.Req.Body().ReadCloser()

	project := model.Project{}
	if err := json.Unmarshal(body, &project); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error parse json"})
		return
	}

	id := ctx.Params("id")

	if err := repo.UpdateProject(id, &project); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on update project. ERROR: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"": ""})
}
