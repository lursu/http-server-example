package server

import (
	"encoding/json"
	"http-server-example/view"
	"log"
	"net/http"
)

func writeSuccess(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	writeJSON(w, body)
}

func writeJSON(w http.ResponseWriter, body interface{}) {
	blob, err := json.Marshal(body)
	if err != nil {
		log.Println("Unable to marshal response JSON error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(blob)
}

func writeError(w http.ResponseWriter, status int, msg string, err error) {
	log.Println(msg, err)
	w.WriteHeader(status)
	writeJSON(w, &view.Failure{
		Error: err.Error(),
		Msg:   msg,
	})
}
