package gata

import (
	"anonymousroom/controller"
	"anonymousroom/module"
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
	processor.Register(&module.Ping{})
	processor.Register(&module.User{})
	processor.Register(&module.GameMessage{})
	processor.Register(&module.Login{})
	processor.Register(&module.Room{})
	processor.Register(&module.Game{})
	processor.Register(&module.Vote{})
	processor.Register(&module.RoomOut{})
	processor.Register(&module.Logout{})

	processor.SetHandler(&module.Ping{}, controller.PingHandle)
	processor.SetHandler(&module.User{}, controller.UserHandle)
	processor.SetHandler(&module.GameMessage{}, controller.GameMessageHandle)
	processor.SetHandler(&module.Login{}, controller.LoginHandle)
	processor.SetHandler(&module.Room{}, controller.RoomHandle)
	processor.SetHandler(&module.Game{}, controller.GameHandle)
	processor.SetHandler(&module.Vote{}, controller.VoteHandle)
	processor.SetHandler(&module.RoomOut{}, controller.RoomOutHandle)
	processor.SetHandler(&module.Logout{}, controller.LogoutHandle)
	return processor
}
