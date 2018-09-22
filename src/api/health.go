package api

import (
	"encoding/json"
	"net/http"
)

type errorStruct struct {
	Message string `json:"error_msg"`
}

// Ping responds with a 200 ok response.
func Ping(w http.ResponseWriter, r *http.Request) {
	type responseStruct struct {
		Status string `json:"status"`
	}

	json.NewEncoder(w).Encode(&responseStruct{Status: "ok"})
}
