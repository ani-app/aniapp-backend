package models

type CustomerAddress struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Line       string   `json:"line" gorm:"not null"`
	ProvinceID uint     `json:"province_id" gorm:"not null"`
	Province   Province `json:"province" gorm:"embedded"`
	CustomerID uint     `json:"customer_id" gorm:"not null"`
}
