package dal

import (
	"time"

	"github.com/ink-paint/ink/config"
	inkLog "github.com/ink-paint/ink/log"
	"github.com/ink-paint/ink/util/xerr"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
)

var (
	DB *gorm.DB
)

func NewGormDB(conf *config.Config, gormLogger logger.Interface) *gorm.DB {
	var err error

	DB, err = initDB(conf, gormLogger)
	if err != nil {
		inkLog.Fatal("connect to postgresql error", zap.Error(err))
	}
	inkLog.Info("connect database success")
	sqlDB, err := DB.DB()
	if err != nil {
		inkLog.Fatal("get database connection error")
	}
	sqlDB.SetMaxIdleConns(200)
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	// SetDefault(DB)
	// dbMigrate()
	return DB
}

func initDB(conf *config.Config, gormLogger logger.Interface) (*gorm.DB, error) {
	postgresConfig := conf.PostgreSQL
	if postgresConfig == nil {
		return nil, xerr.WithMsg(nil, "nil progresql config")
	}
	dsn := postgresConfig.Dsn

	inkLog.Info("try to open postgresql db", zap.String("dsn", `Use dsn in config`))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                   gormLogger,
		PrepareStmt:              true,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})
	return db, err
}
