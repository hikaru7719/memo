package repository

import (
	"fmt"

	"github.com/hikaru7719/memo/server/config"
	"github.com/hikaru7719/memo/server/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			config.Config.GetString("db.host"),
			config.Config.GetString("db.port"),
			config.Config.GetString("db.user"),
			config.Config.GetString("db.name"),
			config.Config.GetString("db.password"),
			config.Config.GetString("db.sslmode")))
	if err != nil {
		panic(err)
	}
	db = d
	db.AutoMigrate(&entity.Done{})
	db.AutoMigrate(&entity.Tag{})
}
