package service

import (
	"../Entity"
	"../dao"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() (service *UserService) {
	service = new(UserService)
	service.userDao = dao.NewUserDao()
	return
}

func (this *UserService) FindUser(user *Entity.User) *Entity.User {
	result := this.userDao.FindUser(user)
	if result != nil && result.Password == user.Password {
		return result
	} else {
		return nil
	}
}

const (
	USER_EXIST       = "username duplicate"
	SERVER_ERROR     = "server error"
	REGISTER_SUCCESS = "register success"
)

func (this *UserService) Register(user *Entity.User) string {
	user2 := this.FindUser(user)
	if user2 != nil {
		return USER_EXIST
	}
	success := this.userDao.AddUser(user)
	if success {
		return REGISTER_SUCCESS
	} else {
		return SERVER_ERROR
	}
}
