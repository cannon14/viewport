package services

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"test/app/models"
)

func GetAllWizProjects(limit int, offset int) ([]models.WizProject, error) {

	var wizProject []models.WizProject

	tx := Db.Preload(clause.Associations).
		Find(&wizProject).
		Limit(limit).
		Offset(offset)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		return wizProject, tx.Error
	}

	return wizProject, nil
}

func GetAllWizProjectsAsPayload(limit int, offset int) models.Payload[models.WizProject] {

	var payload models.Payload[models.WizProject]

	wizProject, err := GetAllWizProjects(limit, offset)

	if err != nil {
		payload.Error = err.Error()
	}
	var count int64
	Db.Table("viewport.wiz_projects").Count(&count)

	payload.RecordsFiltered = int64(len(wizProject))
	payload.RecordsTotal = count
	payload.Data = wizProject

	return payload
}

func GetWizProjectById(id uint16) models.WizProject {
	var wizProject models.WizProject

	tx := Db.First(&wizProject, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return wizProject

}

func CreateWizProject(wizProject models.WizProject) models.Payload[models.WizProject] {

	var payload models.Payload[models.WizProject]

	tx := Db.Select("*").Omit("id").Create(&wizProject)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.WizProject{wizProject}

	return payload
}

func UpdateWizProject(id uint16, wizProject models.WizProject) models.Payload[models.WizProject] {
	tx := Db.Model(&wizProject).Select("*").
		Omit("id", "created_at").Where("id = ?", id).
		Updates(wizProject)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	var payload models.Payload[models.WizProject]
	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.WizProject{wizProject}

	return payload
}

func DeleteWizProject(id uint16) {
	tx := Db.Delete(&models.WizProject{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}
