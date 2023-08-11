package model

import "time"

type UserInfo struct {
	ID             ID
	UUN            string
	Username       string
	HashedPassword string
	Created_at     time.Time
	Email          string
	Role           string
}

type RawUser struct {
	ID       ID
	UUN      string
	Username string
	Password string
	Email    string
	Role     string
}
