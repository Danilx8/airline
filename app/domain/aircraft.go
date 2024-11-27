package domain

type Aircrafts struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"column:Name"`
	MakeModel     string `gorm:"column:MakeModel"`
	TotalSeats    int    `gorm:"column:TotalSeats"`
	EconomySeats  int    `gorm:"column:EconomySeats"`
	BusinessSeats int    `gorm:"column:BusinessSeats"`
}

func (Aircrafts) TableName() string {
	return "Aircrafts"
}
