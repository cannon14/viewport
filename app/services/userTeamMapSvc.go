package services

import (
	"github.com/sirupsen/logrus"
	"test/app/models"
)

func CreateUserTeamMap(userId uint16, teamIds []uint16) {

	for _, id := range teamIds {
		team := models.UserTeamMap{
			UserId: userId,
			TeamId: id,
		}
		tx := Db.Create(&team)

		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}
}

func UpdateUserTeamMap(userId uint16, teamIds []uint16) {

	tx := Db.Where("user_id = ?", userId).
		Delete(models.UserTeamMap{})

	for _, id := range teamIds {
		team := models.UserTeamMap{
			UserId: userId,
			TeamId: id,
		}
		tx = Db.Create(&team)

		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}
}
