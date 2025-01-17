package entities

import (
	"github.com/mshyshov/PassAxion/pkg/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"size:255;not null;uniqueIndex"`
	PasswordHash string `gorm:"size:255"`
	IsAdmin      bool
}

func GetAllUsers(u *[]User) (err error) {
	if err := config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(u *User, id uint) (err error) {
	if err := config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(u *User) (err error) {
	if err := config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutUser(u *User, id string) (err error) {
	if err := config.DB.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(u *User, id string) (err error) {
	if err := config.DB.Where("id = ?", id).Delete(u).Error; err != nil {
		return err
	}
	return nil
}
