package domain

import "time"

type Session struct {
	Id           int       `json:"id"`
	UserId       int       `json:"user_id"`
	Date         string    `json:"date"`
	LoginTime    time.Time `json:"login_time"`
	LogoutTime   time.Time `json:"logout_time"`
	LogoutReason string    `json:"logout_reason"`
}
