package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/internal/task/dto"
	"todo-list/internal/task/service"
	"todo-list/pkg/response"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) ListTasks(ctx *gin.Context) {
	tasks, err := h.service.ListTasks(ctx)
	fmt.Println(tasks)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something went wrong")
		return
	}

	response.JSON(ctx, http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var req dto.CreateTaskReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid parameters")
		return
	}

	id, err := h.service.Create(ctx, &req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something went wrong")
		return
	}

	response.JSON(ctx, http.StatusOK, id)
}

func (h *TaskHandler) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateTaskReq
	err := ctx.ShouldBindJSON(&req)
	if id == "" || err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid parameters")
		return
	}

	task, err := h.service.Update(ctx, id, &req)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			response.Error(ctx, http.StatusBadRequest, "task not found")
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "something went wrong")
		return
	}

	response.JSON(ctx, http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.Error(ctx, http.StatusBadRequest, "invalid parameters")
		return
	}

	err := h.service.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			response.Error(ctx, http.StatusBadRequest, "task not found")
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "something went wrong")
		return
	}

	response.JSON(ctx, http.StatusOK, nil)
}
