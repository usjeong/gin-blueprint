package conf

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// CaseOne default app dependency objects
type CaseOne struct {
	Env      *viper.Viper
	DBReader *sqlx.DB
	DBWriter *sqlx.DB
}

// NewConfig set environment values
func NewConfig(env string) *viper.Viper {
	config := viper.New()
	pwd, _ := os.Getwd()

	config.SetConfigName("Set app Name!")
	config.AddConfigPath(pwd)
	config.SetConfigType("yaml")

	config.SetDefault("SentryKey", "")
	config.SetDefault("Store", "")
	config.SetDefault("StorePassword", "")
	config.SetDefault("Listen", ":8080")

	switch env {
	case "production":
		config.SetDefault("DBWriter", "")
		config.SetDefault("DBReader", "")
	case "develop":
		config.SetDefault("DBWriter", "")
		config.SetDefault("DBReader", "")
	default:
		config.SetDefault("DBWriter", "")
		config.SetDefault("DBReader", "")
	}

	config.ReadInConfig()
	return config
}

// ConnectDB create database connection
func ConnectDB(host string, env *viper.Viper) *sqlx.DB {
	driver := "mysql"
	switch host {
	case "Redshift":
		driver = "postgres"
	}

	sess, err := sqlx.Open(driver, env.GetString(host))
	sess.SetMaxIdleConns(20)
	sess.SetMaxOpenConns(20)

	if err != nil {
		sess.Close()
		log.Panic(err)
	}
	log.Printf("connect Database: %s\n", host)

	return sess
}

//NewCaseOne create default configuration object
func NewCaseOne(mode string) *CaseOne {
	if mode == "" {
		mode = os.Getenv("APP_ENV")
	}
	env := NewConfig(mode)
	caseOne := &CaseOne{
		Env:      env,
		DBWriter: ConnectDB("DBWriter", env),
		DBReader: ConnectDB("DBReader", env),
	}
	log.Println("configure case: One")

	return caseOne
}
