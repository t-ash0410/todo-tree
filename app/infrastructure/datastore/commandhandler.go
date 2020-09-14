package datastore

import (
	"todo-tree/app/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SQLCommandHandler 変更系のクエリ発行を司る
type SQLCommandHandler struct {
	Conn *gorm.DB
}

//NewCommandHandler SQLCommandHandlerの生成
func NewCommandHandler(ctx DBContext) *SQLCommandHandler {
	conn, err := gorm.Open(mysql.Open(ctx.ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	conn.Set("gorm:table_options", "ENGINE=InnoDB")
	SetMigrate(conn)
	handler := new(SQLCommandHandler)
	handler.Conn = conn
	return handler
}

//SetMigrate アプリケーション内で使用するエンティティのマイグレーションを自動化
func SetMigrate(conn *gorm.DB) {
	conn.AutoMigrate(&entity.User{})
}
