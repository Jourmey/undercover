package utils

import "sync"

type JDB struct {
	jt  map[int]interface{}
	p   int
	jtM sync.Mutex
}

func NewJDB() *JDB {
	return &JDB{
		jt:  make(map[int]interface{}, 100),
		p:   0,
		jtM: sync.Mutex{},
	}
}

func (j *JDB) Action(filter func(r0 interface{}) bool, action func(id int, r interface{})) {
	for i, r := range j.jt {
		if filter(r) {
			action(i, r)
		}
	}
}

func (j *JDB) Insert(r interface{}) {
	j.jtM.Lock()
	defer j.jtM.Unlock()
	j.p++
	j.jt[j.p] = r
}

func (j *JDB) Delete(id int) {
	delete(j.jt, id)
}
func (j *JDB) Update(id int, i interface{}) {
	j.jt[id] = i
}
