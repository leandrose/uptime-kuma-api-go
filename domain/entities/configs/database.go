package configs

type DatabaseConfig struct {
	Type         string
	Hostname     string
	Port         int
	Database     string
	Scheme       string
	Username     string
	Password     string
	PanicInError bool
}

const (
	DATABASE_TYPE_MYSQL     string = "mysql"
	DATABASE_TYPE_SQLITE    string = "sqlite"
	DATABASE_TYPE_POSTGRES  string = "postgres"
	DATABASE_TYPE_SQLSERVER string = "sqlserver"
)
