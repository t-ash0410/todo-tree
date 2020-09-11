package entity

//User ユーザ
type User struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Email string
}
