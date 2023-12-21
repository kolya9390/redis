package config

import (
	"os"
	"time"
)

const (
	AppName = "APP_NAME"

	serverPort         = "SERVER_PORT"
	envShutdownTimeout = "SHUTDOWN_TIMEOUT"
	envAccessTTL       = "ACCESS_TTL"
	envRefreshTTL      = "REFRESH_TTL"
	envVerifyLinkTTL   = "VERIFY_LINK_TTL"
)

type AppConf struct {
	AppName     string   `yaml:"app_name"`
	Environment string   `yaml:"environment"`
	Domain      string   `yaml:"domain"`
	APIUrl      string   `yaml:"api_url"`
	Server      Server   `yaml:"server"`
	Token       Token    `yaml:"token"`
	DB          DB       `yaml:"db"`
	AuthorizationDADATA	AuthorizationDADATA
}

type DB struct {
	Net      string `yaml:"net"`
	Driver   string `yaml:"driver"`
	Name     string `yaml:"name"`
	User     string `json:"-" yaml:"user"`
	Password string `json:"-" yaml:"password"`
	Host     string `yaml:"host"`
	MaxConn  int    `yaml:"max_conn"`
	Port     string `yaml:"port"`
	Timeout  int    `yaml:"timeout"`
}

type AuthorizationDADATA struct {
	ApiKeyValue    string
	SecretKeyValue string
}

type Server struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type Token struct {
	AccessTTL     time.Duration `yaml:"access_ttl"`
	RefreshTTL    time.Duration `yaml:"refresh_ttl"`
	AccessSecret  string        `yaml:"access_secret"`
	RefreshSecret string        `yaml:"refresh_secret"`
}


func NewAppConf() AppConf {
	port := os.Getenv(serverPort)

	return AppConf{
		AppName: os.Getenv(AppName),
		Server: Server{
			Port: port,
		},
		DB: DB{
			Net:      os.Getenv("DB_NET"),
			Driver:   os.Getenv("DB_DRIVER"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}
