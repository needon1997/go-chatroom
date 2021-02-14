package controller

import (
	"encoding/json"
	"fmt"
	"net"

	"../Entity"
)

type Dispatcher struct {
	userController *UserController
	msgControler   *MsgController
}

func NewDispatcher() *Dispatcher {
	dispatcher := new(Dispatcher)
	dispatcher.userController = NewUserController()
	return dispatcher
}
func (this *Dispatcher) HandleMsg(msg *Entity.Message, conn net.Conn) (*Entity.ResultType, error) { //request dispatcher
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	switch msg.Type {
	case Entity.LOGIN:
		var user Entity.User
		err := json.Unmarshal([]byte(msg.Data), &user)
		if err != nil {
			return nil, err
		}
		return this.userController.HandleLogin(&user, conn)
	case Entity.REGISTER:
		var user Entity.User
		err := json.Unmarshal([]byte(msg.Data), &user)
		if err != nil {
			return nil, err
		}
		return this.userController.HandleRegister(&user)
	case Entity.MSG:
		this.msgControler.RedirectMsg(msg)
		return nil, nil
	default:
	}
	return nil, nil
}
