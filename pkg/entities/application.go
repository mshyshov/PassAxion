package entities

import (
	"github.com/mshyshov/PassAxion/pkg/config"
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;uniqueIndex"`
	Description string `gorm:"type:text"`
	URI         string `gorm:"type:text"`
}

func GetAllApplications(a *[]Application) (err error) {
	if err := config.DB.Find(a).Error; err != nil {
		return err
	}
	return nil
}

func GetApplication(a *Application, id uint) (err error) {
	if err := config.DB.Where("id = ?", id).First(a).Error; err != nil {
		return err
	}
	return nil
}

func CreateApplication(a *Application) (err error) {
	if err := config.DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}

func PutApplication(a *Application, id string) (err error) {
	if err := config.DB.Save(a).Error; err != nil {
		return err
	}
	return nil
}

func DeleteApplication(a *Application, id string) (err error) {
	if err := config.DB.Where("id = ?", id).Delete(a).Error; err != nil {
		return err
	}
	return nil
}
