package entity



type User struct {
    ID                   uint   `gorm:"primaryKey"`
    FirstName            string `gorm:"size:255;not null"`
    Email                string `gorm:"size:255;unique"`
    Phone                string `gorm:"size:11;not null;unique"`
    Username             string  `gorm:"not null;unique"`
    Password             string `gorm:"size:255"`
    Role                 string `gorm:"size:255;not null"`
    CreatedQuestionsCount int64 `gorm:"not null"`
    SolvedQuestionsCount  int64 `gorm:"not null"`
}