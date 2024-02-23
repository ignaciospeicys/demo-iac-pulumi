package model

import (
	"time"
)

type Resource struct {
	ResourceID            uint            `gorm:"primaryKey"`
	ResourceName          string          `gorm:"size:255;not null"`
	QualifiedResourceName string          `gorm:"size:255;not null"`
	ResourceType          string          `gorm:"size:50;not null"`
	StackName             string          `gorm:"size:255;not null"`
	CreatedAt             time.Time       `gorm:"default:now()"`
	Status                string          `gorm:"size:50;not null"`
	LastUpdated           time.Time       `gorm:"default:now()"`
	Configurations        []Configuration `gorm:"foreignKey:ResourceID"`
}

type Configuration struct {
	ConfigID    uint      `gorm:"primaryKey"`
	ResourceID  uint      `gorm:"index"`
	ConfigKey   string    `gorm:"size:255;not null"`
	ConfigValue string    `gorm:"type:text;not null"`
	LastUpdated time.Time `gorm:"default:now()"`
}
