package entity

import "time"

type Question struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint `gorm:"not null"`
	PublishDate      time.Time
	Status           string    `gorm:"size:255;not null"`
	Title            string    `gorm:"size:255;not null"`
	Body             string    `gorm:"size:10000;not null"`
	TimeLimit        int64     `gorm:"not null"`
	MemoryLimit      int64     `gorm:"not null"`
	InputURL         string    `gorm:"size:255;not null"`
	Deadline         time.Time `gorm:"not null"`
	OutputURL        string    `gorm:"size:255;not null"`
	SubmissionsCount int64     `gorm:"not null"`

	// Foreign key
	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}