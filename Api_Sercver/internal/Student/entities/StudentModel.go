package entities

import (
	"github.com/DavG20/SchoolPortal/Api_Sercver/internal/Department/entities"
	"gorm.io/gorm"
)

type StudentModel struct {
	gorm.Model
	StrudetId       string                   `json:"student_id" gorm:"primary_key"`
	FName           string                   `json:"f_name" gorm:"not null"`
	LName           string                   `json:"l_name" gorm:"not null"`
	DepName         string                   `json:"dep_name"`
	DepartmentModel entities.DepartmentModel `gorm:"foreignKey:DepName"`
	SectionId       string                   `json:"section_id"`
	
}
