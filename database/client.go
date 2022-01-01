package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error

	Connector, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	log.Println("Connection to database succesfull!")
	return nil
}

func Migrate() {

}
