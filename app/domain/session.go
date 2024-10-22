package domain

import "time"

// References user_panel
type Session struct {
	Id           int       `gorm:"column:ID"`
	UserId       int       `gorm:"column:UserID"`
	Date         string    `gorm:"column:Date"`
	LoginTime    time.Time `gorm:"column:LoginTime"`
	LogoutTime   time.Time `gorm:"column:LogoutTime"`
	LogoutReason string    `gorm:"column:LogoutReason"`
}

type SessionRepository interface {
	Start(session *Session) error
	Update(session *Session) error
}

func (Session) TableName() string {
	return "user_panel"
}
