package model

import (
	"time"

	"github.com/pquerna/otp"
)

type UserInfo struct {
	ID             ID
	UUN            string
	Username       string
	HashedPassword string
	Created_at     time.Time
	Email          string
	TotpSecret     *otp.Key
	Role           string
	Status         string
}

type RawUser struct {
	ID         ID
	UUN        string
	Username   string
	Password   string
	Email      string
	TotpSecret *otp.Key
	Role       string
	Status     string
}
