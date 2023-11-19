package service

import "taskList/repository"

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetAllTasks() {}

func (s *TaskService) GetTask() {}

func (s *TaskService) AddTask() {}

func (s *TaskService) UpdateTask() {}

func (s *TaskService) DeleteTask() {}
