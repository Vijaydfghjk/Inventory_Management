package dbrepository

import (
	"errors"
	"fmt"
	models "inventory_management/Models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User_interect interface {
	Register_new_user(a models.User) (error, models.User)
	Logging(id uint, Password string) (error, models.User)
}

type User_db struct {
	userdb *gorm.DB
}

func User_repo(userdb *gorm.DB) *User_db {

	return &User_db{userdb: userdb}
}

func (u *User_db) Register_new_user(a models.User) (error, models.User) {

	if err := u.userdb.Where("email = ?", a.Email).Take(&a).Error; err != nil {

		fmt.Errorf("user already exists with this email")
	}

	hashedPassword, err := HashPassword(a.Password)

	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err), a
	}

	a.Password = hashedPassword

	if err := u.userdb.Create(&a).Error; err != nil {

		fmt.Errorf("Failed to create the user %v", err.Error())
	}

	return nil, a
}

func (u *User_db) Logging(id uint, Password string) (error, models.User) {

	var temp models.User
	if err := u.userdb.First(&temp, "id=?", id).Error; err != nil {

		return errors.New("user not found"), temp
	}
	if err := bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(Password)); err != nil {
		return errors.New("invalid password"), temp
	}
	return nil, temp
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
