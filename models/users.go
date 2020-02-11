package models

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"not null;unique_index"`
}