package http

import (
	"github.com/gin-gonic/gin"
	"todo-list/internal/task/repository"
	"todo-list/internal/task/service"
)

func Routes(r *gin.RouterGroup) {
	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := NewTaskHandler(taskService)
	taskRoute := r.Group("/tasks")
	{
		taskRoute.GET("/", taskHandler.ListTasks)
		taskRoute.POST("/", taskHandler.CreateTask)
		taskRoute.PUT("/:id", taskHandler.UpdateTask)
		taskRoute.DELETE("/:id", taskHandler.DeleteTask)
	}
}
