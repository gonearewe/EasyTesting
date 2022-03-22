package dao

import (
	"time"

	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb() {
	var err error
	db, err = gorm.Open(mysql.Open(viper.GetString("dsn")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // show basically all sql gorm generated for DEBUG
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
		sqlDb.SetMaxOpenConns(200)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDb.SetConnMaxLifetime(time.Hour)
	}
}
