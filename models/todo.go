package models

import "gorm.io/gorm"

type ToDo struct {
	gorm.Model
	Title string
	Note string
	Complete bool
}