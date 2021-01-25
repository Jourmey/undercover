package manager

import (
	"anonymousroom/utils"
)

type RoomTable struct {
	UserId       string
	RoomId       string
	IsUndercover bool //是否为卧底
	IsPrepare    bool //是否准备
	IsOut        bool //是否出局
}

var (
	rt = utils.NewJDB() //RoomTable
)

func InsertRoomTable(r []RoomTable) {
	for _, table := range r {
		rt.Insert(table)
	}
}

func actionRoomTable(filter func(r0 RoomTable) bool, action func(id int, r RoomTable)) {
	rt.Action(func(r0 interface{}) bool {
		return filter(r0.(RoomTable))
	}, func(id int, r interface{}) {
		action(id, r.(RoomTable))
	})
}

func DeleteRoomTable(filter func(r0 RoomTable) bool) {
	actionRoomTable(filter, func(id int, _ RoomTable) {
		rt.Delete(id)
	})
}

func UpdateRoomTable(filter func(r0 RoomTable) bool, action func(r RoomTable) RoomTable) {
	actionRoomTable(filter, func(id int, r RoomTable) {
		rt.Update(id, action(r))
	})
}

func SelectRoomTable(filter func(r0 RoomTable) bool) []RoomTable {
	result := make([]RoomTable, 0, 2)
	actionRoomTable(filter, func(id int, r RoomTable) {
		result = append(result, r)
	})
	return result
}
