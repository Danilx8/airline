package domain

type Route struct {
	ID                int      `gorm:"primaryKey;autoIncrement"`
	DepatureAirportID int      `gorm:"column:DepatureAirportID"`
	DepatureAirport   Airports `gorm:"column:DepatureAirportID;foreignKey:DepatureAirportID"`
	ArrivalAirportID  int      `gorm:"column:DepatureAirportID"`
	ArrivalAirport    Airports `gorm:"column:ArrivalAirportID;foreignKey:ArrivalAirportID"`
	Distance          float32  `gorm:"column:Distance"`
	FlightTime        int      `gorm:"column:FlightTime"`
}

type RouteRepository interface {
	GetByDepatureID(depId int) *Route
	GetByArrivalID(arrId int) *Route
	GetByDepatureIDAndArrivalID(depId, arrId int) *Route
}
