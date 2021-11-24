package handler

import (
	"net/http"
)

func NotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
    http.NotFound(w, req)
    return
}
