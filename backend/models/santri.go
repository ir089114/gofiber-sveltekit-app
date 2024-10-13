package models

import "gorm.io/gorm"

type Santri struct {
	gorm.Model
	Name      string `json:"name"`
	Age       int    `json:"age"`
	ClassName string `json:"className"`
}
