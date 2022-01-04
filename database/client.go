package database

import (
	"log"

	"neopeople-service/model"

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
	Connector.Migrator().CreateTable(&model.Event{})
	Connector.Migrator().CreateTable(&model.Attendance{})
	Connector.Migrator().CreateTable(&model.EventOrder{})
	Connector.Migrator().CreateTable(&model.Session{})
	Connector.Migrator().CreateTable(&model.Counselor{})
	Connector.Migrator().CreateTable(&model.Pantient{})
	Connector.Migrator().CreateTable(&model.Faq{})
	Connector.Migrator().CreateTable(&model.Team{})

	Connector.Migrator().CreateConstraint(&model.Event{}, "EventOrder")
	Connector.Migrator().CreateConstraint(&model.Event{}, "fk_event_order")

	Connector.Migrator().CreateConstraint(&model.Event{}, "Attendance")
	Connector.Migrator().CreateConstraint(&model.Event{}, "fk_attendance")

	Connector.Migrator().CreateConstraint(&model.Session{}, "Counselor")
	Connector.Migrator().CreateConstraint(&model.Session{}, "fk_counselor")

	Connector.Migrator().CreateConstraint(&model.Session{}, "Pantient")
	Connector.Migrator().CreateConstraint(&model.Session{}, "fk_pantient")

}
