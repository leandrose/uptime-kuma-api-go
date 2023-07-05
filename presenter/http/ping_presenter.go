package http

import (
	"encoding/json"
	"net/http"
)

type avgPingPresenter struct {
	ID   int `json:"id"`
	Ping int `json:"ping"`
}

func PingPresenter(w http.ResponseWriter, monitorID int, avgPing int) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(avgPingPresenter{ID: monitorID, Ping: avgPing})
	_, _ = w.Write(b)
	return nil
}

func PingsPresenter(w http.ResponseWriter, avgPings map[int]int) error {
	items := []avgPingPresenter{}
	for k, v := range avgPings {
		items = append(items, avgPingPresenter{ID: k, Ping: v})
	}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(items)
	_, _ = w.Write(b)
	return nil
}
