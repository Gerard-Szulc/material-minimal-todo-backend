package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Uuid     string
	Email    string
	Password string
	Active   bool   `gorm:"default:false"`
	Todos    []Todo `gorm:"one2many:user_todos;"`
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
}

type Validation struct {
	Value string
	Valid string
}

type Register struct {
	Username string
	Email    string
	Password string
}
