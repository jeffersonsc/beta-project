package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jeffersonsc/beta-project/lib/context"
	"github.com/jeffersonsc/beta-project/model"
	"github.com/jeffersonsc/beta-project/repo"
)

func CreateTaskHandler(ctx *context.Context) {
	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on read body"})
		return
	}
	defer ctx.Req.Body().ReadCloser()

	task := model.Task{}
	if err := json.Unmarshal(body, &task); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error parse json"})
		return
	}

	projectID := ctx.Params("project_id")
	boardID := ctx.Query("board_id")

	if err := repo.CreateTask(projectID, boardID, &task); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on save task. ERROR: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"task": task})
}
func UpdateTaskHandler(ctx *context.Context) {
	body, err := ctx.Req.Body().Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on read body"})
		return
	}
	defer ctx.Req.Body().ReadCloser()

	task := model.Task{}
	if err := json.Unmarshal(body, &task); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error parse json"})
		return
	}

	projectID := ctx.Params("project_id")
	taskID := ctx.Params("id")

	if err := repo.UpdateTask(projectID, taskID, &task); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error on save task. ERROR: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"": ""})
}
