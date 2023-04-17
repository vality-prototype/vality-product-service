package datastore

import (
	"log"

	"github.com/vality-prototype/vality-utility-service/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(env configs.Env) *gorm.DB {
	connStr := env.DBConnectionStr
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
