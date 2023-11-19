package service

import (
	"taskList/repository"
	"taskList/service/models"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetAllTasks() {}

func (s *TaskService) GetTask(id string) (models.Task, error) {
	taskId, err := validateId(id)
	if err != nil {
		return models.Task{}, err
	}

	return s.repo.GetTask(taskId)
}

func (s *TaskService) AddTask() {}

func (s *TaskService) UpdateTask() {}

func (s *TaskService) DeleteTask() {}
