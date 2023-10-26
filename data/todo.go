// Created by Keenan Sanford
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

// Gets all incomplete (is_completed = false) todoItems using the provided userId.
func (c *client) GetIncompleteTodoItems(userId string) ([]*TodoItem, error) {
	stmt := `SELECT id, user_id, todo_item, is_completed, created_at, updated_at FROM todo_items
		WHERE user_id = $1 AND is_completed IS NOT TRUE;`

	rows, err := c.db.Query(stmt, userId)

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

// Created todo
func (c *client) CreateTodoItem(userID string, todoItemName string) (*TodoItem, error) {
	stmt := `INSERT INTO todo_items (user_id, todo_item)
		VALUES ($1, $2)
		RETURNING id, user_id, todo_item, is_completed, created_at, updated_at;`

	todoItem := new(TodoItem)
	err := c.db.QueryRow(stmt, userID, todoItemName).Scan(&todoItem.ID, &todoItem.UserID,
		&todoItem.TodoItem, &todoItem.IsCompleted,
		&todoItem.CreatedAt, &todoItem.UpdatedAt)

	return todoItem, err
}
