package manager

import (
	"anonymousroom/common"
	"sync"
)

var votes = make(map[string]*votesInfo, 100) //roomId
type votesInfo struct {
	action     func(userId string)
	userNumber int
	votes      []common.Vote
	m          sync.Mutex
}

func CreatVote(roomId string, user int, action func(userId string)) {
	votes[roomId] = &votesInfo{
		action:     action,
		userNumber: user,
		votes:      make([]common.Vote, 0, user),
	}
}

func InsertVote(roomId string, v common.Vote) error {
	vv, ok := votes[roomId]
	if !ok {
		return common.InvalidRoomInfoError
	}

	vv.m.Lock()
	defer vv.m.Unlock()

	vv.votes = append(vv.votes, v)
	if len(vv.votes) == vv.userNumber {
		var voteResult = make(map[string]int)
		for _, vote := range vv.votes {
			voteResult[vote.VotePlayerNumber]++
		}

		ii := ""
		voteV := 0
		for id, i := range voteResult {
			if i > voteV {
				ii = id
			}
		}
		vv.action(ii)
	}
	return nil
}
