package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName"`
	LastName string `json:"lastName"`
}
