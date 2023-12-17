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

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTask(id string) (models.Task, error) {
	taskId, err := validateId(id)
	if err != nil {
		return models.Task{}, err
	}

	return s.repo.GetTask(taskId)
}

func (s *TaskService) AddTask(task models.Task) (int, error) {
	return s.repo.AddTask(task)
}

func (s *TaskService) UpdateTask(task models.Task) (int, error) {
	return s.repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id string) (int, error) {
	taskId, err := validateId(id)
	if err != nil {
		return 0, err
	}

	return s.repo.DeleteTask(taskId)
}
