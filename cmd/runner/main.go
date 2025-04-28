package main

import (
	"fmt"
	"log"
	"time"

	// مسیر صحیح پکیج entity خودت رو بذار

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=12345678 dbname=judgino_db port=5433 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// ساخت ۵۰۰ تا submission
	var submissions []Submission
	now := time.Now()

	for i := 1; i <= 500; i++ {
		submission := Submission{
			QuestionID: 1, // سوال های فرضی بین ۱ تا ۱۰
			UserID:     1, // کاربر فرضی بین ۱ تا ۱۰۰۰
			SubmitURL:  "./uploads/code/1/main_21312.go",
			Status:     1, // ۱ یعنی "آماده برای جاج شدن"
			IsFinal:    true,
			SubmitTime: now.Add(time.Duration(-i) * time.Minute), // هر سابمیشن یه دقیقه قبل تر
			TryCount:   0,
		}
		submissions = append(submissions, submission)
	}

	// Bulk Insert
	if err := db.Create(&submissions).Error; err != nil {
		log.Fatalf("failed to insert submissions: %v", err)
	}

	fmt.Println("✅ 500 submissions inserted successfully.")
}

type Submission struct {
	ID         int64     `gorm:"column:id;primaryKey"`
	QuestionID int64     `gorm:"column:question_id"`
	UserID     int64     `gorm:"column:user_id"`
	SubmitURL  string    `gorm:"column:submit_url"`
	Status     int64     `gorm:"column:status"`
	IsFinal    bool      `gorm:"column:is_final"`
	SubmitTime time.Time `gorm:"column:submit_time"`
	TryCount   int       `gorm:"column:try_count"`
}
