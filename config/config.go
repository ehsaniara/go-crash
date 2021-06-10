package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type Config struct {
	App      App
	DataBase DataBase
	Redis    Redis
}

type App struct {
	LogMode       string
	GinRunMode    string //"debug","release","test"
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
	profile := os.Getenv("PROFILE")
	fmt.Printf("PROFILE: %s\n", profile)

	// run it as: ~# export PROFILE=dev && go run main.go
	if os.Getenv("PROFILE") != "" {
		var profileName = ""
		if strings.Contains(profile, "dev") {
			// dev configs
			profileName = "_dev"
		} else if strings.Contains(profile, "prod") {
			// prod configs
			profileName = "_prod"
		} else if strings.Contains(profile, "docker") {
			// docker configs
			profileName = "_docker"
		}
		fmt.Printf("App `%s` profile Selected\n", profileName)
		viper.SetConfigName("config/config" + profileName)

		//override them from env
		AppConfig.App.JwtSecret = os.Getenv("APP_JWT_SECRET")
		AppConfig.DataBase.Username = os.Getenv("DATABASE_USER")
		AppConfig.DataBase.Password = os.Getenv("DATABASE_PASSWORD")
		AppConfig.DataBase.DbName = os.Getenv("DATABASE_DB")
		AppConfig.Redis.Password = os.Getenv("REDIS_PASSWORD")

	} else {
		fmt.Println("App has Default profile")
		viper.SetConfigName("config/config")
	}

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic("Unable to unmarshal config")
	}

	fmt.Printf("App Version: %s\n", AppConfig.App.Version)
}
