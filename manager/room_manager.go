package manager

import (
	"anonymousroom/module"
	uuid "github.com/satori/go.uuid"
)

var rooms = make(map[string]*module.Room, 100)
var roomToUser = make(map[string][]string, 100)

func CreatRoom(m *module.Room) (*module.Room, error) {
	if m == nil {
		m = new(module.Room)
	}
	if m.RoomId == "" {
		m.RoomId = uuid.NewV4().String()
	} else {
		if mm, ok := rooms[m.RoomId]; ok {
			return mm, nil
		}
	}
	m.CreateUserId = ""

	roomToUser[m.RoomId] = []string{}
	rooms[m.RoomId] = m
	return m, nil
}
