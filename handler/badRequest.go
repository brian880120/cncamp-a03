package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func BadRequest(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
    w.Header().Set("Content-Type", "application/json")

    resp := make(map[string]string)

    resp["message"] = "Bad Request"
    jsonResp, err := json.Marshal(resp)

    if err != nil {
        log.Fatalf("Error happend in JSON marshal, Err: %s", err)
    }

    w.Write(jsonResp)
}
