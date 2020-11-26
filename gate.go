package main

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type GateModule struct {
	G *gate.Gate
}

func (m *GateModule) OnInit() {
	log.Debug("enter OnInit")
	defer log.Debug("level OnInit")
}

func (m *GateModule) OnDestroy() {
	log.Debug("enter OnDestroy")
	defer log.Debug("level OnDestroy")
}

func (m *GateModule) Run(closeSig chan bool) {
	log.Debug("enter Run")
	defer log.Debug("level Run")

	m.G.Run(closeSig)
}

func NewGateModule() *GateModule {
	gate := &gate.Gate{
		MaxConnNum:      0,
		PendingWriteNum: 0,
		MaxMsgLen:       0,
		Processor:       Processor,
		AgentChanRPC:    nil,
		WSAddr:          "",
		HTTPTimeout:     0,
		CertFile:        "",
		KeyFile:         "",
		TCPAddr:         "localhost:5678",
		LenMsgLen:       0,
		LittleEndian:    false,
	}

	mm := &GateModule{
		G: gate,
	}

	return mm
}
