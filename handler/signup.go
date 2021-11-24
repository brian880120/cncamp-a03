package handler

import (
	"encoding/json"
	"net/http"
)

type SignupUser struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

func Signup(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var user SignupUser
    err := decoder.Decode(&user)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
    }

    jsonResp, err := json.Marshal(user)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
    }

    w.Write(jsonResp)
}
