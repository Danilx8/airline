package domain

type Record struct {
	Type         string
	Date         string
	Time         string
	FlightNumber string
	From         string
	To           string
	AircraftID   int
	EconomyPrice float64
	Confirmed    int
}
