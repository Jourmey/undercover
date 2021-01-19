package controller

import (
	"anonymousroom/common"
	"anonymousroom/manager"
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"strconv"
)

func LoginHandle(args []interface{}) {
	log.Debug("enter LoginHandle")
	defer log.Debug("level LoginHandle")

	m := args[0].(*common.Login)
	a := args[1].(gate.Agent)

	mm := manager.Login(a, m)
	a.SetUserData(&common.UserData{L: mm.UserId})
	a.WriteMsg(&common.GameMessage{
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

	m := args[0].(*common.GameMessage)
	a := args[1].(gate.Agent)
	//l := a.UserData().(*module.Login)

	message := &common.GameMessage{
		Msg:    fmt.Sprintf("%s: %s", m.UserName, m.Msg),
		Data:   nil,
		Type:   "Messagee",
		Status: 1,
	}
	err := manager.SendMessageByRoom(m.RoomId, func(agent gate.Agent) {
		agent.WriteMsg(message)
	})

	if err != nil {
		a.WriteMsg(&common.GameMessage{
			Msg:    "房间不存在。" + err.Error(),
			Status: 0,
		})
	}
}

func RoomHandle(args []interface{}) {
	log.Debug("enter RoomHandle")
	defer log.Debug("level RoomHandle")

	m := args[0].(*common.Room)
	a := args[1].(gate.Agent)
	l := a.UserData().(*common.UserData)

	var room *common.Room
	var err error
	if m.RoomId == "" {
		room, err = manager.CreatRoom(m, l.L)
	} else if !m.IsPrepare {
		room, err = manager.InRoom(m.RoomId, l.L)
	} else { // 准备逻辑
		prepareHandle(m.RoomId, a, l)
		return
	}

	if err != nil {
		a.WriteMsg(&common.GameMessage{
			Msg:    "创建房间失败！原因：" + err.Error(),
			Data:   nil,
			Type:   "Room",
			Status: 0,
		})
		return
	}
	var result = make(map[string]interface{})
	result["RoomInfo"] = room
	a.WriteMsg(&common.GameMessage{
		Msg:    "进入房间: " + m.RoomId,
		Data:   result,
		Type:   "Room",
		Status: 1,
	})

	l.R = room.RoomId
}

func GameHandle(args []interface{}) {
	log.Debug("enter GameHandle()")
	defer log.Debug("level GameHandle()")

	m := args[0].(*common.Game)
	a := args[1].(gate.Agent)
	l := a.UserData().(*common.UserData)

	switch m.Stage {
	case common.GameStage_Start:
		err := gameStart(m, a, l)
		if err != nil {
			log.Error("GameHandle. gameStart(m, a, l) failed. err = ", err)
		}
		break
	case common.GameStage_Vote:
		err := vote(m, a, l)
		if err != nil {
			log.Error("GameHandle. gameStart(m, a, l) failed. err = ", err)
		}
		break
	}

}

func VoteHandle(args []interface{}) {
	log.Debug("enter VoteHandle()")
	defer log.Debug("level VoteHandle()")

	m := args[0].(*common.Vote)
	a := args[1].(gate.Agent)
	l := a.UserData().(*common.UserData)

	g, _ := manager.GetGame(l.R)
	if g.Stage != common.GameStage_Vote {
		log.Error("VoteHandle failed. game.stage isn't vote")
		return
	}

	err := manager.InsertVote(l.R, *m)
	if err != nil {
		a.WriteMsg(common.NewErrorGameMessage(err).WithType("VoteSuccess"))
	} else {
		a.WriteMsg(common.NewSuccessGameMessage("投票成功").WithType("VoteSuccess"))
	}
}

func RoomOutHandle(args []interface{}) {
	log.Release("enter RoomOutHandle %+v", args)
}

func LogoutHandle(args []interface{}) {
	log.Release("enter LogoutHandle %+v", args)
}

func prepareHandle(roomid string, a gate.Agent, l *common.UserData) {
	log.Debug("enter prepareHandle()")
	defer log.Debug("level prepareHandle()")

	r, err := manager.GetRoom(roomid)
	if err != nil {
		a.WriteMsg(common.NewErrorGameMessage(err))
	}

	_, isPrepared := r.PrepareList[l.L]

	var str string
	if isPrepared {
		str = "取消准备"
		delete(r.PrepareList, l.L)
	} else {
		str = "已准备"
		r.PrepareList[l.L] = l.L
	}

	m := common.NewSuccessGameMessage(l.L + ": " + str).WithType("Messagee")
	err = manager.SendMessageByRoom(roomid, func(agent gate.Agent) {
		agent.WriteMsg(m)
	})

	a.WriteMsg(common.NewSuccessGameMessage("").WithType("Prepare"))

	ca, err := manager.GetAgents(r.CreateUserId)
	if err != nil {
		return
	}
	totalNumber, _ := strconv.Atoi(r.TotalNumber)
	if len(r.PrepareList) == totalNumber-1 && totalNumber == r.Number {
		ca.WriteMsg(common.NewSuccessGameMessage("开始").WithType("Start"))
	} else {
		ca.WriteMsg(common.NewSuccessGameMessage("准备").WithType("Start"))
	}
}

func gameStart(m *common.Game, a gate.Agent, l *common.UserData) error {
	r, err := manager.GetRoom(m.RoomId)
	if err != nil {
		return err
	}
	g, err := manager.CreatGame(r)
	if err != nil {
		return err
	}

	//l.G = g

	isU := false
	return manager.SendMessageByRoom(m.RoomId, func(agent gate.Agent) {
		k := new(common.KeywordResult)
		if !isU {
			k.Keyword = g.Keyword.UndercoverWord
			isU = true
		} else {
			k.Keyword = g.Keyword.NormalWord
		}
		agent.WriteMsg(common.NewSuccessGameMessage("游戏开始").WithType("StartGame").WithData(k))
	})
}

func vote(m *common.Game, a gate.Agent, l *common.UserData) error {
	log.Debug("enter vote()")
	defer log.Debug("level vote()")

	r, err := manager.GetGame(l.R)
	if err != nil {
		return err
	}

	r.Stage = common.GameStage_Vote
	// 构造投票对象
	s := make(map[string]*common.User)
	rs := manager.SelectRoomTable(func(r0 manager.RoomTable) bool {
		return r0.RoomId == r.RoomId
	})
	for _, rt := range rs {
		l, err := manager.GetLogin(rt.UserId)
		if err != nil {
			continue
		}
		s[rt.UserId] = &common.User{
			Openid: l.UserId,
			No:     l.UserName,
			Name:   l.UserName,
			Status: 0,
			RoomId: r.RoomId,
		}
	}
	r.SurvivalUserList = s

	manager.CreatVote(l.R, len(s), func(userId string) {
		//投出了userid
		voteOver(r, userId)
	})

	return manager.SendMessageByRoom(l.R, func(agent gate.Agent) {
		agent.WriteMsg(common.NewSuccessGameMessage("开始投票").WithType("Vote").WithData(r))
	})
}

func voteOver(g *common.Game, id string) {
	if g.Stage != common.GameStage_Vote {
		return
	}

}
