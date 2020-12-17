package entity

//User ユーザ
type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `gorm:"type:varchar(200);not null"`
}
