package entity

import "time"

//User ユーザ
type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `gorm:"type:varchar(200);not null"`
	Email     string    `gorm:"type:varchar(200);not null"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-:update"`
}
