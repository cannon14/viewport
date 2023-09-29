package models

import (
	"time"
)

type GhaHistory struct {
	Id              uint16    `gorm:"primaryKey;auto_increment;not_null" json:"id"`
	ApplicationName string    `gorm:"type:varchar(64);not_null" json:"applicationName"`
	Environment     string    `gorm:"type:varchar(64);not null" json:"environment"`
	Event           string    `gorm:"type:varchar(64);not null" json:"event"`
	Reference       string    `gorm:"type:varchar(64);default: null" json:"reference"`
	Status          string    `gorm:"type:varchar(64);not null" json:"status"`
	Actor           string    `gorm:"type:varchar(54);not null" json:"actor"`
	Timestamp       time.Time `gorm:"type:timestamp;default:now()" json:"timestamp"`
}
