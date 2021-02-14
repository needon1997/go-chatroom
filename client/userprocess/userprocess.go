package userprocess

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"../Entity"
	"../online"
	"../onlineMan"
	"../util"
)

const (
	hostAddr    string = "127.0.0.1:8888"
	connectType string = "tcp"
)

func connect() net.Conn {
	conn, err := net.Dial(connectType, hostAddr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return conn
}
func Login(username, password string) (*Entity.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password cannot be empty")
	}
	conn := connect()
	defer func() {
		go OnlineReadDeamon(conn)
		go OnlineWriteDeamon(conn)
	}()
	if conn == nil {
		return nil, errors.New("connection error")
	}
	var transfer *util.Transfer = new(util.Transfer)
	transfer.Conn = conn
	data := Entity.User{Username: username, Password: password}
	data_b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	msg := Entity.Message{Type: Entity.LOGIN, Data: string(data_b)}
	err = transfer.SendMsg(msg)
	if err != nil {
		return nil, err
	}
	result, err := transfer.ReadResult()
	if err != nil {
		return nil, err
	}
	user, err := handleLoginResult(result)
	//use channel to inform when the main should return
	return user, err
}
func Register(username, password, nickname string) error {
	if username == "" || password == "" || nickname == "" {
		return errors.New("username, password, and nickname cannot be empty")
	}
	conn := connect()
	defer conn.Close()
	if conn == nil {
		return errors.New("connection error")
	}
	var transfer *util.Transfer = new(util.Transfer)
	transfer.Conn = conn
	data := Entity.User{Username: username, Password: password, Nickname: nickname}
	data_b, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	msg := Entity.Message{Type: Entity.REGISTER, Data: string(data_b)}
	err = transfer.SendMsg(msg)
	if err != nil {
		return err
	}
	result, err := transfer.ReadResult()
	if err != nil {
		return err
	}
	handleRegisterResult(result)
	//use channel to inform when the main should return
	return nil
}

func OnlineReadDeamon(conn net.Conn) {
	defer conn.Close()
	var transfer *util.Transfer = new(util.Transfer)
	transfer.Conn = conn
	for {
		msg, err := transfer.ReadMsg()
		if err != nil {
			fmt.Println(err)
			continue
		}
		HandleMsg(msg, conn)
	}
}
func OnlineWriteDeamon(conn net.Conn) {
	defer conn.Close()
	var transfer *util.Transfer = new(util.Transfer)
	transfer.Conn = conn
	for {
		msg := <-online.MsgWriteChan
		transfer.SendMsg(msg)
	}
}

func handleLoginResult(result *Entity.ResultType) (user *Entity.User, err error) {
	if result.Status == Entity.OK {
		reponseBody := result.Data
		loginResponse := new(Entity.LoginResponse)
		err = json.Unmarshal([]byte(reponseBody), loginResponse)
		if err != nil {
			return nil, err
		}
		user = &(loginResponse.UserInfo)
		onlineMan.Manager.UserMap = &(loginResponse.OnlineMap)
	}
	return
}

func handleRegisterResult(result *Entity.ResultType) {
	fmt.Println(result.Msg)
	return
}
