package models

type Breed struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name" gorm:"not null"`
	BreedTypeID uint      `json:"breed_type_id" gorm:"not null"`
	BreedType   BreedType `json:"breed_type" gorm:"embedded"`
}
