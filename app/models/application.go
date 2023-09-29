package models

import (
	"time"
)

type Application struct {
	Id              uint16     `gorm:"primaryKey;auto_increment;not_null" json:"id"`
	Name            string     `gorm:"type:varchar(64);unique;not null" json:"name"`
	Slug            string     `gorm:"type:varchar(64);unique" json:"slug"`
	Repository      string     `gorm:"type:varchar(255);unique;not null" json:"repository"`
	TeamId          uint16     `gorm:"default:null" json:"teamId"`
	Team            Team       `gorm:"foreignKey:TeamId" json:"team"`
	Critical        uint8      `gorm:"default:0" json:"critical"`
	High            uint8      `gorm:"default:0" json:"high"`
	ExemptTo        *time.Time `gorm:"type:timestamp;default:null" json:"exemptTo"`
	ExemptReason    string     `gorm:"type:text" json:"exemptReason"`
	RequestedBy     *uint16    `gorm:"default:null" json:"requestedBy"`
	RequestedByUser *User      `gorm:"foreignKey:RequestedBy" json:"requestedByUser"`
	ApprovedBy      *uint16    `gorm:"default:null" json:"approvedBy"`
	ApprovedByUser  *User      `gorm:"foreignKey:ApprovedBy" json:"approvedByUser"`
	//WizApplication  *WizApplication `json:"wizApplication"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updatedAt"`
}
