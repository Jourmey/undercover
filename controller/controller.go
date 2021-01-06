package controller

import (
	"anonymousroom/manager"
	"anonymousroom/module"
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"strconv"
)

func LoginHandle(args []interface{}) {
	log.Debug("enter LoginHandle")
	defer log.Debug("level LoginHandle")

	m := args[0].(*module.Login)
	a := args[1].(gate.Agent)

	mm := manager.Login(a, m)
	a.SetUserData(mm)
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

	m := args[0].(*module.GameMessage)
	a := args[1].(gate.Agent)
	//l := a.UserData().(*module.Login)

	message := &module.GameMessage{
		Msg:    fmt.Sprintf("%s: %s", m.UserName, m.Msg),
		Data:   nil,
		Type:   "Messagee",
		Status: 1,
	}
	err := manager.SendMessageByRoom(m.RoomId, func(agent gate.Agent) {
		agent.WriteMsg(message)
	})

	if err != nil {
		a.WriteMsg(&module.GameMessage{
			Msg:    "房间不存在。" + err.Error(),
			Status: 0,
		})
	}
}

func RoomHandle(args []interface{}) {
	log.Debug("enter RoomHandle")
	defer log.Debug("level RoomHandle")

	m := args[0].(*module.Room)
	a := args[1].(gate.Agent)
	l := a.UserData().(*module.Login)

	var room *module.Room
	var err error
	if m.RoomId == "" {
		room, err = manager.CreatRoom(m, l)
	} else if !m.IsPrepare {
		room, err = manager.InRoom(m.RoomId, l)
	} else { // 准备逻辑
		PrepareHandle(m.RoomId, a, l)
		return
	}

	if err != nil {
		a.WriteMsg(&module.GameMessage{
			Msg:    "创建房间失败！原因：" + err.Error(),
			Data:   nil,
			Type:   "Room",
			Status: 0,
		})
		return
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

func PrepareHandle(roomid string, a gate.Agent, l *module.Login) {
	r, err := manager.GetRoom(roomid)
	if err != nil {
		a.WriteMsg(module.NewErrorGameMessage(err))
	}

	_, isPrepared := r.PrepareList[l.UserName]

	var str string
	if isPrepared {
		str = "取消准备"
		delete(r.PrepareList, l.UserName)
	} else {
		str = "已准备"
		r.PrepareList[l.UserName] = l.UserName
	}

	m := module.NewSuccessGameMessage(l.UserName + ": " + str).WithType("Messagee")
	err = manager.SendMessageByRoom(roomid, func(agent gate.Agent) {
		agent.WriteMsg(m)
	})

	a.WriteMsg(module.NewSuccessGameMessage("").WithType("Prepare"))

	ca, err := manager.GetAgents(r.CreateUserId)
	if err != nil {
		totalNumber, _ := strconv.Atoi(r.TotalNumber)
		if r.PrepareNum == totalNumber-1 && totalNumber == r.Number {
			ca.WriteMsg(module.NewSuccessGameMessage("开始").WithType("Start"))
		} else {
			ca.WriteMsg(module.NewSuccessGameMessage("准备").WithType("Start"))
		}
	}

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
