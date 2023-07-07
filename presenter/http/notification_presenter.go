package http

import (
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"net/http"
)

type notificationsPresenter struct {
	Notifications []entities.Notification `json:"notifications"`
}

func NotificationsPresenter(w http.ResponseWriter, notifications []entities.Notification) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(notificationsPresenter{
		Notifications: notifications,
	})
	_, _ = w.Write(b)
	return nil
}

func NotificationPresenter(w http.ResponseWriter, notification entities.Notification) error {
	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(notification)
	_, _ = w.Write(b)
	return nil
}
