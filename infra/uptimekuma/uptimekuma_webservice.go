package uptimekuma

import (
	"context"
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities/configs"
	"github.com/leandrose/uptime-kuma-api-go/domain/services"
	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"time"
)

var uptimeKumaWebServiceInstance *services.IUptimeKuma

type uptimeKumaWebService struct {
	conn   *websocket.Conn
	cancel context.CancelFunc
	config configs.UptimeKumaConfig
}

func LoadUptimeKumaWebService(config configs.UptimeKumaConfig) {
	uptimeKumaWebServiceInstance = NewUptimeKumaWebService(config)
}

func GetInstanceUptimeKuma() *services.IUptimeKuma {
	return uptimeKumaWebServiceInstance
}

func NewUptimeKumaWebService(config configs.UptimeKumaConfig) *services.IUptimeKuma {
	ctx, cancel := context.WithCancel(context.Background())

	var service services.IUptimeKuma
	service = uptimeKumaWebService{
		config: config,
		cancel: cancel,
	}

	go service.Consume(ctx)

	return &service
}

func (s uptimeKumaWebService) Send() {

}

func (s uptimeKumaWebService) Consume(ctx context.Context) {
	var err error
	var newConn *websocket.Conn
	for {
		time.Sleep(5 * time.Second)

		newConn, _, err = websocket.Dial(ctx, s.config.Uri, nil)
		if err == nil {
			break
		}

		logrus.Infof("Failed to reconnect: %s", err)
	}
	s.conn = newConn
	//go s.pong(ctx)

	for {
		msg, b, err := s.conn.Read(ctx)
		if err != nil {
			logrus.Infof("Connection closed: %s", err)
			break
		}

		// PROCESSAR
		bb := string(b[0])
		logrus.Debugf("bb: %+v", string(b[0]))
		switch bb {
		case "0":
			var item entities.UptimeKumaWebService
			err = json.Unmarshal(b[1:], &item)
			if err != nil {
				logrus.Errorf("Error unmarshal: %s", err)
				break
			}
			logrus.Debugf("Received: %+v", item)
			break
		case "2":
			logrus.Infof("PONG")
			_ = s.Write(ctx, websocket.MessageText, []byte("3"))
			break
		}
		logrus.Debugf("READ msg: %s", msg)
		logrus.Debugf("READ b: %s", string(b))
	}

	//s.cancel()
	go s.Consume(ctx)
}

func (s *uptimeKumaWebService) pong(ctx context.Context) {
	select {
	case <-time.After(25 * time.Second):
		_ = s.conn.Write(ctx, websocket.MessageText, []byte("3"))
		break
	case <-ctx.Done():
		return
	}
}

func (s uptimeKumaWebService) Write(ctx context.Context, messageType websocket.MessageType, b []byte) (err error) {
	return s.conn.Write(ctx, messageType, b)
}
