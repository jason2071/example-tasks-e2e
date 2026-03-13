package repository

import (
	"database/sql"
	"example-tasks/model"
	"example-tasks/utils"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type TaskRepository interface {
	CreateTask(task model.TaskRequest) (bool, error)
	GetTasks(cursor int64, size, priority int, sortWith, sortBy string) ([]model.Task, error)
	GetTaskByID(id int64) (model.Task, error)
	UpdateTask(id int64, task model.TaskRequest) (string, error)
	DeleteTask(id int64) (string, error)
}

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{
		db: db,
	}
}

func (r *TaskRepositoryImpl) CreateTask(task model.TaskRequest) (bool, error) {
	query := `
		INSERT INTO example.tasks (title, status, priority) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (title) DO NOTHING 
		RETURNING true;
	`

	var inserted bool
	err := r.db.QueryRow(query, *task.Title, *task.Status, *task.Priority).Scan(&inserted)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Task with the same title already exists:", *task.Title)
			return false, nil
		}
		log.Println("Failed to create task:", err)
		return false, err
	}

	log.Println("Task created successfully:", *task.Title)
	return inserted, nil
}

func (r *TaskRepositoryImpl) GetTasks(cursor int64, size int, priority int, sortWith, sortBy string) ([]model.Task, error) {
	query := `
        SELECT id, title, status, created_at, updated_at, priority 
        FROM example.tasks 
    `

	args := []interface{}{}
	argIndex := 1

	if cursor > 0 {
		operator := ">"
		if sortBy == "desc" {
			operator = "<"
		}

		if priority > 0 {
			query += fmt.Sprintf(" AND id %s $%d ", operator, argIndex)
		} else {
			query += fmt.Sprintf(" WHERE id %s $%d ", operator, argIndex)
		}
		args = append(args, cursor)
		argIndex++
	}

	query += fmt.Sprintf(" ORDER BY %s %s ", sortWith, sortBy)
	query += fmt.Sprintf(" LIMIT $%d; ", argIndex)
	args = append(args, size)

	var tasks []model.Task

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Println("Failed to retrieve tasks:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt, &task.UpdatedAt, &task.Priority); err != nil {
			log.Println("Failed to scan task row:", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Println("Row iteration error:", err)
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) taskTotalRecords(priority int) (int64, error) {
	var totalRows int64
	countQuery := "SELECT COUNT(*) FROM example.tasks"

	args := []interface{}{}
	argIndex := 1

	if priority > 0 {
		countQuery += fmt.Sprintf(" WHERE priority = $%d", argIndex)
		args = append(args, priority)
		argIndex++
	}

	err := r.db.QueryRow(countQuery, args...).Scan(&totalRows)
	if err != nil {
		log.Println("Failed to count total tasks:", err)
		return 0, err
	}
	log.Println("Total tasks count:", totalRows)
	return totalRows, nil
}

func (r *TaskRepositoryImpl) GetTaskByID(id int64) (model.Task, error) {
	query := `
		SELECT id, title, status, created_at, updated_at, priority
		FROM example.tasks 
		WHERE id = $1;`

	var task model.Task
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt, &task.UpdatedAt, &task.Priority)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Task not found with ID:", id)
			return model.Task{}, nil
		}
		log.Println("Failed to retrieve task:", err)
		return model.Task{}, err
	}
	log.Println("Task retrieved successfully with ID:", id)
	return task, nil
}

func (r *TaskRepositoryImpl) UpdateTask(id int64, task model.TaskRequest) (string, error) {
	query := "update example.tasks set "
	args := []interface{}{}
	argIndex := 1

	if task.Title != nil && *task.Title != "" {
		query += fmt.Sprintf("title = $%d, ", argIndex)
		args = append(args, *task.Title)
		argIndex++
	}

	if task.Status != nil && *task.Status != "" {
		query += fmt.Sprintf("status = $%d, ", argIndex)
		args = append(args, *task.Status)
		argIndex++
	}

	if task.Priority != nil {
		query += fmt.Sprintf("priority = $%d, ", argIndex)
		args = append(args, *task.Priority)
		argIndex++
	}

	query = query[:len(query)-2]
	query += fmt.Sprintf(" WHERE id = $%d", argIndex)
	args = append(args, id)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			log.Println("Duplicate key error:", pqErr.Message)
			return utils.ErrDuplicateEntry.ErrorCode, nil
		}

		log.Println("Failed to update task:", err)
		return utils.ErrInternalServer.ErrorCode, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Failed to get rows affected:", err)
		return utils.ErrInternalServer.ErrorCode, err
	}

	if rowsAffected == 0 {
		log.Println("Task not found with ID:", id)
		return utils.ErrNotFound.ErrorCode, nil
	}

	log.Println("Task updated successfully with ID:", id)
	return utils.Success.ErrorCode, nil
}

func (r *TaskRepositoryImpl) DeleteTask(id int64) (string, error) {
	query := `
		DELETE FROM example.tasks 
		WHERE id = $1;
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Println("Failed to delete task:", err)
		return utils.ErrInvalidRequest.ErrorCode, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Failed to get rows affected:", err)
		return utils.ErrInternalServer.ErrorCode, err
	}

	if rowsAffected == 0 {
		log.Println("Task not found with ID:", id)
		return utils.ErrNotFound.ErrorCode, nil
	}

	log.Println("Task deleted successfully with ID:", id)
	return utils.Success.ErrorCode, nil
}
