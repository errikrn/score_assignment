package models

import "time"

type Scores struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	AssignmentTitle string    `gorm:"not null" json:"assignment_title" form:"assignment_title" valid:"required~Assignment title is required"`
	Description     string    `gorm:"not null" json:"description" form:"description" valid:"required~Description is required"`
	StudentID       uint      `gorm:"not null" json:"studentID" form:"studentID"`
	Score           int       `gorm:"not null" json:"score" form:"score" valid:"required~Score is required"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"-"`
}
