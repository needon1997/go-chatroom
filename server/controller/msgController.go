package controller

import (
	"../Entity"
	"../onlineMan"
)

type MsgController struct {
}

func (this *MsgController) RedirectMsg(msg *Entity.Message) {
	onlineMan.Manager.RedirectMsg(msg)
}
