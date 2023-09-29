package services

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"test/app/models"
	"test/app/utils"
	"time"
)

func Login(email string, password string) (models.User, error) {

	user, err := GetUserByEmail(email)
	authenticated, err := utils.CheckPasswordHash(password, user.Password)

	if !authenticated || err != nil {
		logrus.Error(err)
		return user, errors.New("Unauthorized")
	}

	if err == nil {
		user.LastLogin = time.Now()
		_, err = UpdateUser(user.Id, user)
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func ResetPassword(userId uint16, password string) error {

	hash, _ := utils.HashPassword(password)
	user := models.User{
		Password: hash,
	}
	_, err := UpdateUser(userId, user)

	if err != nil {
		return err
	}

	return nil

}
