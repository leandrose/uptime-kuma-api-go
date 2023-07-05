package http

import (
	"encoding/json"
	"net/http"
)

type successPresenter struct {
	Success bool `json:"success"`
}

func SuccessPresenter(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(successPresenter{
		Success: true,
	})
	_, _ = w.Write(b)
	return nil
}
