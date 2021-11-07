package dao

import (
	"time"

	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/fmt/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(viper.GetString("dsn")), &gorm.Config{})
	if err != nil {
		panic(errors.Because(errors.New("initialization failed"), err, ""))
	}
	if sqlDb, err := db.DB(); err != nil {
		panic(errors.Because(errors.New("initialization failed"), err, ""))
	} else {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDb.SetMaxIdleConns(10)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDb.SetMaxOpenConns(100)
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDb.SetConnMaxLifetime(time.Hour)
	}
}
