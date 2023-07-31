package database

import (
	"fmt"
	"github.com/miniyus/gofiber/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	"time"
)

type Config struct {
	Name        string
	Driver      string
	Host        string
	Dbname      string
	Username    string
	Password    string
	Port        string
	TimeZone    string
	SSLMode     bool
	AutoMigrate []interface{}
	Logger      logger.Config
	MaxIdleConn int
	MaxOpenConn int
	MaxLifeTime time.Duration
}

var connections = make(map[string]*gorm.DB)

func GetDB(name ...string) *gorm.DB {
	if len(name) == 0 {
		return connections["default"]
	}

	return connections[name[0]]
}

func switchDriver(cfg Config) gorm.Dialector {
	switch cfg.Driver {
	case "postgres", "pgsql":
		var sslMode string = "disable"
		if cfg.SSLMode {
			sslMode = "enable"
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			cfg.Host, cfg.Username, cfg.Password, cfg.Dbname, cfg.Port, sslMode, cfg.TimeZone,
		)
		return postgres.Open(dsn)
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname, url.QueryEscape(cfg.TimeZone),
		)
		return mysql.Open(dsn)
	case "sqlite", "":
		dsn := ":memory"
		if cfg.Dbname != "" {
			dsn = fmt.Sprintf("%s.db", cfg.Dbname)
		}

		return sqlite.Open(dsn)
	default:
		return sqlite.Open(":memory")
	}
}

// New
// gorm.DB 객체 생성 함수
func New(cfg Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		cfg.Logger,
	)

	driver := switchDriver(cfg)

	db, err := gorm.Open(driver, &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("Failed: Connect DB %v", err)
	}

	log.Println("Success: Connect DB")

	if cfg.AutoMigrate != nil && len(cfg.AutoMigrate) != 0 {
		migrations.Migrate(db, cfg.AutoMigrate...)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed: Connect sqlDB %v", err)
	}

	sqlDB.SetConnMaxLifetime(cfg.MaxLifeTime)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	connections[cfg.Name] = db

	return db
}
