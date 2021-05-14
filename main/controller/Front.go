package controller

import (
	"encoding/json"
	"net/http"
)

func RegisterController() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
