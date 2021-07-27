package gate

import (
	"undercover/game"
	"undercover/msg"
)

func init() {
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Ping{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.User{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GameMessage{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Login{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Room{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Game{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Vote{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RoomOut{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.UserOut{}, game.ChanRPC)
}
