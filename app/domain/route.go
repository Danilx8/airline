package domain

type Routes struct {
	ID                int      `gorm:"primaryKey;autoIncrement"`
	DepatureAirportID int      `gorm:"column:DepatureAirportID"`
	DepatureAirport   Airports `gorm:"column:DepatureAirportID;foreignKey:DepatureAirportID"`
	ArrivalAirportID  int      `gorm:"column:DepatureAirportID"`
	ArrivalAirport    Airports `gorm:"column:ArrivalAirportID;foreignKey:ArrivalAirportID"`
	Distance          float32  `gorm:"column:Distance"`
	FlightTime        int      `gorm:"column:FlightTime"`
}
