package cmd

import (
	"github.com/golobby/container/v3"
	"github.com/leandrose/uptime-kuma-api-go/domain/services/uptimekuma"
	"github.com/leandrose/uptime-kuma-api-go/infra/http"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	serveCmd.Flags().IntVarP(&servePort, "port", "p", 3000, "specify the port number to be used when initializing the web server")
}

var servePort int
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Initialize the Web Server",
	Long:  "Start web server using command-line",
	Run: func(cmd *cobra.Command, args []string) {
		var service uptimekuma.IUptimeKumaService
		_ = container.Resolve(&service)

		s := http.NewServer(servePort)
		if err := s.Start(); err != nil {
			logrus.Errorf("HTTP Server Error: %s", err)
		}
	},
}
