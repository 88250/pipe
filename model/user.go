package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `sql:"type:VARCHAR(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"`
}
