// Created by Keenan Sanford
package view

import "time"

type TodoItem struct {
	ID          string
	UserID      string
	TodoItem    string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TodoItems struct {
	TodoItems []*TodoItem
}
