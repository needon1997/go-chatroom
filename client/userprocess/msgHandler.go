package userprocess

import (
	"encoding/json"
	"fmt"
	"net"

	"../Entity"
	"../onlineMan"
)

func HandleMsg(msg *Entity.Message, conn net.Conn) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	switch msg.Type {
	case Entity.ONLINE:
		var user Entity.User
		err := json.Unmarshal([]byte(msg.Data), &user)
		if err != nil {
			return
		}
		onlineMan.Manager.Add(user)
		return
	case Entity.OFFLINE:
		var user Entity.User
		err := json.Unmarshal([]byte(msg.Data), &user)
		if err != nil {
			return
		}
		onlineMan.Manager.Delete(user)
		return
	case Entity.MSG:
		onlineMan.Manager.Append(msg)
	default:
	}
	return
}
