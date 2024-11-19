package domain

type Schedules struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Date         int       `gorm:"Date"`
	Time         int       `gorm:"Time"`
	AircraftID   int       `gorm:"column:AircraftID"`
	Aircraft     Aircrafts `gorm:"column:AircraftID;foreignKey:AircraftID"`
	RouteID      int       `gorm:"column:RouteID"`
	Route        Routes    `gorm:"column:RouteID;foreignKey:RouteID"`
	FlightNumber int       `gorm:"column:FlightNumber"`
	EconomyPrice int       `gorm:"EconomyPrice"`
	Confirmed    bool      `gorm:"Confirmed"`
}
