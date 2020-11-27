package controller

import (
	"anonymousroom/manager"
	"anonymousroom/module"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

const (
	StrLoginSuccess = "登录成功"
)

func LoginHandle(args []interface{}) {
	log.Debug("enter LoginHandle")
	defer log.Debug("level LoginHandle")

	m := args[0].(*module.Login)
	a := args[1].(gate.Agent)

	mm := manager.Login(m)

	a.WriteMsg(&module.GameMessage{
		Msg:    StrLoginSuccess,
		Data:   mm,
		Type:   "Login",
		Status: 1,
	})
}

func PingHandle(args []interface{}) {
	log.Release("enter PingHandle %+v", args)
}

func UserHandle(args []interface{}) {
	log.Release("enter UserHandle %+v", args)
}

func GameMessageHandle(args []interface{}) {
	log.Release("enter GameMessageHandle %+v", args)
}

func RoomHandle(args []interface{}) {
	log.Debug("enter RoomHandle")
	defer log.Debug("level RoomHandle")

	//m := args[0].(*module.Room)
	//a := args[1].(gate.Agent)
	//
	//a.WriteMsg(&module.GameMessage{
	//	Msg:    "进入房间: ",
	//	Data:   data,
	//	Type:   t,
	//	Status: 1,
	//})
	//
	//Success(a, "进入房间: ", "Room", "success")
}

func GameHandle(args []interface{}) {
	log.Release("enter GameHandle %+v", args)
}

func VoteHandle(args []interface{}) {
	log.Release("enter VoteHandle %+v", args)
}

func RoomOutHandle(args []interface{}) {
	log.Release("enter RoomOutHandle %+v", args)
}

func LogoutHandle(args []interface{}) {
	log.Release("enter LogoutHandle %+v", args)
}
