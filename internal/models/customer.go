package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Phone           string `gorm:"not null"`
	Email           string `gorm:"not null"`
	Username        string `gorm:"not null"`
	Password        string
	IsGoogleAccount bool         `gorm:"not null"`
	CustomerTypeID  uint         `gorm:"not null"`
	CustomerType    CustomerType `gorm:"embedded"`
}

type CustomerDTO struct {
	ID           uint         `json:"id"`
	Phone        string       `json:"phone"`
	Email        string       `json:"email"`
	Password     string       `json:"password,omitempty"`
	Username     string       `json:"username"`
	CustomerType CustomerType `json:"customer_type"`
}

type CustomerDTOTable struct {
	ID           uint         `json:"id"`
	Email        string       `json:"email"`
	Username     string       `json:"username"`
	CustomerType CustomerType `json:"customer_type"`
}

func (c Customer) toCustomerDTO() CustomerDTO {
	return CustomerDTO{
		ID:           c.ID,
		Phone:        c.Phone,
		Email:        c.Email,
		Username:     c.Username,
		CustomerType: c.CustomerType,
	}
}

func (c Customer) toCustomerDTOTable() CustomerDTOTable {
	return CustomerDTOTable{
		ID:           c.ID,
		Email:        c.Email,
		Username:     c.Username,
		CustomerType: c.CustomerType,
	}
}

func (cd CustomerDTO) toCustomer() Customer {
	return Customer{
		Model:          gorm.Model{ID: cd.ID},
		Phone:          cd.Phone,
		Email:          cd.Email,
		Username:       cd.Username,
		Password:       cd.Password,
		CustomerTypeID: cd.CustomerType.ID,
		CustomerType:   cd.CustomerType,
	}
}

func ToCustomerDTOs(customers []Customer) []CustomerDTOTable {
	customerDTOs := make([]CustomerDTOTable, len(customers))
	for i, c := range customers {
		customerDTOs[i] = c.toCustomerDTOTable()
	}
	return customerDTOs
}
