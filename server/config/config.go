package config

import "github.com/spf13/viper"

// Config is global configuration object
var Config = NewConfig()

// NewConfig returns Configulation Object
func NewConfig() *viper.Viper {
	v := viper.New()
	// Default database host is "db" because of using Docker-Compose
	v.SetDefault("db.host", "localhost")
	// Default database port is "5432" because of using PostgresSQL
	v.SetDefault("db.port", "5432")
	v.SetDefault("db.name", "postgres")
	v.SetDefault("db.user", "postgres")
	v.SetDefault("db.password", "mypassword")
	v.SetDefault("db.sslmode", "disable")
	return v
}
