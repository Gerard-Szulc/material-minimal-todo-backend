package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool `json:"completed"`
}

//type TodoResponse struct {
//	Id        int    `json:"id"`
//	Title     string `json:"title"`
//	Completed bool   `json:"completed"`
//}
