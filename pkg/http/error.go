package http

import (
	"encoding/json"
	"net/http"
)

func SendHttpError(w http.ResponseWriter, err error) {

}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	converted, _ := json.Marshal(data)
	w.Write(converted)
}
