package databases

import (
	"fmt"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDatabase(config configs.DatabaseConfig) (db *gorm.DB) {
	var err error
	var port int
	var dsn string
	log := logrus.WithField("package", "infra/databases")

	log.Debugf("Database connecting in \"%s\"", config.Type)
	switch config.Type {
	case configs.DATABASE_TYPE_MYSQL:
		port = 3306
		if config.Port > 1 {
			port = config.Port
		}
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Hostname,
			port,
			config.Database,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case configs.DATABASE_TYPE_POSTGRES:
		port = 5432
		if config.Port > 1 {
			port = config.Port
		}
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			config.Hostname,
			config.Username,
			config.Password,
			config.Database,
			port,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case configs.DATABASE_TYPE_SQLITE:
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		break
	case configs.DATABASE_TYPE_SQLSERVER:
		port = 1433
		if config.Port > 1 {
			port = config.Port
		}
		dsn := fmt.Sprintf(
			"sqlserver://%s:%s@%s:%d?database=%s",
			config.Username,
			config.Password,
			config.Hostname,
			port,
			config.Database,
		)
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		break
	default:
		log.Errorf("Type \"%s\" is not supported. Please verify.", config.Type)
		return nil
	}

	if err != nil {
		if config.PanicInError {
			panic(err)
		}
		log.Errorf("Database Initialization Error: %s", err)
	}

	log.Debugf("Database \"%s\" connected.", config.Type)
	return db
}
