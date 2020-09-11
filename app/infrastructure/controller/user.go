package controller

import (
	"strconv"
	"todo-tree/app/entity"
	"todo-tree/app/interface/controller"
	"todo-tree/app/interface/datastore/user"
	"todo-tree/app/usecase"
)

//UserController Userに関する操作の窓口
type UserController struct {
	Interactor usecase.UserInteractor
}

//NewUserController UserControllerのコンストラクタ
func NewUserController(command user.IUserCommand, query user.IUserQuery) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			Command: command,
			Query:   query,
		},
	}
}

//Create Userの追加
func (controller *UserController) Create(c controller.Context) {
	u := entity.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, nil)
}

//Get Userの取得
func (controller *UserController) Get(c controller.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, user)
}
