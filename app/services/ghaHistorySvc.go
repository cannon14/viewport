package services

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"test/app/models"
	"test/app/utils"
	"time"
)

func GetGhaStats(applicationName string, date string) models.Payload[models.GhaStats] {

	var payload models.Payload[models.GhaStats]

	var result []models.GhaStats
	targetDate := time.Now()

	if len(date) > 0 {
		parsedDate, err := time.Parse("2023-01-01", date)

		if err != nil {
			payload.Error = err.Error()
			return payload

		}

		targetDate = parsedDate
	}

	tx := Db.Table("viewport.gha_histories").Select("environment", "event", "application_name", "count(*)")

	if len(applicationName) > 0 {
		tx.Where("application_name = ?", applicationName)
	}
	tx.Where("timestamp >= ?", utils.BeginningOfMonth(targetDate)).
		Where("timestamp <= ?", utils.EndOfMonth(targetDate)).
		Group("environment").
		Group("event").
		Group("application_name").
		Order("application_name").
		Order("environment").
		Find(&result)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
		return payload
	}

	payload.Data = result

	return payload

}

func GetAllGhaHistory(offset int, limit int, filters map[string]string) ([]models.GhaHistory, error) {

	var history []models.GhaHistory

	tx := Db.Preload(clause.Associations)

	from, ok := filters["from"]
	if ok {
		tx.Where("timestamp >= ?", from)
		delete(filters, "from")
	}

	to, ok := filters["to"]
	if ok {
		tx.Where("timestamp <= ?", to)
		delete(filters, "to")
	}

	tx.Where(filters)

	if limit > 0 {
		tx.Limit(limit)
	}

	tx.Offset(offset).
		Order("timestamp desc").
		Find(&history)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		return history, tx.Error
	}

	return history, nil
}

func GetAllGhaHistoryAsPayload(offset int, limit int, filters map[string]string) models.Payload[models.GhaHistory] {

	var payload models.Payload[models.GhaHistory]

	history, err := GetAllGhaHistory(offset, limit, filters)

	if err != nil {
		payload.Error = err.Error()
	}
	var count int64
	Db.Table("viewport.gha_histories").Count(&count)

	payload.RecordsFiltered = count
	payload.RecordsTotal = count
	payload.Data = history

	return payload
}

func CreateGhaHistory(history models.GhaHistory) models.GhaHistory {

	tx := Db.Select("*").Create(&history)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return history
}

func DeleteGhaHistory(id int) {
	tx := Db.Delete(&models.GhaHistory{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}
