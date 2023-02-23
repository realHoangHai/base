package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	C = new(Config)
)

func unmarshal(rawVal interface{}) bool {
	if err := viper.Unmarshal(rawVal); err != nil {
		fmt.Printf("config file loaded fail: %v\n", err)
		return false
	}
	return true
}

func Load(fpath string) bool {
	_, err := os.Stat(fpath)
	if err != nil {
		return false
	}

	viper.SetConfigFile(fpath)
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file %s read fail: %v\n", err)
		return false
	}

	if !unmarshal(&C) {
		return false
	}
	fmt.Printf("config %s load ok!\n", fpath)
	return true
}

type Config struct {
	AppConfig AppConfig      `mapstructure:"app"`
	DBConfig  PostgresConfig `mapstructure:"postgres"`
}

type AppConfig struct {
	Port            int    `mapstructure:"port"`
	Mode            string `mapstructure:"mode"`
	ShutdownTimeout int64  `mapstructure:"shutdown_timeout"`
}

type PostgresConfig struct {
	Driver    string `mapstructure:"driver"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Database  string `mapstructure:"database"`
	SSLMode   string `mapstructure:"ssl_mode"`
	Migration bool   `mapstructure:"migration"`
}

func (c PostgresConfig) DSN() string {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl != "" {
		return dbUrl
	}
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		c.Driver, c.User, c.Password, c.Host, c.Port, c.Database, c.SSLMode)
}
