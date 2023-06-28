package main

import (
	"github.com/leandrose/uptime-kuma-api-go/app"
	"github.com/leandrose/uptime-kuma-api-go/cmd"
	"github.com/leandrose/uptime-kuma-api-go/config"
)

func main() {
	config.LoadConfig()
	app.InitializeApp()

	cmd.Execute()
}
