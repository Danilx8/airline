package domain

type Role struct {
	ID    int64  `gorm:"primaryKey"`
	Title string `gorm:"column:Title"`
}
