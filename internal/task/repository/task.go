package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"todo-list/internal/task/model"
)

var ErrNotFound = errors.New("record not found")

var tasks []*model.Task // in-memory db

type TaskRepository interface {
	ListTasks(ctx context.Context) ([]*model.Task, error)
	Create(ctx context.Context, task *model.Task) (string, error)
	GetByID(ctx context.Context, id string) (*model.Task, error)
	Update(ctx context.Context, task *model.Task, completed bool) (*model.Task, error)
	Delete(ctx context.Context, id string) error
}

type taskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (r *taskRepository) ListTasks(ctx context.Context) ([]*model.Task, error) {
	return tasks, nil
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) (string, error) {
	task.ID = uuid.New().String()
	task.Completed = false
	tasks = append(tasks, task)
	return task.ID, nil
}

func (r *taskRepository) GetByID(ctx context.Context, id string) (*model.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, ErrNotFound
}

func (r *taskRepository) Update(ctx context.Context, task *model.Task, completed bool) (*model.Task, error) {
	task.Completed = completed
	return task, nil
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return nil
}
