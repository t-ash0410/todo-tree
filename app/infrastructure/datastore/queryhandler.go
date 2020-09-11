package datastore

import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

//SQLQueryHandler 参照系のクエリ発行を司る
type SQLQueryHandler struct {
	Conn *sql.DB
}

//NewQueryHandler QueryHandlerの生成
func NewQueryHandler(connectionString string) *SQLQueryHandler {
	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	handler := new(SQLQueryHandler)
	handler.Conn = conn
	return handler
}
