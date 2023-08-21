package models

import "time"

type Score struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	AssignmentTitle string `gorm:"not null" json:"assignmentTitle" form:"assignmentTitle" valid:"required~Assignment title is required"`
	Description     string `gorm:"not null" json:"description" form:"description" valid:"required~Description is required"`
	StudentID       uint   `json:"student"`
	Nilai           int
	CreatedAt       time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"-"`
}
