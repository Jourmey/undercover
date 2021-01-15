package manager

import (
	"anonymousroom/common"
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"time"
)

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

	InsertRoomTable([]RoomTable{{
		UserId: l.UserId,
		RoomId: m.RoomId,
	}})

	rooms[m.RoomId] = m
	m.Number = 1
	m.PrepareList = make(map[string]string, 5)
	return m, nil
}

func InRoom(id string, l *common.Login) (*common.Room, error) {
	room, ok := rooms[id]
	if ok {
		InsertRoomTable([]RoomTable{{
			UserId: l.UserId,
			RoomId: id,
		}})
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

	user := SelectRoomTable(func(r0 RoomTable) bool {
		return r0.RoomId == id
	})

	if len(user) == 0 {
		return common.InvalidRoomIDError
	}

	for _, s := range user {
		a, ok := agents[s.UserId]
		if ok {
			f(a)
		} else {
			log.Error("GetAgentOfRoom failed. agents ", s, "is invalid.")
		}
	}
	return nil
}
