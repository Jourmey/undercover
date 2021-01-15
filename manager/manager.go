package manager

import (
	"anonymousroom/common"
	"github.com/name5566/leaf/gate"
)

var (
	users  = make(map[string]*common.Login, 100) //UserId
	agents = make(map[string]gate.Agent, 100)    //UserId
	games  = make(map[string]*common.Game, 100)  //RoomId
	rooms  = make(map[string]*common.Room, 100)  //RoomId

)
