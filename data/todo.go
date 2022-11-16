// Created by Keenan
package data

import "database/sql"

type TodoItems interface {
	GetIncompleteTodoItems(userId string) ([]*TodoItem, error)
	CreateTodoItem(userID, todoItem string) (*TodoItem, error)
}

type TodoItem struct {
	ID          string
	UserID      string
	TodoItem    string
	IsCompleted bool
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}

// var todoItems []TodoItems

//
func (c *client) GetIncompleteTodoItems(userId string) ([]*TodoItem, error) {
	stmt := `SELECT id, user_id, todo_item, is_completed, created_at, updated_at FROM todo_items
		WHERE id = $1 AND is_completed = $2;`

	rows, err := c.db.Query(stmt, userId, false)

	var todoItems []*TodoItem

	for rows.Next() {
		todoItem := new(TodoItem)
		rows.Scan(&todoItem.ID, &todoItem.UserID,
			&todoItem.TodoItem, &todoItem.IsCompleted,
			&todoItem.CreatedAt, &todoItem.UpdatedAt)

		todoItems = append(todoItems, todoItem)
	}

	return todoItems, err
}

func (c *client) CreateTodoItem(userID string, todoItemName string) (*TodoItem, error) {
	stmt := `INSERT INTO todo_items (user_id, todo_item)
		VALUES ($1, $2)
		RETURNING id, user_id, todo_item, is_completed, created_at, updated_at`

	todoItem := new(TodoItem)
	err := c.db.QueryRow(stmt, userID, todoItemName).Scan(&todoItem.ID, &todoItem.UserID,
		&todoItem.TodoItem, &todoItem.IsCompleted,
		&todoItem.CreatedAt, &todoItem.UpdatedAt)

	return todoItem, err
}
