package services

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"test/app/models"
	"test/app/utils"
	"time"
)

func GetAllApplications(limit int, offset int) ([]models.Application, error) {

	var applications []models.Application

	tx := Db.
		Joins("Team").
		Joins("RequestedByUser").
		Joins("ApprovedByUser").
		Find(&applications).
		Limit(limit).
		Offset(offset)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		return applications, tx.Error
	}

	return applications, nil
}

func GetAllApplicationsAsPayload(limit int, offset int) models.Payload[models.Application] {

	var payload models.Payload[models.Application]

	applications, err := GetAllApplications(limit, offset)

	if err != nil {
		payload.Error = err.Error()
	}
	var count int64
	Db.Table("viewport.applications").Count(&count)

	payload.RecordsFiltered = count
	payload.RecordsTotal = count
	payload.Data = applications

	return payload
}

func GetApplicationById(id uint16) models.Application {
	var application models.Application

	tx := Db.First(&application, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return application
}

func CreateApplication(application models.Application) models.Payload[models.Application] {

	var payload models.Payload[models.Application]

	tx := Db.Select("*").Create(&application)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.Application{application}

	return payload
}

func UpdateApplication(id uint16, application models.Application) models.Payload[models.Application] {

	var payload models.Payload[models.Application]

	tx := Db.Model(&application).Select("*").Omit("id", "created_at").Where("id = ?", id).Updates(application)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.Application{application}

	return payload
}

func DeleteApplication(id uint16) {
	tx := Db.Delete(&models.Application{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}

func ApplicationStatus(slug string) (map[string]interface{}, error) {
	var application models.Application
	slug = strings.ToLower(slug)

	err := Db.Model(&application).
		Where("slug = ?", slug).
		First(&application).
		Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Not Found")
	} else if err != nil {
		logrus.Error(err)
	}

	today := time.Now()
	isExempt := application.ExemptTo == nil || (today.Before(*application.ExemptTo) && application.ApprovedBy != nil)

	payload := make(map[string]interface{})
	payload["isExempt"] = isExempt

	return payload, nil
}

func RecordAndCheckStatus(slug string, vulnerabilities string) (map[string]interface{}, error) {
	var application models.Application
	slug = strings.ToLower(slug)

	err := Db.Model(&application).
		Where("slug = ?", slug).
		First(&application).
		Error

	//If err is that app wasn't found
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New(slug + " application not found")
	} else if err != nil {
		logrus.Error(err)
	}

	vuls, err := utils.ExtractVulnerabilityCounts(vulnerabilities)

	if err != nil {
		//TODO do something with this error
		logrus.Error(err)
	} else {
		application.Critical = vuls["critical"]
		application.High = vuls["high"]
		application = UpdateApplication(application.Id, application).Data[0]
	}

	isExempt := utils.IsExempt(application, vuls)

	payload := make(map[string]interface{})
	payload["isExempt"] = isExempt

	return payload, nil
}
