package user

//User ユーザ
type User struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(200);not null"`
}
