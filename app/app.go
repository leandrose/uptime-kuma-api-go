package app

import (
	"github.com/leandrose/uptime-kuma-api-go/app/providers"
	"github.com/leandrose/uptime-kuma-api-go/config"
	"github.com/leandrose/uptime-kuma-api-go/infra/uptimekuma"
)

func InitializeApp() {
	providers.LoadProviders()
	uptimekuma.LoadUptimeKumaWebService(config.GetConfig().UptimeKuma)
}
