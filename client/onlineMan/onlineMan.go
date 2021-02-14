package onlineMan

import (
	"fmt"

	Entity "../Entity"
)

var Manager = &OnlineManager{MsgList: make([]string, 50)}

type OnlineManager struct {
	UserMap *map[string]string
	MsgList []string
}

func (this *OnlineManager) GetTotal() int {
	return len(*this.UserMap)
}
func (this *OnlineManager) String() string {
	return fmt.Sprintln(*(this.UserMap))
}
func (this *OnlineManager) Add(user Entity.User) {
	(*this.UserMap)[user.Username] = user.Nickname
}
func (this *OnlineManager) Delete(user Entity.User) {
	delete(*this.UserMap, user.Username)
}
func (this *OnlineManager) Append(msg *Entity.Message) {
	this.MsgList = append(this.MsgList, msg.Data)
}
func (this *OnlineManager) ShowMsg() {
	for i := 0; i < len(this.MsgList); i++ {
		fmt.Println(this.MsgList[i])
	}
}
