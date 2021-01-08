package gata

import (
	"anonymousroom/common"
	"anonymousroom/controller"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/json"
)

type GateModule struct {
	*gate.Gate
}

func (m *GateModule) OnInit() {
	log.Debug("enter OnInit")
	defer log.Debug("level OnInit")
}

func NewGateModule(tcpAddr, wsAddr string) *GateModule {
	processor := getProcessor()

	g := &gate.Gate{
		MaxConnNum:      0,
		PendingWriteNum: 0,
		MaxMsgLen:       0,
		Processor:       processor,
		AgentChanRPC:    nil,
		WSAddr:          wsAddr,
		HTTPTimeout:     0,
		CertFile:        "",
		KeyFile:         "",
		TCPAddr:         tcpAddr,
		LenMsgLen:       0,
		LittleEndian:    false,
	}

	mm := &GateModule{
		g,
	}

	return mm
}

func getProcessor() *json.Processor {
	var processor = json.NewProcessor()
	processor.Register(&common.Ping{})
	processor.Register(&common.User{})
	processor.Register(&common.GameMessage{})
	processor.Register(&common.Login{})
	processor.Register(&common.Room{})
	processor.Register(&common.Game{})
	processor.Register(&common.Vote{})
	processor.Register(&common.RoomOut{})
	processor.Register(&common.Logout{})

	processor.SetHandler(&common.Ping{}, controller.PingHandle)
	processor.SetHandler(&common.User{}, controller.UserHandle)
	processor.SetHandler(&common.GameMessage{}, controller.GameMessageHandle)
	processor.SetHandler(&common.Login{}, controller.LoginHandle)
	processor.SetHandler(&common.Room{}, controller.RoomHandle)
	processor.SetHandler(&common.Game{}, controller.GameHandle)
	processor.SetHandler(&common.Vote{}, controller.VoteHandle)
	processor.SetHandler(&common.RoomOut{}, controller.RoomOutHandle)
	processor.SetHandler(&common.Logout{}, controller.LogoutHandle)
	return processor
}
