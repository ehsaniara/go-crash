package models

import "gorm.io/gorm"

type Customer struct {
	Model
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Title     string `json:"title" validate:"required"`
}

func GetCustomerById(id int) (*Customer, error) {
	var customer Customer
	err := db.Where("id = ?", id).First(&customer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &customer, nil
}
