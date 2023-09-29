package models

import (
	"time"
)

type UserTeamMap struct {
	UserId    uint16    `gorm:"not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserId" json:"user"`
	TeamId    uint16    `gorm:"not null" json:"teamId"`
	Team      Team      `gorm:"foreignKey:TeamId" json:"team"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updatedAt"`
}
