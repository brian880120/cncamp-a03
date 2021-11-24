package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal, Err: %s", err)
	}

	w.Write(jsonResp)
}
