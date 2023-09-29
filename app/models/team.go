package models

import "time"

type Team struct {
	Id        uint      `gorm:"primaryKey;auto_increment;not_null" json:"id"`
	Name      string    `gorm:"type:varchar(64);unique;not null" json:"name"`
	Users     []*User   `gorm:"many2many:user_team_maps;" json:"users"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updatedAt"`
}
