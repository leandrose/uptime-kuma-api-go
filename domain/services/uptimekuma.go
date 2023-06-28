package services

import (
	"context"
	"nhooyr.io/websocket"
)

type IUptimeKuma interface {
	Send()
	Consume(ctx context.Context)
	Write(ctx context.Context, messageType websocket.MessageType, b []byte) error
}
