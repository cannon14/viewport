package models

import "time"

type User struct {
	Id        uint16    `gorm:"primaryKey;auto_increment;not_null" json:"id"`
	Teams     []*Team   `gorm:"many2many:user_team_maps;constraint:OnDelete:CASCADE" json:"teams"`
	Name      string    `gorm:"type:varchar(64);unique;not null" json:"name"`
	Title     string    `gorm:"type:varchar(65);not null" json:"title"`
	Email     string    `gorm:"type:varchar(255);unique not null" json:"email"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updatedAt"`
	Password  string    `gorm:"type:varchar(71);" json:"password"`
	IsAdmin   bool      `gorm:"type:boolean;default:false" json:"isAdmin"`
	LastLogin time.Time `gorm:"type:timestamp;default:null" json:"lastLogin"`
}
