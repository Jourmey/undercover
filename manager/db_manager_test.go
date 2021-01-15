package manager

import (
	"testing"
)

var (
	a = RoomTable{
		UserId: "aaa",
		RoomId: "AAAA",
	}
	b = RoomTable{
		UserId: "bbb",
		RoomId: "BBBB",
	}
	c = RoomTable{
		UserId: "ccc",
		RoomId: "CCCC",
	}
	d = RoomTable{
		UserId: "ddd",
		RoomId: "DDDDD",
	}
)

func TestInsertRoomTable(t *testing.T) {
	aa := []RoomTable{a, b, c, d}
	InsertRoomTable(aa)

	bb := SelectRoomTable(func(r0 RoomTable) bool {
		return true
	})
	if len(bb) != len(aa) {
		t.Fatal()
	}

	bb[0].RoomId = "XXX"

	cc := SelectRoomTable(func(r0 RoomTable) bool {
		return true
	})

	t.Log(cc)
}

func TestUpdateRoomTable(t *testing.T) {
	aa := []RoomTable{a, b, c, d}
	InsertRoomTable(aa)

	UpdateRoomTable(func(r0 RoomTable) bool {
		return r0.RoomId == "DDDDD" || r0.RoomId == "CCCC"
	}, func(r RoomTable) RoomTable {
		r.RoomId += "XXX"
		return r
	})

	SelectRoomTable(func(_ RoomTable) bool {
		return true
	})
}

func TestDeleteRoomTable(t *testing.T) {
	aa := []RoomTable{a, b, c, d}
	InsertRoomTable(aa)

	DeleteRoomTable(func(r0 RoomTable) bool {
		return r0.RoomId == "DDDDD" || r0.RoomId == "CCCC"
	})

	SelectRoomTable(func(_ RoomTable) bool {
		return true
	})
}
