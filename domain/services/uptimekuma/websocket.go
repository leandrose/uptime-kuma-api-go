package uptimekuma

import (
	"context"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"nhooyr.io/websocket"
)

// Handle para ouvir eventos recebidos pelo WebSocket, o data recebido retornara em string
type Handle func(args ...interface{})

type IUptimeKumaWebSocket interface {
	Send()
	IsAuthenticated() bool
	SetToken(token string)
	Close()
	Consume(ctx context.Context, cancel context.CancelFunc)
	On(event string, handleFunc Handle)
	WriteText(ctx context.Context, b []byte) (c chan entities.Event, err error)
	Write(ctx context.Context, messageType websocket.MessageType, b []byte) error
}
