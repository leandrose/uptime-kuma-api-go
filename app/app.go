package app

import (
	"github.com/leandrose/uptime-kuma-api-go/app/providers"
)

func InitializeApp() {
	providers.LoadProviders()
}
