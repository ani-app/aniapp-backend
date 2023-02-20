package models

type TreatmentType struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type VeterinaryTreatmentType struct {
	ID              uint `gorm:"primaryKey"`
	VeterinaryID    uint `gorm:"not null"`
	TreatmentTypeID uint `gorm:"not null"`
}
