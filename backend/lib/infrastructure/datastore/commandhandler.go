package datastore

import (
	entity "github.com/todo-tree/entity/task"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SQLCommandHandler 変更系のクエリ発行を司る
type SQLCommandHandler struct {
	Conn *gorm.DB
}

//NewCommandHandler SQLCommandHandlerの生成
func NewCommandHandler(ctx DBContext) SQLCommandHandler {
	conn, err := gorm.Open(mysql.Open(ctx.ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	conn.Set("gorm:table_options", "ENGINE=InnoDB")
	migrate(conn)
	handler := new(SQLCommandHandler)
	handler.Conn = conn
	return *handler
}

//migrate アプリケーション内で使用するエンティティのマイグレーションを自動化
func migrate(conn *gorm.DB) {
	conn.AutoMigrate(&entity.Task{})
}
