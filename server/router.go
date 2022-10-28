package server

import "github.com/gorilla/mux"

// builds a router will all of the necicary endpoints
// if the handler needs the db or any clients stored in the app use an app method for a handler
func ConfigureRouter(a *app) *mux.Router {
	r := mux.NewRouter()
	// add a healthchek endpoint on the base router
	r.HandleFunc("/health", a.HealthCheck)

	// use a subrouter for add the path prefix to all api related endpoints
	sub := r.PathPrefix("/api/v1").Subrouter()
	// use a user subrouter for all user related endpoints
	user := sub.PathPrefix("/user").Subrouter()
	user.HandleFunc("/{id}", a.GetUserByID).Methods("GET")
	user.HandleFunc("/create", a.CreateUser).Methods("POST")
	user.HandleFunc("/update/{id}", a.UpdateUser).Methods("PATCH")

	return r
}
