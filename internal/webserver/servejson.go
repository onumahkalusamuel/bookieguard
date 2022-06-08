package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/onumahkalusamuel/bookieguard/config"
)

func ServeJSON(w http.ResponseWriter, responseBody config.BodyStructure) {

	m, err := json.Marshal(responseBody)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&m)
	if err != nil {
		return
	}

	return
}
