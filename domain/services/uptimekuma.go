package services

import "context"

type IUptimeKuma interface {
	Send()
	Consume(ctx context.Context)
}
