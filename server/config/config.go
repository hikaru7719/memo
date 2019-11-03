package config

import "github.com/spf13/viper"

// Config is global configuration object
var Config = NewConfig()

// NewConfig returns Configulation Object
func NewConfig() *viper.Viper {
	v := viper.New()
	v.SetDefault("db.host", "localhost")
	v.SetDefault("db.port", "5432")
	v.SetDefault("db.name", "postgres")
	v.SetDefault("db.user", "postgres")
	v.SetDefault("db.password", "mypassword")
	v.SetDefault("db.sslmode", "disable")
	return v
}
