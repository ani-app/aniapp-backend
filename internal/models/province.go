package models

type Province struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	CityID uint   `json:"city_id,omitempty" gorm:"not null"`
	City   City   `json:"city,omitempty" gorm:"embedded"`
	Name   string `json:"name" gorm:"not null"`
}
