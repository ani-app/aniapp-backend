package models

type VeterinaryAddress struct {
	ID           uint     `gorm:"primaryKey" json:"id"`
	Line         string   `json:"line" gorm:"not null"`
	ProvinceID   uint     `json:"province_id" gorm:"not null"`
	Province     Province `json:"province" gorm:"embedded"`
	VeterinaryID uint     `json:"veterinary_id" gorm:"not null"`
}
