package domain

type Airports struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CountryID int       `gorm:"column:CountryID"`
	Country   Countries `gorm:"column:CountryID;foreignKey:CountryID"`
	IATACode  string    `gorm:"column:IATACode"`
	Name      string    `gorm:"column:Name"`
}

func (Airports) TableName() string {
	return "Airports"
}
