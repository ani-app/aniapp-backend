package models

import "gorm.io/gorm"

type Veterinary struct {
	gorm.Model
	Phone           string `gorm:"not null"`
	Email           string `gorm:"not null"`
	Password        string `gorm:"not null"`
	Name            string `gorm:"not null"`
	ProfilePhotoUrl string
	Score           uint
	IsOpen          bool `gorm:"not null"`
}

type VeterinaryDTO struct {
	ID              uint   `json:"id"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Password        string `json:"password,omitempty"`
	Name            string `json:"name"`
	ProfilePhotoUrl string `json:"profile_photo_url"`
	Score           uint   `json:"score"`
	IsOpen          bool   `json:"is_open"`
}

func (v Veterinary) toVeterinaryDTO() VeterinaryDTO {
	return VeterinaryDTO{
		ID:              v.ID,
		Phone:           v.Phone,
		Email:           v.Email,
		Password:        v.Password,
		Name:            v.Name,
		ProfilePhotoUrl: v.ProfilePhotoUrl,
		Score:           v.Score,
		IsOpen:          v.IsOpen,
	}
}

func (vd VeterinaryDTO) toVeterinary() Veterinary {
	return Veterinary{
		Model:           gorm.Model{ID: vd.ID},
		Phone:           vd.Phone,
		Email:           vd.Email,
		Password:        vd.Password,
		Name:            vd.Name,
		ProfilePhotoUrl: vd.ProfilePhotoUrl,
		Score:           vd.Score,
		IsOpen:          vd.IsOpen,
	}
}

func ToVeterinaryDTOs(veterinaries []Veterinary) []VeterinaryDTO {
	veterinaryDTOs := make([]VeterinaryDTO, len(veterinaries))
	for i, v := range veterinaries {
		veterinaryDTOs[i] = v.toVeterinaryDTO()
	}
	return veterinaryDTOs
}
