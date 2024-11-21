package domain

import "time"

type Schedules struct {
	ID           int        `gorm:"primaryKey;autoIncrement"`
	Date         *time.Time `gorm:"Date"`
	Time         *time.Time `gorm:"Time"`
	AircraftID   int        `gorm:"column:AircraftID"`
	Aircraft     Aircrafts  `gorm:"column:AircraftID;foreignKey:AircraftID"`
	RouteID      int        `gorm:"column:RouteID"`
	Route        Route      `gorm:"column:RouteID;foreignKey:RouteID"`
	FlightNumber int        `gorm:"column:FlightNumber"`
	EconomyPrice int        `gorm:"EconomyPrice"`
	Confirmed    bool       `gorm:"Confirmed"`
}
