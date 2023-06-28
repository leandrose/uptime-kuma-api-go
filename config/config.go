package config

import (
	"github.com/leandrose/uptime-kuma-api-go/domain/entities/configs"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	configs.HttpConfig
	configs.LoggerConfig

	Database   configs.DatabaseConfig
	UptimeKuma configs.UptimeKumaConfig
}

var config Config

func LoadConfig() {
	var err error

	if _, err := os.Stat(".env"); err == nil {
		if err = godotenv.Load(); err != nil {
			panic(err.Error())
		}
	}
	if err = loadVariables(); err != nil {
		panic(err.Error())
	}
	if err = loadLogrus(); err != nil {
		panic(err.Error())
	}
}

func GetConfig() *Config {
	return &config
}

func loadVariables() (err error) {
	config.CrossOrigin = strings.Split(os.Getenv("CROSS_ORIGIN"), ",")
	config.DebugMode, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	config.TraceMode, _ = strconv.ParseBool(os.Getenv("TRACE"))

	// Database
	config.Database.Type = os.Getenv("DATABASE_TYPE")
	config.Database.Hostname = os.Getenv("DATABASE_HOSTNAME")
	if len(os.Getenv("DATABASE_PORT")) > 0 {
		port, err := strconv.ParseInt(os.Getenv("DATABASE_PORT"), 10, 32)
		if err == nil {
			config.Database.Port = int(port)
		}
	}
	config.Database.Database = os.Getenv("DATABASE_DATABASE")
	config.Database.Username = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")

	config.UptimeKuma.Uri = os.Getenv("UPTIMEKUMA_URI")
	config.UptimeKuma.Username = os.Getenv("UPTIMEKUMA_USERNAME")
	config.UptimeKuma.Password = os.Getenv("UPTIMEKUMA_PASSWORD")

	return
}

func loadLogrus() (err error) {
	level := logrus.WarnLevel
	if config.DebugMode {
		level = logrus.DebugLevel
	}
	if config.TraceMode {
		level = logrus.TraceLevel
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(level)

	return
}
