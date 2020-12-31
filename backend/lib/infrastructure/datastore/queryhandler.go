package datastore

import (
	"database/sql"
)

//SQLQueryHandler 参照系のクエリ発行を司る
type SQLQueryHandler struct {
	Conn *sql.DB
}

//NewQueryHandler QueryHandlerの生成
func NewQueryHandler(ctx DBContext) SQLQueryHandler {
	conn, err := sql.Open("mysql", ctx.ConnectionString)
	if err != nil {
		panic(err)
	}
	handler := new(SQLQueryHandler)
	handler.Conn = conn
	return *handler
}
