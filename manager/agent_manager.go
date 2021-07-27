package manager

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"undercover/msg"
	"time"
)

func Login(a gate.Agent, m *msg.Login) *msg.Login {
	if m == nil {
		m = new(msg.Login)
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
		return nil, msg.InvalidUserIDError
	}
	a, ok := agents[userid]
	if ok {
		return a, nil
	}

	return nil, msg.InvalidUserIDError
}

func GetLogin(id string) (*msg.Login, error) {
	if id == "" {
		return nil, msg.InvalidRoomIDError
	}
	x, ok := users[id]
	if ok {
		return x, nil
	}
	return nil, msg.InvalidRoomIDError
}
