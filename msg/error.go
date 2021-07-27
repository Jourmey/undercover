package msg

import "errors"

var (
	LoginEmptyError      = errors.New("login is empty.")
	InvalidRoomIDError   = errors.New("invalid roomid.")
	InvalidUserIDError   = errors.New("invalid user id.")
	InvalidRoomInfoError = errors.New("invalid room info.")
)
