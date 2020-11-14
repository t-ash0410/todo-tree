package user

import (
	"todo-tree/entity"
	"todo-tree/infrastructure/datastore"
)

//DBUserCommand Userの変更系クエリを司る
type DBUserCommand struct {
	*datastore.SQLCommandHandler
}

//NewDBUserCommand DBUserCommandのコンストラクタ
func NewDBUserCommand(commandHandler *datastore.SQLCommandHandler) *DBUserCommand {
	return &DBUserCommand{
		commandHandler,
	}
}

//Create ユーザの作成
func (repo *DBUserCommand) Create(u entity.User) (id int, err error) {
	result := repo.Conn.Create(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	return u.ID, nil
}

//Update ユーザの更新
func (repo *DBUserCommand) Update(u entity.User) {
	repo.Conn.First(&u)
	repo.Conn.Save(&u)
}

//Delete ユーザの削除
func (repo *DBUserCommand) Delete(id int) {
	repo.Conn.Delete(&entity.User{}, id)
}
