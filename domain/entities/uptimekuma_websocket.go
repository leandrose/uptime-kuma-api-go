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

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token" default:""`
}

type Event struct {
	Ok    bool    `json:"ok"`
	Token *string `json:"token,omitempty"`
	Msg   *string `json:"msg,omitempty"`
}

type Monitor struct {
	ID *int `json:"id,omitempty"`
	//notificationIDList	{}
	//tags	[]
	Name                     string   `json:"name"`
	Active                   bool     `json:"active" default:"1"`
	Interval                 int      `json:"interval" default:"20"`
	Url                      *string  `json:"url"`
	Type                     string   `json:"type"`
	Weight                   int      `json:"weight" default:"2000"`
	Hostname                 string   `json:"hostname"`
	Port                     int      `json:"port"`
	CreatedDate              string   `json:"created_date"`
	Keyword                  *string  `json:"keyword"`
	Maxretries               int      `json:"maxretries" default:"1"`
	IgnoreTls                bool     `json:"ignore_tls" default:"0"`
	UpsideDown               bool     `json:"upside_down" default:"0"`
	Maxredirects             int      `json:"maxredirects" default:"10"`
	AcceptedStatuscodes      []string `json:"accepted_statuscodes"`
	DnsResolveType           string   `json:"dns_resolve_type" default:"A"`
	DnsResolveServer         string   `json:"dns_resolve_server" default:"1.1.1.1"`
	DnsLastResult            *string  `json:"dns_last_result"`
	RetryInterval            int      `json:"retry_interval" default:"0"`
	PushToken                *string  `json:"push_token"`
	Method                   string   `json:"METHOD" default:"GET"`
	Body                     *string  `json:"BODY"`
	Headers                  *string  `json:"headers"`
	BasicAuthUser            *string  `json:"basic_auth_user"`
	BasicAuthPass            *string  `json:"basic_auth_pass"`
	DockerHost               *int     `json:"docker_host"`
	DockerContainer          string   `json:"docker_container" default:""`
	ProxyId                  int      `json:"proxy_id"`
	ExpiryNotification       bool     `json:"expiry_notification" default:"1"`
	MqttTopic                *string  `json:"mqtt_topic"`
	MqttSuccessMessage       *string  `json:"mqtt_success_message"`
	MqttUsername             *string  `json:"mqtt_username"`
	MqttPassword             *string  `json:"mqtt_password"`
	DatabaseConnectionString *string  `json:"database_connection_string"`
	DatabaseQuery            *string  `json:"database_query"`
	AuthMethod               string   `json:"auth_method" default:""`
	AuthDomain               *string  `json:"auth_domain"`
	AuthWorkstation          *string  `json:"auth_workstation"`
	GrpcUrl                  *string  `json:"grpc_url"`
	GrpcProtobuf             *string  `json:"grpc_protobuf"`
	GrpcBody                 *string  `json:"grpc_body"`
	GrpcMetadata             *string  `json:"grpc_metadata"`
	GrpcMethod               *string  `json:"grpc_method"`
	GrpcServiceName          *string  `json:"grpc_service_name"`
	GrpcEnableTls            bool     `json:"grpc_enable_tls" default:"0"`
	RadiusUsername           *string  `json:"radius_username"`
	RadiusPassword           *string  `json:"radius_password"`
	RadiusCallingStationId   *string  `json:"radius_calling_station_id"`
	RadiusCalledStationId    *string  `json:"radius_called_station_id"`
	RadiusSecret             *string  `json:"radius_secret"`
	ResendInterval           int      `json:"resend_interval" default:"0"`
	PacketSize               int      `json:"packet_size" default:"56"`
	Game                     *string  `json:"game"`
	HttpBodyEncoding         string   `json:"http_body_encoding" default:"json"`
	Description              *string  `json:"description"`
	TlsCa                    *string  `json:"tls_ca"`
	TlsCert                  *string  `json:"tls_cert"`
	TlsKey                   *string  `json:"tls_key"`
}
