package manager

import (
	"github.com/name5566/leaf/gate"
	"undercover/msg"
)

var (
	users  = make(map[string]*msg.Login, 100) //UserId
	agents = make(map[string]gate.Agent, 100) //UserId
	games  = make(map[string]*msg.Game, 100)  //RoomId
	rooms  = make(map[string]*msg.Room, 100)  //RoomId
)
