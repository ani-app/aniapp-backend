package models

import (
	"gorm.io/gorm"
	"time"
)

type Treatment struct {
	gorm.Model
	PetID           uint       `gorm:"not null"`
	Pet             Pet        `gorm:"embedded"`
	VeterinaryID    uint       `gorm:"not null"`
	Veterinary      Veterinary `gorm:"embedded"`
	Date            time.Time  `gorm:"not null"`
	IsApproved      bool       `gorm:"default:false"`
	ApprovedAt      time.Time
	IsCompleted     bool `gorm:"default:false"`
	CompletedAt     time.Time
	TreatmentTypeID uint          `gorm:"not null"`
	TreatmentType   TreatmentType `gorm:"embedded"`
}
