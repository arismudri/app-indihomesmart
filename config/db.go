package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

type ConnectionConfiguration struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbSslMode  string
}

func DbConfiguration() (string, string) {
	var (
		master  ConnectionConfiguration
		replica ConnectionConfiguration
	)

	master.DbName = viper.GetString("MASTER_DB_NAME")
	master.DbUser = viper.GetString("MASTER_DB_USER")
	master.DbPassword = viper.GetString("MASTER_DB_PASSWORD")
	master.DbHost = viper.GetString("MASTER_DB_HOST")
	master.DbPort = viper.GetString("MASTER_DB_PORT")
	master.DbSslMode = viper.GetString("MASTER_SSL_MODE")

	replica.DbName = viper.GetString("REPLICA_DB_NAME")
	replica.DbUser = viper.GetString("REPLICA_DB_USER")
	replica.DbPassword = viper.GetString("REPLICA_DB_PASSWORD")
	replica.DbHost = viper.GetString("REPLICA_DB_HOST")
	replica.DbPort = viper.GetString("REPLICA_DB_PORT")
	replica.DbSslMode = viper.GetString("REPLICA_SSL_MODE")

	masterDBDSN := dbSetConnection("mysql", master)
	replicaDBDSN := dbSetConnection("mysql", replica)

	return masterDBDSN, replicaDBDSN
}

func dbSetConnection(dbType string, attr ConnectionConfiguration) (connection string) {
	if dbType == "mysql" {
		connection = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			attr.DbUser, attr.DbPassword, attr.DbHost, attr.DbPort, attr.DbName,
		)
	} else {
		connection = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			attr.DbHost, attr.DbUser, attr.DbPassword, attr.DbName, attr.DbPort, attr.DbSslMode,
		)
	}
	return
}
