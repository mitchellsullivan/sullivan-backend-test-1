package users

import (
	"errors"
	"github.com/mitchellsullivan/sullivan-backend-test-1/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `gorm:"column:id;primary_key"`
	Username     string `gorm:"column:username;unique_index"`
	Email        string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;not null"`
}

func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password is empty")
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *User) checkPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

// user-related db stuff

func AutoMigrate() {
	err := common.GetDb().AutoMigrate(&User{})
	if err != nil {
		return
	}
}

func FindOne(condition interface{}) (User, error) {
	db := common.GetDb()
	var model User
	err := db.Where(condition).First(&model).Error
	return model, err
}

func Save(data interface{}) error {
	db := common.GetDb()
	err := db.Save(data).Error
	return err
}
