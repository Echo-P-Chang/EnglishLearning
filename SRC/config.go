package main

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

//Config...
type Config struct {
	//MariaDB
	DBUser     string
	DBPasswd   string
	DBHost     string
	DBPort     int
	DBDatabase string
	/*
		//Log
		LogLevel         Log.Level
		LogPath          string
		LogConsoleOutput bool

		//FTP
		FTPUser         string
		FTPHost         string
		FTPPort         int
		FTPRemoteFolder string

		//Worker
		WorkerOffset      int
		WorkerExportTable string
		WorkerRegion      string
		WorderFilePrefix  string */
}

func readConfig() (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.SetConfigType("toml")
	viper.SetConfigName("env")
	viper.AddConfigPath(dir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//the validation of db configuration
	if !viper.IsSet("mssql.user") {
		return nil, errors.New("Config lost: mariadb.user")
	}
	if !viper.IsSet("mssql.passwd") {
		return nil, errors.New("Config lost: mariadb.passwd")
	}
	if !viper.IsSet("mssql.host") {
		return nil, errors.New("Config lost: mariadb.host")
	}
	if !viper.IsSet("mssql.port") {
		return nil, errors.New("Config lost: mariadb.port")
	}
	if !viper.IsSet("mssql.database") {
		return nil, errors.New("Config lost: mariadb.database")
	}

	c := Config{
		DBUser:     viper.GetString("mssql.user"),
		DBPasswd:   viper.GetString("mssql.passwd"),
		DBHost:     viper.GetString("mssql.host"),
		DBPort:     viper.GetInt("mssql.port"),
		DBDatabase: viper.GetString("mssql.database"),
	}

	return &c, nil
}

/*
func chgToLevel(lv string) (Log.Level, error) {

	r := Log.LevelDebug
	err := errors.New("the value of log level is not right")

	switch strings.ToLower(lv) {
	case "debug":
		r = Log.LevelDebug
		err = nil
	case "info":
		r = Log.LevelInfo
		err = nil
	case "warn", "warning":
		r = Log.LevelWarn
		err = nil
	case "error":
		r = Log.LevelError
		err = nil
	case "fatal":
		r = Log.LevelFatal
		err = nil
	case "panic":
		r = Log.LevelPanic
		err = nil
	}
	return r, err
} */
