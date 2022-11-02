package database

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	DB  *gorm.DB
	err error
)

// DbConnection create database connection
func DbConnection(masterDSN, replicaDSN string) error {
	var db = DB

	logMode := viper.GetBool("DB_LOG_MODE")
	debug := viper.GetBool("DEBUG")

	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       masterDSN,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})
	if !debug {
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				mysql.Open(replicaDSN),
			},
			Policy: dbresolver.RandomPolicy{},
		}))
	}
	if err != nil {
		log.Fatalf("Db connection error")
		return err
	}
	DB = db
	return nil
}

// GetDB connection
func GetDB() *gorm.DB {
	return DB
}
