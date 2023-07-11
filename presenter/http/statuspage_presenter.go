package http

import (
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"net/http"
)

type statuspagePresenter struct {
	StatusPages []entities.StatusPage `json:"statuspages"`
}

func StatusPagesPresenter(w http.ResponseWriter, statuspages []entities.StatusPage) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(statuspagePresenter{StatusPages: statuspages})
	_, _ = w.Write(b)
	return nil
}

func StatusPagePresenter(w http.ResponseWriter, statuspage entities.StatusPage) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(statuspage)
	_, _ = w.Write(b)
	return nil
}
