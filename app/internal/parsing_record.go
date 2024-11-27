package internal

import "app/app/domain"

func ParsingRecord(record []string, recordModel *domain.Record) {
	recordModel.Type = record[0]
	recordModel.Date = record[1]
	recordModel.Time = record[2]
	recordModel.FlightNumber = record[3]
	recordModel.From = record[4]
	recordModel.To = record[5]
	recordModel.AircraftID = record[6]
	recordModel.EconomyPrice = record[7]
	if record[8] == "OK" {
		recordModel.Confirmed = "1"
	}
}
