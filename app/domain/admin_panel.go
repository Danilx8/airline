package domain

type AdminPanel struct {
	ID     int64 `gorm:"primaryKey;autoIncrement"`
	UserID int64 `gorm:"column:UserID"`
	Age    int64 `gorm:"column:Age"`
}
