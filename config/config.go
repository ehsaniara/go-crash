package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	App      App
	DataBase DataBase
	Redis    Redis
}

type App struct {
	RunMode       string
	Version       string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	EndPointPort  int
	JwtSecret     string
	PageSize      int
	ObjectCashTtl int
}

type DataBase struct {
	Host     string
	Port     int
	DbName   string
	Username string
	Password string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var AppConfig = &Config{}

func Setup() {

	// Set the file name of the configurations file
	viper.SetConfigName("config/config")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("Error reading config file, %s\n", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic("Unable to unmarshal config")
	}

	fmt.Printf("App Version: %s\n", AppConfig.App.Version)
}
