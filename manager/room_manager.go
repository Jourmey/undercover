package manager

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"undercover/msg"
	"time"
)

func CreatRoom(m *msg.Room, l string) (*msg.Room, error) {
	if l == "" {
		return nil, msg.LoginEmptyError
	}
	if m == nil {
		m = new(msg.Room)
	}

	if m.TotalNumber == "" || m.UndercoverNumber == "" {
		log.Error("creat room failed。")
		return nil, msg.InvalidRoomInfoError
	}

	if m.RoomId == "" {
		m.RoomId = fmt.Sprintf("1%d", time.Now().UnixNano())
	}
	m.CreateUserId = l

	InsertRoomTable([]RoomTable{{
		UserId: l,
		RoomId: m.RoomId,
	}})

	rooms[m.RoomId] = m
	m.Number = 1
	//m.PrepareList = make(map[string]string, 5)
	return m, nil
}

func InRoom(id string, l string) (*msg.Room, error) {
	room, ok := rooms[id]
	if ok {
		InsertRoomTable([]RoomTable{{
			UserId: l,
			RoomId: id,
		}})
		room.Number++
		return room, nil
	} else {
		return nil, msg.InvalidRoomIDError
	}
}

func GetRoom(id string) (*msg.Room, error) {
	if id == "" {
		return nil, msg.InvalidRoomIDError
	}
	room, ok := rooms[id]
	if ok {
		return room, nil
	}
	return nil, msg.InvalidRoomIDError
}

func SendMessageByRoom(id string, f func(gate.Agent)) error {
	if id == "" {
		return msg.InvalidRoomIDError
	}

	user := SelectRoomTable(func(r0 RoomTable) bool {
		return r0.RoomId == id
	})

	if len(user) == 0 {
		return msg.InvalidRoomIDError
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
