package user

import (
	"todo-tree/app/entity"
	"todo-tree/app/infrastructure/datastore"
)

//DBUserQuery Userの参照系クエリを司る
type DBUserQuery struct {
	*datastore.SQLQueryHandler
}

//NewDBUserQuery DBUserQueryのコンストラクタ
func NewDBUserQuery(queryHandler *datastore.SQLQueryHandler) *DBUserQuery {
	return &DBUserQuery{
		queryHandler,
	}
}

//FindByID UserをIDで検索して返却
func (query *DBUserQuery) FindByID(userID int) (user entity.User, err error) {
	row, err := query.Conn.Query("SELECT id, name, email FROM users WHERE id = ?", userID)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	var email string
	row.Next()
	if err = row.Scan(&id, &name, &email); err != nil {
		return
	}
	user.ID = id
	user.Name = name
	user.Email = email
	return
}
