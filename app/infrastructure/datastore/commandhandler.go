package datastore

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SQLCommandHandler 変更系のクエリ発行を司る
type SQLCommandHandler struct {
	Conn *gorm.DB
}

//NewCommandHandler SQLCommandHandlerの生成
func NewCommandHandler(connectionString string) *SQLCommandHandler {
	conn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	handler := new(SQLCommandHandler)
	handler.Conn = conn
	return handler
}
