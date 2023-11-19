package repository

import (
	"github.com/jmoiron/sqlx"
	"taskList/service/models"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) GetAllTasks() {}

func (t *TaskRepository) GetTask(id int) (models.Task, error) {
	return models.Task{}, nil
}

func (t *TaskRepository) UpdateTask() {}

func (t *TaskRepository) AddTask() {}

func (t *TaskRepository) DeleteTask() {}
