package services

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"test/app/models"
)

func GetAllWizSubscriptions(limit int, offset int) ([]models.WizSubscription, error) {

	var wizSubscription []models.WizSubscription

	tx := Db.Preload(clause.Associations).
		Find(&wizSubscription).
		Limit(limit).
		Offset(offset)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		return wizSubscription, tx.Error
	}

	return wizSubscription, nil
}

func GetAllWizSubscriptionAsPayload(limit int, offset int) models.Payload[models.WizSubscription] {

	var payload models.Payload[models.WizSubscription]

	wizSubscription, err := GetAllWizSubscriptions(limit, offset)

	if err != nil {
		payload.Error = err.Error()
	}
	var count int64
	Db.Table("viewport.wiz_subscriptions").Count(&count)

	payload.RecordsFiltered = int64(len(wizSubscription))
	payload.RecordsTotal = count
	payload.Data = wizSubscription

	return payload
}

func GetWizSubscriptionByProjectId(id uint16) []models.WizSubscription {
	var wizSubscription []models.WizSubscription

	tx := Db.Where("project_id = ?", id).Find(&wizSubscription)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return wizSubscription

}

func GetWizSubscriptionById(id uint16) models.WizSubscription {
	var wizSubscription models.WizSubscription

	tx := Db.First(&wizSubscription, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return wizSubscription

}

func CreateWizSubscription(wizSubscription models.WizSubscription) models.Payload[models.WizSubscription] {

	var payload models.Payload[models.WizSubscription]

	tx := Db.Select("*").Omit("id").Create(&wizSubscription)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.WizSubscription{wizSubscription}

	return payload
}

func UpdateWizSubscription(id uint16, wizSubscription models.WizSubscription) models.Payload[models.WizSubscription] {
	tx := Db.Model(&wizSubscription).Select("*").
		Omit("id", "created_at").Where("id = ?", id).
		Updates(wizSubscription)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	var payload models.Payload[models.WizSubscription]
	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.WizSubscription{wizSubscription}

	return payload
}

func DeleteWizSubscription(id uint16) {
	tx := Db.Delete(&models.WizSubscription{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}
