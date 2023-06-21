package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DisplayName string  `gorm:"not null"`
	Email     string    `gorm:"not null;unique_index"`
	Password  string    `gorm:"not null"`
}

type Analytix struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Openness  float32   `gorm:"not null"`
	Conscientiousness  float32   `gorm:"not null"`
	Extraversion  float32   `gorm:"not null"`
	Agreeableness  float32   `gorm:"not null"`
	Neuroticism  float32   `gorm:"not null"`
}