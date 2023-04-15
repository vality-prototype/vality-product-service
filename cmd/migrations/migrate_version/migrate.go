package migrate_version

import (
	"log"

	"github.com/vality-prototype/vality-product-service/pkg/entities/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(connection string) {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	models.Migrate(db)
}
