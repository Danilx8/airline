package domain

type Office struct {
	ID        int64  `gorm:"primaryKey"`
	CountryID int64  `gorm:"column:CountryID"`
	Title     string `gorm:"column:Title"`
	Phone     string `gorm:"column:Phone"`
	Contact   string `gorm:"column:Contact"`
}
