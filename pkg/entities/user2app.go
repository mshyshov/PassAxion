package entities

import (
	"github.com/mshyshov/PassAxion/pkg/config"
	"gorm.io/gorm"
)

type User2App struct {
	gorm.Model
	UserID        uint
	ApplicationID uint
}

func GetAllUser2Apps(u *[]User2App) (err error) {
	if err := config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUser2App(u *User2App, id uint) (err error) {
	if err := config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser2App(u *User2App) (err error) {
	if err := config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutUser2App(u *User2App, id string) (err error) {
	if err := config.DB.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser2App(u *User2App, id string) (err error) {
	if err := config.DB.Where("id = ?", id).Delete(u).Error; err != nil {
		return err
	}
	return nil
}
