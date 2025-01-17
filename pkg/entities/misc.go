package entities

import (
	"log"

	"gorm.io/gorm"
)

func MigrateEntities(d *gorm.DB) error {
	log.Println("Starting automatic migrations")
	err := d.AutoMigrate(&Application{})
	if err != nil {
		return err
	}

	err = d.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	err = d.AutoMigrate(&User2App{})
	if err != nil {
		return err
	}

	log.Println("All models migrated successfully")
	return nil
}
