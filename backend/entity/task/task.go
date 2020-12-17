package entity

import "time"

//Task タスク
type Task struct {
	Id        int       `json:"Id" binding:"-" gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `json:"Name" binding:"required,max=200" gorm:"type:varchar(200);not null"`
	Description      string    `json:"Description" binding:"required,max=2000" gorm:"type:varchar(2000);not null"`
	CreatedAt time.Time `binding:"-" gorm:"<-:create"`
	UpdatedAt time.Time `binding:"-" gorm:"<-:update"`
}
