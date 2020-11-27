package manager

import (
	"anonymousroom/module"
	"github.com/satori/go.uuid"
	"time"
)

var users = make(map[string]*module.Login, 100)

func Login(m *module.Login) *module.Login {
	if m == nil {
		m = new(module.Login)
	}
	if m.UserId == "" {
		m.UserId = uuid.NewV4().String()
	} else {
		if mm, ok := users[m.UserId]; ok {
			return mm
		}
	}
	if m.UserName == "" {
		m.UserId = time.Now().String()
	}
	users[m.UserId] = m
	return m
}
