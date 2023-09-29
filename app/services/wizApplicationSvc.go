package services

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"test/app/models"
)

func GetAllWizApplicationsAsPayload(limit int, offset int) models.Payload[models.WizApplication] {

	var payload models.Payload[models.WizApplication]
	var wizApplication []models.WizApplication

	tx := Db.Preload(clause.Associations).
		Find(&wizApplication).
		Limit(limit).
		Offset(offset).
		Order("id")

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}
	var count int64
	Db.Table("viewport.wiz_applications").Count(&count)

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = count
	payload.Data = wizApplication

	return payload
}

func GetWizApplicationsBySubscriptionId(id uint16) []models.WizApplication {
	var wizApplication []models.WizApplication

	tx := Db.Where("subscription_id = ?", id).Find(&wizApplication)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return wizApplication

}

func GetWizApplicationByApplicationId(id uint16) models.WizApplication {
	var wizApplication models.WizApplication

	tx := Db.First(&wizApplication, "application_id = ?", id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return wizApplication
}

func GetWizApplicationById(id uint16) models.WizApplication {
	var wizApplication models.WizApplication

	tx := Db.First(&wizApplication, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return wizApplication
}

func CreateWizApplication(wizApplication models.WizApplication) models.Payload[models.WizApplication] {

	var payload models.Payload[models.WizApplication]

	tx := Db.Select("*").Create(&wizApplication)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.WizApplication{wizApplication}

	return payload
}

func UpdateWizApplication(id uint16, wizApplication models.WizApplication) models.Payload[models.WizApplication] {
	tx := Db.Model(&wizApplication).Select("*").
		Omit("id", "created_at").Where("id = ?", id).
		Updates(wizApplication)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	var payload models.Payload[models.WizApplication]
	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.WizApplication{wizApplication}

	return payload
}

func DeleteWizApplication(id uint16) {
	tx := Db.Delete(&models.WizApplication{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}
