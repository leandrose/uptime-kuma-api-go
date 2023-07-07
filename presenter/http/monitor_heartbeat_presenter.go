package http

import (
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"net/http"
)

type monitorHeartbeatsPresenter struct {
	HeartbeatsPresenter []entities.Heartbeat `json:"heartbeats_presenter"`
}

func MonitorHeartbeatsPresenter(w http.ResponseWriter, heartbeats []entities.Heartbeat) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(monitorHeartbeatsPresenter{HeartbeatsPresenter: heartbeats})
	_, _ = w.Write(b)
	return nil
}
