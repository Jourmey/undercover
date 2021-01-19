package manager

import (
	"anonymousroom/common"
	"strconv"
)

func CreatGame(r *common.Room) (*common.Game, error) {
	if r == nil {
		return nil, common.InvalidRoomInfoError
	}

	undercoverNum, _ := strconv.Atoi(r.UndercoverNumber)
	game := new(common.Game)
	game.Round = 1
	game.SurvivalUserList = make(map[string]*common.User)
	game.Keyword = &common.Keyword{
		Code:           "",
		NormalWord:     "枕头",
		UndercoverWord: "抱枕",
		Vension:        0,
	}
	game.UndercoverNum = undercoverNum
	game.Stage = "start"
	game.ActionTime = 60
	game.VoteTime = 60
	game.VoteList = make(map[string]*common.Vote)
	game.RoomId = r.RoomId
	//game.VoteChan = make(chan *common.Vote)
	game.VoteNum = 1
	game.WinRole = ""
	game.OutUser = make([]*common.User, 0)

	games[game.RoomId] = game
	return game, nil
}

func GetGame(roomId string) (*common.Game, error) {
	if roomId == "" {
		return nil, common.InvalidRoomIDError
	}
	g, ok := games[roomId]
	if ok {
		return g, nil
	}
	return nil, common.InvalidRoomIDError
}
