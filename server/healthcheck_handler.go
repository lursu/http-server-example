package server

import (
	"http-server-example/view"
	"net/http"
	"time"
)

func (a *app) HealthCheck(w http.ResponseWriter, r *http.Request) {
	err := a.Healthy()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "client health checks failed error: ", err)
		return
	}

	writeSuccess(w, view.HealthCheck{
		Uptime: time.Since(a.start),
		Msg:    "All healthchecks are good",
	})
}

// func (a *app) Greeting(w http.ResponseWriter, r *http.Request) {
// 	writeSuccess(w, "Hello world!")
// }
