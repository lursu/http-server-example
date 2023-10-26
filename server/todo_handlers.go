// Created by Keenan
package server

import (
	"encoding/json"
	"fmt"
	"http-server-example/data"
	"http-server-example/view"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *app) GetIncompleteTodoItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	log.Println(vars)
	log.Println(userID)

	todos, err := a.db.GetIncompleteTodoItems(userID)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Unable to find todos for userId: %s error: ", userID), err)
		return
	}

	viewTodos := convertTodoItems(todos)

	writeSuccess(w, &view.TodoItems{
		TodoItems: viewTodos,
	})
}

func (a *app) CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todoItem view.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&todoItem); err != nil {
		writeError(w, http.StatusBadRequest, "The body posted does not match the spec unable to marshal error: ", err)
		return
	}

	todos, err := a.db.CreateTodoItem(todoItem.UserID, todoItem.TodoItem)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Unable to create todo %s for userId: %s error: ", todoItem.TodoItem, todoItem.UserID), err)
		return
	}

	writeSuccess(w, &view.TodoItem{
		ID:          todos.ID,
		UserID:      todos.UserID,
		TodoItem:    todos.TodoItem,
		IsCompleted: todos.IsCompleted,
		UpdatedAt:   todos.UpdatedAt.Time,
		CreatedAt:   todos.CreatedAt.Time,
	})
}

// Helper function converts data TodoItem -> view TodoItem
func convertTodoItems(todoItems []*data.TodoItem) []*view.TodoItem {
	result := make([]*view.TodoItem, len(todoItems))

	for i, item := range todoItems {
		result[i] = &view.TodoItem{
			ID:          item.ID,
			UserID:      item.UserID,
			TodoItem:    item.TodoItem,
			IsCompleted: item.IsCompleted,
			CreatedAt:   item.CreatedAt.Time,
			UpdatedAt:   item.UpdatedAt.Time,
		}
	}

	return result
}
