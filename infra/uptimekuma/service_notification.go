package uptimekuma

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/sirupsen/logrus"
	"time"
)

func (s *uptimekumaService) OnNotificationList(args ...interface{}) {
	log := logrus.WithField("func", "OnNotificationList")
	err := s.fill(args[0], &s.notifications)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debugf("count notifications: %d", len(s.notifications))
	log.Tracef("notifications=%+v", s.notifications)
}

func (s *uptimekumaService) GetNotifications() []entities.Notification {
	return s.notifications
}

func (s *uptimekumaService) CreateNotification(notification entities.Notification) (*entities.Notification, error) {
	//423["addNotification",{"name":"My Webhook Alert (1)","type":"webhook","isDefault":false,"webhookURL":"https://sistema.smsnet.com.br/webhook/uptimekuma","webhookContentType":"json"},null]
	//433[{"ok":true,"msg":"Saved","id":1}]
	notificationID := notification.Id
	log := logrus.WithField("func", "CreateNotification")

	b, _ := json.Marshal([]interface{}{
		"addNotification",
		notification,
		notificationID,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		log.Error(err)
	}

	select {
	case ok := <-c:
		if ok.Ok {
			for _, v := range s.notifications {
				if *v.Id == *ok.ID {
					return &v, nil
				}
			}
		}
	case <-time.After(5 * time.Second):
		log.Errorf("expired request: notification=%+v", notification)
		return nil, errors.New("expired request")
	}

	return nil, errors.New("error ocurred")
}

func (s *uptimekumaService) DeleteNotification(notificationID int) error {
	// 4210["deleteNotification",1]
	// 4310[{"ok":true,"msg":"Deleted"}]
	log := logrus.WithField("func", "DeleteNotification")

	b, _ := json.Marshal([]interface{}{
		"deleteNotification",
		notificationID,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		log.Error(err)
		return err
	}

	select {
	case ok := <-c:
		if ok.Ok {
			return nil
		} else {
			return errors.New(fmt.Sprintf("%s", ok.Msg))
		}
	case <-time.After(5 * time.Second):
		log.Error("expired request")
		return errors.New("expired request")
	}
}
