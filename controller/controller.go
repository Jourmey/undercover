package controller

import (
	"anonymousroom/manager"
	"anonymousroom/module"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func LoginHandle(args []interface{}) {
	log.Debug("enter LoginHandle")
	defer log.Debug("level LoginHandle")

	m := args[0].(*module.Login)
	a := args[1].(gate.Agent)

	mm := manager.Login(m)

	a.WriteMsg(&module.GameMessage{
		Msg:    "登录成功！",
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
	log.Debug("enter GameMessageHandle")
	defer log.Debug("level GameMessageHandle")
}

func RoomHandle(args []interface{}) {
	log.Debug("enter RoomHandle")
	defer log.Debug("level RoomHandle")

	m := args[0].(*module.Room)
	a := args[1].(gate.Agent)

	room, err := manager.CreatRoom(m)
	if err != nil {
		a.WriteMsg(&module.GameMessage{
			Msg:    "创建房间失败！",
			Data:   nil,
			Type:   "Room",
			Status: 0,
		})
	}
	var result = make(map[string]interface{})
	result["RoomInfo"] = room
	a.WriteMsg(&module.GameMessage{
		Msg:    "进入房间: " + m.RoomId,
		Data:   result,
		Type:   "Room",
		Status: 1,
	})
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
