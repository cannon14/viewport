package services

import (
	"github.com/sirupsen/logrus"
	"test/app/models"
)

func GetAllTeams(limit int, offset int) models.Payload[models.Team] {

	var payload models.Payload[models.Team]
	var teams []models.Team

	tx := Db.Find(&teams).
		Limit(limit).
		Offset(offset).
		Order("name")

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()
	}
	var count int64
	Db.Table("viewport.teams").Count(&count)

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = count
	payload.Data = teams

	return payload
}

func GetTeamById(id uint) models.Team {
	var team models.Team

	tx := Db.First(&team, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	return team

}

func CreateTeam(team models.Team) models.Payload[models.Team] {

	var payload models.Payload[models.Team]

	tx := Db.Select("*").Create(&team)

	if tx.Error != nil {
		logrus.Error(tx.Error)
		payload.Error = tx.Error.Error()

	}

	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.Team{team}

	return payload
}

func UpdateTeam(id uint, team models.Team) models.Payload[models.Team] {
	tx := Db.Model(&team).Select("*").Omit("id", "created_at").Where("id = ?", id).Updates(team)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}

	var payload models.Payload[models.Team]
	payload.RecordsFiltered = tx.RowsAffected
	payload.RecordsTotal = 1
	payload.Data = []models.Team{team}

	return payload
}

func DeleteTeam(id uint) {
	tx := Db.Delete(&models.Team{}, id)

	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
}
