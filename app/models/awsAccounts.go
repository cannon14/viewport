package models

import (
	"time"
)

type AwsAccount struct {
	Id        uint      `gorm:"primaryKey;auto_increment;not_null" json:"id"`
	Name      string    `gorm:"type:varchar(64);unique;not null" json:"name"`
	Alias     string    `gorm:"type:varchar(64);unique" json:"alias"`
	AccountId string    `gorm:"type:varchar(64);unique;not null" json:"accountId"`
	Notes     string    `gorm:"type:text" json:"notes"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updatedAt"`
}
