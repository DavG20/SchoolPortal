package entities

import "gorm.io/gorm"

type DepartmentModel struct {
	gorm.Model
	SchoolName string `json:"school_name"`
	DepName    string `json:"dep_name" gorm:"not null"`

}
