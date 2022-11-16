package server

import "github.com/gorilla/mux"

// builds a router will all of the necicary endpoints
// if the handler needs the db or any clients stored in the app use an app method for a handler
func ConfigureRouter(a *app) *mux.Router {
	r := mux.NewRouter()

	// get a non-404 error when pinging from a browser
	// r.HandleFunc("/", a.Greeting)

	// add a healthchek endpoint on the base router
	r.HandleFunc("/health", a.HealthCheck)

	// use a subrouter for add the path prefix to all api related endpoints
	sub := r.PathPrefix("/api/v1").Subrouter()
	// use a user subrouter for all user related endpoints
	user := sub.PathPrefix("/user").Subrouter()
	user.HandleFunc("/{id}", a.GetUserByID).Methods("GET")
	user.HandleFunc("/create", a.CreateUser).Methods("POST")
	user.HandleFunc("/update/{id}", a.UpdateUser).Methods("PATCH")

	// use a todo subrouter for all todo related endpoints
	todo := sub.PathPrefix("/todoItems").Subrouter()
	todo.HandleFunc("/{userId}", a.GetIncompleteTodoItems).Methods("GET")
	todo.HandleFunc("/create", a.CreateTodoItem).Methods("POST")
	return r
}
