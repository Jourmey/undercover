package manager

import (
	"anonymousroom/common"
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"time"
)

var rooms = make(map[string]*common.Room, 100)
var roomToUser = make(map[string][]string, 100)

func CreatRoom(m *common.Room, l *common.Login) (*common.Room, error) {
	if l == nil {
		return nil, common.LoginEmptyError
	}
	if m == nil {
		m = new(common.Room)
	}

	if m.TotalNumber == "" || m.UndercoverNumber == "" {
		log.Error("creat room failed。")
		return nil, common.InvalidRoomInfoError
	}

	if m.RoomId == "" {
		m.RoomId = fmt.Sprintf("1%d", time.Now().UnixNano())
	}
	m.CreateUserId = l.UserId

	roomToUser[m.RoomId] = []string{l.UserId}
	rooms[m.RoomId] = m
	m.Number = 1
	m.PrepareList = make(map[string]string, 5)
	return m, nil
}

func InRoom(id string, l *common.Login) (*common.Room, error) {
	room, ok := rooms[id]
	if ok {
		roomToUser[id] = append(roomToUser[id], l.UserId)
		room.Number++
		return room, nil
	} else {
		return nil, common.InvalidRoomIDError
	}
}

func GetRoom(id string) (*common.Room, error) {
	if id == "" {
		return nil, common.InvalidRoomIDError
	}
	room, ok := rooms[id]
	if ok {
		return room, nil
	}
	return nil, common.InvalidRoomIDError
}

func SendMessageByRoom(id string, f func(gate.Agent)) error {
	if id == "" {
		return common.InvalidRoomIDError
	}

	user, ok := roomToUser[id]
	if !ok {
		return common.InvalidRoomIDError
	}

	for _, s := range user {
		a, ok := agents[s]
		if ok {
			f(a)
		} else {
			log.Error("GetAgentOfRoom failed. agents ", s, "is invalid.")
		}
	}
	return nil
}
