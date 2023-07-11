package uptimekuma

import "github.com/leandrose/uptime-kuma-api-go/domain/entities"

type IUptimeKumaService interface {
	Auth() bool
	// HEARTBEAT
	GetMonitorHeartbets(int) []entities.Heartbeat
	// MONITOR
	GetMonitors() []entities.Monitor
	GetMonitorById(id int) (*entities.Monitor, error)
	CreateMonitor(monitor entities.Monitor) (*entities.Monitor, error)
	DeleteMonitor(monitorId int) error
	EditMonitor(monitor entities.Monitor) (*entities.Monitor, error)
	ResumeMonitor(int) error
	PauseMonitor(int) error
	AddTagInMonitor(monitorID, tagID int) error
	DeleteTagInMonitor(monitorID, tagID int) error
	// NOTIFICATIONS
	GetNotifications() []entities.Notification
	CreateNotification(entities.Notification) (*entities.Notification, error)
	DeleteNotification(int) error
	// PING
	GetPingAverage(int) (*int, error)
	GetPingsAverage() map[int]int
	// UPTIMES
	GetUptime(int) (*[]entities.Uptime, error)
	GetUptimes() []entities.Uptime
	// STATUS PAGES
	GetStatusPages() []entities.StatusPage
	GetStatusPage(string) (*entities.StatusPage, error)
	DeleteStatusPage(string) error
	CreateStatusPage(entities.StatusPage) (*entities.StatusPage, error)
	UpdateStatusPage(string, entities.StatusPage) (*entities.StatusPage, error)
	// TAGS
	GetTag(int) (*entities.Tag, error)
	GetTags() (*[]entities.Tag, error)
	CreateTag(entities.Tag) (*entities.Tag, error)
	UpdateTag(tag entities.Tag) (*entities.Tag, error)
	DeleteTag(int) error
}
