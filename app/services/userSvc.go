package services

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"test/app/models"
)

func GetAllUsers(limit int, offset int) models.Payload[models.User] {

	var payload models.Payload[models.User]
	var users []models.User

	tx := Db.Preload(clause.Associations).
		Find(&users).
		Limit(limit).
		Offset(offset).
		Order("name")

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}
	var count int64
	Db.Table("viewport.users").Count(&count)

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = count
	payload.Data = users

	return payload
}

func GetUserById(id uint16) models.User {
	var user models.User

	tx := Db.First(&user, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return user

}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	if result := Db.Select("*").
		Where("email = ?", email).
		First(&user); result.Error != nil {

		logrus.Error(result.Error)
		return user, result.Error
	}

	return user, nil

}

func CreateUser(user models.User, teamIds []uint16) models.Payload[models.User] {

	var payload models.Payload[models.User]

	tx := Db.Select("*").Create(&user)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}

	CreateUserTeamMap(user.Id, teamIds)

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.User{user}

	return payload
}

func UpdateUser(id uint16, user models.User) (models.User, error) {
	tx := Db.Model(&user).Select("*").
		Omit("id", "created_at").Where("id = ?", id).
		Updates(user)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		return user, tx.Error
	}

	return user, nil
}

func UpdateUserForPayload(id uint16, user models.User) (models.Payload[models.User], error) {
	tx := Db.Model(&user).Select("*").
		Omit("id", "created_at").Where("id = ?", id).
		Updates(user)

	var payload models.Payload[models.User]

	if tx.Error != nil {
		logrus.Error(tx.Error)
		return payload, tx.Error
	}

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.User{user}

	return payload, nil
}

func DeleteUser(id uint16) {
	tx := Db.Delete(&models.User{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}
