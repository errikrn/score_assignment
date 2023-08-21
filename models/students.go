package models

import "time"

type Student struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	Age       int       `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
	Scores    []Score   `gorm:"foreignKey:StudentID" json:"scores"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}
