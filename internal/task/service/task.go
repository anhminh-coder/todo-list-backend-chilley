package service

import (
	"context"
	"errors"
	"todo-list/internal/task/dto"
	"todo-list/internal/task/model"
	"todo-list/internal/task/repository"
	"todo-list/pkg/utils"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskService interface {
	ListTasks(ctx context.Context) ([]*model.Task, error)
	Create(ctx context.Context, req *dto.CreateTaskReq) (string, error)
	Update(ctx context.Context, id string, req *dto.UpdateTaskReq) (*model.Task, error)
	Delete(ctx context.Context, id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) ListTasks(ctx context.Context) ([]*model.Task, error) {
	return s.repo.ListTasks(ctx)
}

func (s *taskService) Create(ctx context.Context, req *dto.CreateTaskReq) (string, error) {
	var task model.Task
	utils.Copy(&req, &task)
	return s.repo.Create(ctx, &task)
}

func (s *taskService) Update(ctx context.Context, id string, req *dto.UpdateTaskReq) (*model.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	return s.repo.Update(ctx, task, req.Completed)
}

func (s *taskService) Delete(ctx context.Context, id string) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrTaskNotFound
		}
		return err
	}

	return s.repo.Delete(ctx, id)
}
