package handler

import (
	"encoding/json"
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

// FindProjectsHandler select all projetcs
func FindProjectsHandler(ctx *context.Context) {
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
