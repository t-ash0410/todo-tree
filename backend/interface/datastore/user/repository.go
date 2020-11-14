package user

import (
	"todo-tree/entity"
)

//IUserCommand UserCommandのインターフェース
type IUserCommand interface {
	Create(u entity.User) (id int, err error)
	Update(u entity.User)
	Delete(id int)
}

//IUserQuery UserQueryのインターフェース
type IUserQuery interface {
	FindByID(id int) (user entity.User, err error)
}
