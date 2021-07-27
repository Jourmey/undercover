package manager

import (
	"undercover/msg"
	"sync"
)

var votes = make(map[string]*votesInfo, 100) //roomId
type votesInfo struct {
	action     func(v []msg.Vote)
	userNumber int
	votes      []msg.Vote
	m          sync.Mutex
}

func CreatVote(roomId string, user int, action func(v []msg.Vote)) {
	votes[roomId] = &votesInfo{
		action:     action,
		userNumber: user,
		votes:      make([]msg.Vote, 0, user),
	}
}

func InsertVote(roomId string, v msg.Vote) error {
	vv, ok := votes[roomId]
	if !ok {
		return msg.InvalidRoomInfoError
	}

	vv.m.Lock()
	defer vv.m.Unlock()

	vv.votes = append(vv.votes, v)
	if len(vv.votes) == vv.userNumber {

		vv.action(vv.votes)

		defer delete(votes, roomId)
	}
	return nil
}
