package providers

import (
	"github.com/golobby/container/v3"
	"github.com/leandrose/uptime-kuma-api-go/domain/services"
	"github.com/leandrose/uptime-kuma-api-go/infra/uptimekuma"
)

func LoadProviders() {
	// DATABASE
	//_ = container.Singleton(func() *gorm.DB {
	//	return databases.NewDatabase(config.GetConfig().Database)
	//})
	// UPTIME KUMA
	_ = container.Singleton(func() *services.IUptimeKuma {
		return uptimekuma.GetInstanceUptimeKuma()
	})
}
