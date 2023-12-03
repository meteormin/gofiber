package monitor

import (
	"github.com/miniyus/gofiber/database"
	myReflect "github.com/miniyus/gofiber/internal/reflect"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type databaseInfo struct {
	Connections []connectionInfo
	Gorm        gormInfo
}

type connectionInfo struct {
	Connection        string
	Driver            string
	Host              string
	Dbname            string
	Port              string
	TimeZone          string
	SSlMode           bool
	AutoMigrate       []string
	Logger            gormLogger.Config
	MaxIdleConnection int
	MaxOpenConnection int
	MaxLifeTime       time.Duration
}

type gormInfo struct {
	SkipDefaultTransaction                   bool
	NamingStrategy                           schema.Namer
	FullSaveAssociations                     bool
	Logger                                   gormLogger.Interface
	DryRun                                   bool
	PrepareStmt                              bool
	DisableAutomaticPing                     bool
	DisableForeignKeyConstraintWhenMigrating bool
	// IgnoreRelationshipsWhenMigrating
	IgnoreRelationshipsWhenMigrating bool
	// DisableNestedTransaction disable nested transaction
	DisableNestedTransaction bool
	// AllowGlobalUpdate allow global update
	AllowGlobalUpdate bool
	// QueryFields executes the SQL query with all fields of the table
	QueryFields bool
	// CreateBatchSize default create batch size
	CreateBatchSize int
	// TranslateError enabling error translation
	TranslateError bool
	// ClauseBuilders clause builder
	ClauseBuilders map[string]clause.ClauseBuilder
	// ConnPool db conn pool
	ConnPool gorm.ConnPool
	// Plugins registered plugins
	Plugins map[string]gorm.Plugin
}

func newDbInfo(cfg map[string]database.Config, db *gorm.DB) databaseInfo {
	dbInfo := databaseInfo{}
	conn := make([]connectionInfo, 0)
	for _, dbCfg := range cfg {
		migrates := make([]string, 0)
		for _, ent := range dbCfg.AutoMigrate {
			migrates = append(migrates, myReflect.GetType(ent))
		}

		conn = append(conn, connectionInfo{
			Connection:        dbCfg.Name,
			Driver:            dbCfg.Driver,
			Host:              dbCfg.Host,
			Dbname:            dbCfg.Dbname,
			Port:              dbCfg.Port,
			TimeZone:          dbCfg.TimeZone,
			SSlMode:           dbCfg.SSLMode,
			AutoMigrate:       migrates,
			Logger:            dbCfg.Logger,
			MaxIdleConnection: dbCfg.MaxIdleConn,
			MaxOpenConnection: dbCfg.MaxOpenConn,
			MaxLifeTime:       dbCfg.MaxLifeTime,
		})
	}

	gormInfoVar := gormInfo{}
	if db != nil && db.Config != nil {
		gormCfg := db.Config
		gormInfoVar = gormInfo{
			SkipDefaultTransaction:                   gormCfg.SkipDefaultTransaction,
			NamingStrategy:                           gormCfg.NamingStrategy,
			FullSaveAssociations:                     gormCfg.FullSaveAssociations,
			Logger:                                   gormCfg.Logger,
			DryRun:                                   gormCfg.DryRun,
			PrepareStmt:                              gormCfg.PrepareStmt,
			DisableAutomaticPing:                     gormCfg.DisableAutomaticPing,
			DisableForeignKeyConstraintWhenMigrating: gormCfg.DisableForeignKeyConstraintWhenMigrating,
			IgnoreRelationshipsWhenMigrating:         gormCfg.IgnoreRelationshipsWhenMigrating,
			DisableNestedTransaction:                 gormCfg.DisableNestedTransaction,
			AllowGlobalUpdate:                        gormCfg.AllowGlobalUpdate,
			QueryFields:                              gormCfg.QueryFields,
			CreateBatchSize:                          gormCfg.CreateBatchSize,
			TranslateError:                           gormCfg.TranslateError,
			ClauseBuilders:                           gormCfg.ClauseBuilders,
			ConnPool:                                 gormCfg.ConnPool,
			Plugins:                                  gormCfg.Plugins,
		}
	}

	dbInfo.Connections = conn
	dbInfo.Gorm = gormInfoVar

	return dbInfo
}
