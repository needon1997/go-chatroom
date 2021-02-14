package onlineMan

import (
	"encoding/json"
	"net"

	"../Entity"
	"../util"
)

var Manager = &OnlineManager{userConnMap: make(map[Entity.User]net.Conn), connUserMap: make(map[net.Conn]Entity.User)}

type OnlineManager struct {
	userConnMap map[Entity.User]net.Conn
	connUserMap map[net.Conn]Entity.User
}

func (this *OnlineManager) Online(user *Entity.User, conn net.Conn) {
	this.Inform(user, Entity.ONLINE)
	this.userConnMap[*user] = conn
	this.connUserMap[conn] = *user

	return
}

func (this *OnlineManager) Offline(conn net.Conn) {
	user := this.connUserMap[conn]
	delete(this.connUserMap, conn)
	delete(this.userConnMap, user)
	this.Inform(&user, Entity.OFFLINE)

	return
}

func (this *OnlineManager) Inform(user *Entity.User, msgType int) {
	dataBytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	msg := &Entity.Message{Type: msgType, Data: string(dataBytes)}
	for key, _ := range this.connUserMap {
		var transfer *util.Transfer = new(util.Transfer)
		transfer.Conn = key
		transfer.SendMsg(msg)
	}
}

func (this *OnlineManager) RedirectMsg(msg *Entity.Message) {
	for _, val := range this.userConnMap {
		var transfer *util.Transfer = new(util.Transfer)
		transfer.Conn = val
		transfer.SendMsg(msg)
	}

}

func (this *OnlineManager) GetMap() map[string]string {
	userInfoMap := make(map[string]string)
	for key, _ := range this.userConnMap {
		userInfoMap[key.Username] = key.Nickname
	}
	return userInfoMap
}
