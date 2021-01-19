package manager

import (
	"anonymousroom/common"
	"fmt"
	"github.com/name5566/leaf/gate"
	"time"
)

func Login(a gate.Agent, m *common.Login) *common.Login {
	if m == nil {
		m = new(common.Login)
	}
	if m.UserId == "" {
		m.UserId = fmt.Sprintf("0%d", time.Now().UnixNano())
	} else {
		if mm, ok := users[m.UserId]; ok {
			return mm
		}
	}
	if m.UserName == "" {
		m.UserId = time.Now().String()
	}
	users[m.UserId] = m
	agents[m.UserId] = a
	return m
}

func GetAgents(userid string) (gate.Agent, error) {
	if userid == "" {
		return nil, common.InvalidUserIDError
	}
	a, ok := agents[userid]
	if ok {
		return a, nil
	}

	return nil, common.InvalidUserIDError
}

func GetLogin(id string) (*common.Login, error) {
	if id == "" {
		return nil, common.InvalidRoomIDError
	}
	x, ok := users[id]
	if ok {
		return x, nil
	}
	return nil, common.InvalidRoomIDError
}
