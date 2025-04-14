package entity

import "time"

type Submission struct {
	ID         uint      `gorm:"primaryKey"`
	QuestionID uint      `gorm:"not null"`
	UserID     uint      `gorm:"not null"`
	SubmitURL  string    `gorm:"size:255;not null"`
	Status     int64     `gorm:"not null"`
	IsFinal    bool      `gorm:"default:true"`
	SubmitTime time.Time `gorm:"not null"`

	// Foreign keys
	Question Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User     User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
