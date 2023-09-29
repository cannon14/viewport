package models

import (
	"time"
)

type WizProject struct {
	Id        uint16    `gorm:"primaryKey;autoincrement;not null" json:"id"`
	Name      string    `gorm:"type:varchar(64);unique;not null" json:"name"`
	ProjectId string    `gorm:"type:varchar(64);unique;not null" json:"projectId"`
	Notes     string    `gorm:"type:text" json:"notes"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updatedAt"`
}

type WizSubscription struct {
	Id             uint16     `gorm:"primaryKey;autoincrement;not null" json:"id"`
	Name           string     `gorm:"type:varchar(64);unique;not null" json:"name"`
	SubscriptionId string     `gorm:"type:varchar(64);unique;not null" json:"subscriptionId"`
	ProjectId      uint16     `gorm:"not null" json:"projectId"`
	Project        WizProject `gorm:"references:Id" json:"project"`
	Environment    string     `gorm:"type:varchar(15);not null" json:"environment"`
	Notes          string     `gorm:"type:text" json:"notes"`
	CreatedAt      time.Time  `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt      time.Time  `gorm:"type:timestamp;default:now()" json:"updatedAt"`
}

type WizApplication struct {
	Id             uint16          `gorm:"primaryKey;autoincrement;not null" json:"id"`
	SubscriptionId uint16          `gorm:"not null" json:"subscriptionId"`
	Subscription   WizSubscription `gorm:"references:Id" json:"subscription"`
	ApplicationId  uint16          `gorm:"not null" json:"applicationId"`
	Application    Application     `gorm:"references:Id" json:"application"`
	Name           string          `gorm:"type:varchar(64);unique;not null" json:"name"`
	Critical       uint8           `gorm:"type:int;default:0" json:"critical"`
	High           uint8           `gorm:"type:int;default:0" json:"high"`
	Notes          string          `gorm:"type:text" json:"notes"`
	CreatedAt      time.Time       `gorm:"type:timestamp;default:now()" json:"createdAt"`
	UpdatedAt      time.Time       `gorm:"type:timestamp;default:now()" json:"updatedAt"`
}
