package domain

type Route struct {
	ID                 int      `gorm:"primaryKey;autoIncrement"`
	DepartureAirportID int      `gorm:"column:DepartureAirportID"`
	DepartureAirport   Airports `gorm:"column:DepartureAirportID;foreignKey:DepartureAirportID"`
	ArrivalAirportID   int      `gorm:"column:ArrivalAirportID"`
	ArrivalAirport     Airports `gorm:"column:ArrivalAirportID;foreignKey:ArrivalAirportID"`
	Distance           float32  `gorm:"column:Distance"`
	FlightTime         []uint8  `gorm:"column:FlightTime"`
}

type RouteRepository interface {
	GetByDepartureID(depId int) (*Route, error)
	GetByArrivalID(arrId int) (*Route, error)
	GetByDepartureIDAndArrivalID(depId, arrId int) (*Route, error)
}
