package manager

import (
	"undercover/msg"
	"strconv"
)

func CreatGame(r *msg.Room) (*msg.Game, error) {
	if r == nil {
		return nil, msg.InvalidRoomInfoError
	}

	undercoverNum, _ := strconv.Atoi(r.UndercoverNumber)
	game := new(msg.Game)
	game.Round = 1
	game.SurvivalUserList = make(map[string]*msg.User)
	game.Keyword = &msg.Keyword{
		Code:           "",
		NormalWord:     "枕头",
		UndercoverWord: "抱枕",
		Vension:        0,
	}
	game.UndercoverNum = undercoverNum
	game.Stage = "start"
	game.ActionTime = 60
	game.VoteTime = 60
	//game.VoteList = make(map[string]*msg.Vote)
	game.RoomId = r.RoomId
	//game.VoteChan = make(chan *msg.Vote)
	//game.VoteNum = 1
	game.WinRole = ""
	//game.OutUser = make([]*msg.User, 0)

	games[game.RoomId] = game
	return game, nil
}

func GetGame(roomId string) (*msg.Game, error) {
	if roomId == "" {
		return nil, msg.InvalidRoomIDError
	}
	g, ok := games[roomId]
	if ok {
		return g, nil
	}
	return nil, msg.InvalidRoomIDError
}
