package uptimekuma

import "github.com/leandrose/uptime-kuma-api-go/domain/entities"

type IUptimeKumaService interface {
	Auth() bool
	// MONITOR
	GetMonitors() []entities.Monitor
	GetMonitorById(id int) (*entities.Monitor, error)
	CreateMonitor(monitor entities.Monitor) (*entities.Monitor, error)
	DeleteMonitor(monitorId int) error
	EditMonitor(monitor entities.Monitor) (*entities.Monitor, error)
	// PING
	GetPingAverage(int) (*int, error)
	GetPingsAverage() map[int]int
	// UPTIMES
	GetUptime(int) (*[]entities.Uptime, error)
	GetUptimes() []entities.Uptime
}
