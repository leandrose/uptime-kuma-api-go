package http

import (
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"net/http"
)

func UptimesPresenter(w http.ResponseWriter, items []entities.Uptime) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(items)
	_, _ = w.Write(b)

	return nil
}
