package domain

type Schedules struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	Date       string `gorm:"column:Date"`
	Time       string `gorm:"column:Time"`
	AircraftID int    `gorm:"column:AircraftID"`
	//Aircraft   Aircrafts `gorm:"column:AircraftID;foreignKey:AircraftID"`
	RouteID int `gorm:"column:RouteID"`
	//Route           Route     `gorm:"column:RouteID;foreignKey:RouteID"`
	FlightNumber string  `gorm:"column:FlightNumber"`
	EconomyPrice float64 `gorm:"column:EconomyPrice"`
	//BussinesPrice   float64   `gorm:"-" json:"BussinesPrice"`
	//FirstClassPrice float64   `gorm:"-" json:"FirstClassPrice"`
	Confirmed int `gorm:"column:Confirmed"`
}

type ScheduleRepository interface {
	GetSchedules(schedules *[]Schedules, defaultQuery map[string]string) error
	UpdateFlightByNum(schedule *Schedules) error
	AddRoute(schedule *Schedules) error
	EditRoute(id int, schedule *Schedules) error
	GetIDByFields(routeId int, schedule *Schedules) int
}
