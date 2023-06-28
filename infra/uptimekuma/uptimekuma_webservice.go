package uptimekuma

import (
	"context"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities/configs"
	"github.com/leandrose/uptime-kuma-api-go/domain/services"
	"log"
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

		log.Println("Failed to reconnect:", err)
	}
	s.conn = newConn

	for {
		msg, b, err := s.conn.Read(ctx)
		if err != nil {
			log.Println("Connection closed: ", err)
			break
		}

		// PROCESSAR
		log.Printf("READ msg: %s", msg)
		log.Printf("READ b: %s", string(b))
	}

	go s.Consume(ctx)
}
