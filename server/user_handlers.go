package server

import (
	"encoding/json"
	"fmt"
	"http-server-example/view"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *app) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// gets the url path vars out of the request
	vars := mux.Vars(r)
	userID := vars["id"]
	log.Println(vars)
	log.Println(userID)

	user, err := a.db.UserByID(userID)
	if err != nil {
		writeError(w, http.StatusNotFound, fmt.Sprintf("Unable to find user for id: %s error: ", userID), err)
		return
	}

	writeSuccess(w, &view.User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	})
}

func (a *app) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user view.UserUpdate
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeError(w, http.StatusBadRequest, "The body posted does not match the spec unable to marshal error: ", err)
		return
	}

	id, err := a.db.CreateUser(user.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Unable to write user to the db error: ", err)
		return
	}

	writeSuccess(w, &view.User{
		ID:   id,
		Name: user.Name,
	})
}

func (a *app) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// gets the url path vars out of the request
	vars := mux.Vars(r)
	userID := vars["id"]

	defer r.Body.Close()
	var update *view.UserUpdate
	if err := json.NewDecoder(r.Body).Decode(update); err != nil {
		writeError(w, http.StatusBadRequest, "The body posted does not match the spec unable to marshal error: ", err)
		return
	}

	user, err := a.db.UpdateUser(userID, update.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to update user with id: %s error: ", userID), err)
	}

	writeSuccess(w, &view.User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	})
}
