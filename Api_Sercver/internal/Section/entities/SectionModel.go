package entities

import (
	"github.com/DavG20/SchoolPortal/Api_Sercver/internal/Department/entities"
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Name       string                   `json:"name" gorm:"not null"`
	DepName    string                   `json:"dep_name"`
	Department entities.DepartmentModel `gorm:"foreignKey:DepName"`
}
