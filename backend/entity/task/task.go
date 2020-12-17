package entity

import "time"

//Task タスク
type Task struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `gorm:"type:varchar(200);not null"`
	Description      string    `gorm:"type:varchar(2000);not null"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-:update"`
}
