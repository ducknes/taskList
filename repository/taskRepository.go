package repository

import (
	_ "embed"
	"fmt"
	"taskList/service/models"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

//go:embed sql/getAllTasks.sql
var getAllTasks string

func (t *TaskRepository) GetAllTasks() ([]models.Task, error) {
	rows, err := t.db.Query(getAllTasks)
	if err != nil {
		return nil, fmt.Errorf("err while getting all tasks. err %v", err)
	}

	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		task := models.Task{}
		if err = rows.Scan(&task.Id, &task.Status, &task.Message); err != nil {
			return nil, fmt.Errorf("err while scaning all tasks. err %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

//go:embed sql/getTaskById.sql
var getTask string

func (t *TaskRepository) GetTask(id int) (models.Task, error) {
	row := t.db.QueryRow(getTask, id)

	task := models.Task{}
	if err := row.Scan(&task.Id, &task.Status, &task.Message); err != nil {
		return task, fmt.Errorf("err white scaning one task. err: %v", err)
	}

	return task, nil
}

//go:embed sql/updateTaskById.sql
var updateTask string

func (t *TaskRepository) UpdateTask(id int) (int, error) {
	var updatedTaskId int
	err := t.db.QueryRowx(updateTask, id).Scan(&updatedTaskId)
	if err != nil {
		return 0, fmt.Errorf("err while updating task. err: %v", err)
	}

	return updatedTaskId, nil
}

//go:embed sql/addTask.sql
var addTask string

func (t *TaskRepository) AddTask(task models.Task) (int, error) {
	var addedTaskId int
	err := t.db.QueryRowx(addTask, task.Status, task.Message).Scan(&addedTaskId)
	if err != nil {
		return 0, fmt.Errorf("err while adding new task. err: %v", err)
	}

	return addedTaskId, nil
}

//go:embed sql/deleteTaskById.sql
var deleteTask string

func (t *TaskRepository) DeleteTask(id int) (int, error) {
	var deletedTaskId int
	err := t.db.QueryRowx(deleteTask, id).Scan(&deletedTaskId)
	if err != nil {
		return 0, fmt.Errorf("err while deleting task. err: %v", err)
	}

	return deletedTaskId, nil
}
