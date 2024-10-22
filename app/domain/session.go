package domain

import "time"

// References user_panel
type UserPanel struct {
	Id           int       `gorm:"id"`
	UserId       int       `gorm:"user_id"`
	Date         string    `gorm:"date"`
	LoginTime    time.Time `gorm:"login_time"`
	LogoutTime   time.Time `gorm:"logout_time"`
	LogoutReason string    `gorm:"logout_reason"`
}

type SessionRepository interface {
	Start(session *UserPanel) error
	Update(session *UserPanel) error
}
