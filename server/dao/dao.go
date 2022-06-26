package dao

import (
	"io"
	"log"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb(logWriter io.Writer) {
	var err error
	db, err = gorm.Open(mysql.Open(viper.GetString("dsn")), &gorm.Config{
		Logger: logger.New(
			log.New(logWriter, "[GORM] ", log.LstdFlags|log.Llongfile|log.Lmicroseconds),
			logger.Config{
				SlowThreshold:             100 * time.Millisecond, // Slow SQL threshold
				Colorful:                  viper.GetBool("enable_console_color"),
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				LogLevel:                  logger.Info, // show basically all sql gorm generated for DEBUG
			}),
	})
	if err != nil {
		panic(errors.Because(errors.New("initialization failed"), err, ""))
	}
	if sqlDb, err := db.DB(); err != nil {
		panic(errors.Because(errors.New("initialization failed"), err, ""))
	} else {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDb.SetMaxIdleConns(20)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDb.SetMaxOpenConns(100)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDb.SetConnMaxLifetime(time.Hour)
	}
}
