package domain

import "time"

// References user_panel
type Session struct {
	Id           int       `gorm:"ID"`
	UserId       int       `gorm:"UserID"`
	Date         string    `gorm:"Date"`
	LoginTime    time.Time `gorm:"LoginTime"`
	LogoutTime   time.Time `gorm:"LogoutTime"`
	LogoutReason string    `gorm:"LogoutReason"`
}

type SessionRepository interface {
	Start(session *Session) error
	Update(session *Session) error
}

func (Session) TableName() string {
	return "user_panel"
}
