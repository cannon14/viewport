package services

import (
	"github.com/sirupsen/logrus"
	"test/app/models"
)

func GetAllAccounts(offset int, limit int) models.Payload[models.AwsAccount] {

	var accounts []models.AwsAccount

	tx := Db.
		Limit(limit).
		Offset(offset).
		Find(&accounts)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	var count int64
	Db.Table("viewport.aws_accounts").Count(&count)

	var payload models.Payload[models.AwsAccount]
	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = count
	payload.Data = accounts

	return payload
}

func GetAccountById(id uint) models.AwsAccount {
	var account models.AwsAccount

	tx := Db.First(&account, id).Error

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return account

}

func CreateAccount(account models.AwsAccount) models.Payload[models.AwsAccount] {

	tx := Db.Select("*").Create(&account)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	var payload models.Payload[models.AwsAccount]
	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.AwsAccount{account}

	return payload
}

func UpdateAccount(id uint, account models.AwsAccount) models.Payload[models.AwsAccount] {

	tx := Db.Model(&account).Select("*").Omit("id", "created_at").Where("id = ?", id).Updates(account)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	var payload models.Payload[models.AwsAccount]
	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.AwsAccount{account}

	return payload
}

func DeleteAccount(id int) {
	tx := Db.Delete(&models.AwsAccount{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}
