package global

import (
	"gorm.io/gorm"
	"visualization_chart/config"
)

type DbConfig struct {
	DataBase	*gorm.DB
	Tables		string
}

var (
	Config  *config.Server
	DB		*DbConfig

)

func init() {
	Config = &config.Server{}
	DB = &DbConfig{}
}
