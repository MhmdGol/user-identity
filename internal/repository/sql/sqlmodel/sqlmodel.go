package sqlmodel

import "time"

type UserInfo struct {
	ID             string    `gorm:"primaryKey;column:id"`
	UUN            string    `gorm:"column:uun"`
	Username       string    `gorm:"unique;column:username"`
	HashedPassword string    `gorm:"column:hashed_password"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	Email          string    `gorm:"column:email"`
	TotpSecret     string    `gorm:"column:totp_secret"`
	Role           string    `gorm:"column:role"`
	Status         string    `gorm:"column:status"`
}

type Session struct {
	ID         string    `gorm:"primaryKey;column:id"`
	UserID     string    `gorm:"foreignKey:UserID;column:user_id"`
	SessionExp time.Time `gorm:"column:session_exp"`
}

type TrackInfo struct {
	ID        string    `gorm:"primaryKey;column:id"`
	UserID    string    `gorm:"foreignKey:UserID;column:user_id"`
	Action    string    `gorm:"column:action"`
	Timestamp time.Time `gorm:"column:timestamp"`
}
