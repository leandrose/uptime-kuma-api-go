package entities

type UptimeKumaWebService struct {
	Sid          *string `json:"sid,omitempty"`
	Upgrades     *string `json:"upgrades,omitempty"`
	PingInterval *int    `json:"pingInterval,omitempty"`
	PingTimeout  *int    `json:"pingTimeout,omitempty"`
	MaxPayload   *int64  `json:"maxPayload,omitempty"`
}
