package http

import (
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"net/http"
)

func MonitorPresenter(w http.ResponseWriter, monitor entities.Monitor) error {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(monitor)
	if err != nil {
		return Error500Presenter(w, err)
	}
	_, err = w.Write(b)
	if err != nil {
		return Error500Presenter(w, err)
	}
	return nil
}
