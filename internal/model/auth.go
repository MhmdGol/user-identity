package model

type LoginInfo struct {
	Username string
	Password string
}

type UpdatePassword struct {
	Username string
	OldPass  string
	NewPass  string
}
