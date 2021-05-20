package handler

import (
	"encoding/json"
	"net/http"
	"github.com/paharip/microservice-tutorial"
)

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var req helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
	}
	h.next.ServeHTTP(rw, r)
}
