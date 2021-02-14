package controller

import (
	"encoding/json"
	"fmt"
	"net"

	"../Entity"
	"../onlineMan"
	"../service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() (controller *UserController) {
	controller = new(UserController)
	controller.userService = service.NewUserService()
	return
}

func (this *UserController) HandleLogin(user *Entity.User, conn net.Conn) (*Entity.ResultType, error) { // request controller
	user2 := this.userService.FindUser(user)
	result := new(Entity.ResultType)
	if user2 != nil {
		result.Status = Entity.OK
		result.Msg = "login success"
		response := &Entity.LoginResponse{UserInfo: *user2, OnlineMap: onlineMan.Manager.GetMap()}
		data_buf, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		onlineMan.Manager.Online(user2, conn) //update before write back //might need to fix later
		result.Data = string(data_buf)
	} else {
		result.Status = Entity.ERROR
		result.Msg = "username or password incorrect"
		result.Data = ""
	}
	return result, nil
}

func (this *UserController) HandleRegister(user *Entity.User) (*Entity.ResultType, error) { // request controller
	result := new(Entity.ResultType)
	result.Status = Entity.OK
	result.Msg = this.userService.Register(user)
	return result, nil
}
