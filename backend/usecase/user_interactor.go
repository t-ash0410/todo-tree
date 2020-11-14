package usecase

import (
	"todo-tree/entity"
	"todo-tree/interface/datastore/user"
)

//UserInteractor ユーザに関するビジネスロジックを定義
type UserInteractor struct {
	Query   user.IUserQuery
	Command user.IUserCommand
}

//Add ユーザの追加
func (interactor *UserInteractor) Add(u entity.User) (err error) {
	_, err = interactor.Command.Create(u)
	if err != nil {
		panic(err)
	}
	return
}

//UserByID ユーザIDで検索して返却
func (interactor *UserInteractor) UserByID(id int) (user entity.User, err error) {
	user, err = interactor.Query.FindByID(id)
	if err != nil {
		panic(err)
	}
	return
}
