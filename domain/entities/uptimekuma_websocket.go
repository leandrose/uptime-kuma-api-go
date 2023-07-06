package entities

type UptimeKumaWebService struct {
	Sid *string `json:"sid,omitempty"`
}

type UptimeKumaInfo struct {
	Version              string  `json:"version"`
	LatestVersion        string  `json:"latestVersion"`
	PrimaryBaseURL       *string `json:"primaryBaseURL"`
	ServerTimezone       string  `json:"serverTimezone"`
	ServerTimezoneOffset string  `json:"serverTimezoneOffset"`
}

type Uptime struct {
	ID       int     `json:"id"`
	Duration int     `json:"duration"`
	Uptime   float64 `json:"uptime"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token" default:""`
}

type Event struct {
	Ok        bool    `json:"ok"`
	Token     *string `json:"token,omitempty"`
	Msg       *string `json:"msg,omitempty"`
	MonitorId *int    `json:"monitorId,omitempty"`
	Tag       *Tag    `json:"tag"`
	Tags      *[]Tag  `json:"tags"`
}

type Tag struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func NewMonitor() Monitor {
	return Monitor{
		Active:              true,
		Interval:            60,
		Maxretries:          1,
		IgnoreTls:           false,
		UpsideDown:          false,
		Maxredirects:        10,
		DnsResolveType:      "A",
		DnsResolveServer:    "1.1.1.1",
		RetryInterval:       60,
		Method:              "GET",
		DockerContainer:     "",
		ExpiryNotification:  false,
		ResendInterval:      0,
		PacketSize:          56,
		HttpBodyEncoding:    "json",
		AcceptedStatuscodes: []string{"200-299"},
		MqttUsername:        "",
		MqttPassword:        "",
		MqttTopic:           "",
		MqttSuccessMessage:  "",
	}
}

type Monitor struct {
	ID                 *int         `json:"id,omitempty"`
	NotificationIDList map[int]bool `json:"notificationIDList"`
	//Tags                     []Tags               `json:"tags,omitempty"`
	Name                     string   `json:"name"`
	Active                   bool     `json:"active"`
	Interval                 int      `json:"interval"`
	Url                      *string  `json:"url,omitempty"`
	Type                     string   `json:"type"`
	Weight                   *int     `json:"weight,omitempty"`
	Hostname                 *string  `json:"hostname,omitempty"`
	Port                     *int     `json:"port,omitempty"`
	CreatedDate              *string  `json:"createdDate,omitempty"`
	Keyword                  *string  `json:"keyword,omitempty"`
	Maxretries               int      `json:"maxretries"`
	IgnoreTls                bool     `json:"ignoreTls"`
	UpsideDown               bool     `json:"upsideDown"`
	Maxredirects             int      `json:"maxredirects"`
	AcceptedStatuscodes      []string `json:"accepted_statuscodes"`
	DnsResolveType           string   `json:"dns_resolve_type"`
	DnsResolveServer         string   `json:"dns_resolve_server"`
	DnsLastResult            *string  `json:"dns_last_result,omitempty"`
	RetryInterval            int      `json:"retryInterval"`
	PushToken                *string  `json:"pushToken,omitempty"`
	Method                   string   `json:"method"`
	Body                     *string  `json:"body,omitempty"`
	Headers                  *string  `json:"headers,omitempty"`
	BasicAuthUser            *string  `json:"basicAuthUser,omitempty"`
	BasicAuthPass            *string  `json:"basicAuthPass,omitempty"`
	DockerHost               *int     `json:"docker_host"`
	DockerContainer          string   `json:"docker_container"`
	ProxyId                  *int     `json:"proxyId"`
	ExpiryNotification       bool     `json:"expiryNotification"`
	MqttTopic                string   `json:"mqttTopic"`
	MqttSuccessMessage       string   `json:"mqttSuccessMessage"`
	MqttUsername             string   `json:"mqttUsername"`
	MqttPassword             string   `json:"mqttPassword"`
	DatabaseConnectionString *string  `json:"databaseConnectionString,omitempty"`
	DatabaseQuery            *string  `json:"databaseQuery,omitempty"`
	AuthMethod               *string  `json:"authMethod"`
	AuthDomain               *string  `json:"authDomain,omitempty"`
	AuthWorkstation          *string  `json:"authWorkstation,omitempty"`
	GrpcUrl                  *string  `json:"grpcUrl,omitempty"`
	GrpcProtobuf             *string  `json:"grpcProtobuf,omitempty"`
	GrpcBody                 *string  `json:"grpcBody,omitempty"`
	GrpcMetadata             *string  `json:"grpcMetadata,omitempty"`
	GrpcMethod               *string  `json:"grpcMethod,omitempty"`
	GrpcServiceName          *string  `json:"grpcServiceName,omitempty"`
	GrpcEnableTls            *bool    `json:"grpcEnableTls,omitempty"`
	RadiusUsername           *string  `json:"radiusUsername,omitempty"`
	RadiusPassword           *string  `json:"radiusPassword,omitempty"`
	RadiusCallingStationId   *string  `json:"radiusCallingStationId,omitempty"`
	RadiusCalledStationId    *string  `json:"radiusCalledStationId,omitempty"`
	RadiusSecret             *string  `json:"radiusSecret,omitempty"`
	ResendInterval           int      `json:"resendInterval"`
	PacketSize               int      `json:"packetSize"`
	Game                     *string  `json:"game,omitempty"`
	HttpBodyEncoding         string   `json:"httpBodyEncoding"`
	Description              *string  `json:"description,omitempty"`
	TlsCa                    *string  `json:"tlsCa,omitempty"`
	TlsCert                  *string  `json:"tlsCert,omitempty"`
	TlsKey                   *string  `json:"tlsKey,omitempty"`
}
