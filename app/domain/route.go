package domain

type Route struct {
	ID                 int      `gorm:"primaryKey;autoIncrement"`
	DepartureAirportID int      `gorm:"column:DepartureAirportID"`
	DepartureAirport   Airports `gorm:"column:DepartureAirportID;foreignKey:DepartureAirportID"`
	ArrivalAirportID   int      `gorm:"column:ArrivalAirportID"`
	ArrivalAirport     Airports `gorm:"column:ArrivalAirportID;foreignKey:ArrivalAirportID"`
	Distance           float32  `gorm:"column:Distance"`
	FlightTime         string   `gorm:"column:FlightTime"`
}

func (Route) TableName() string {
	return "Routes"
}

type RouteRepository interface {
	GetByDepartureID(route *Route, depId int) error
	GetByArrivalID(route *Route, arrId int) error
	GetByDepartureIDAndArrivalID(route *Route, depId, arrId int) error
	GetAllRoutes(routes *[]Route) error
}
