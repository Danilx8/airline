package internal

import (
	"app/app/domain"
	"strconv"
)

func ParsingRecord(record []string, recordModel *domain.Record) error {
	recordModel.Type = record[0]
	recordModel.Date = record[1]
	recordModel.Time = record[2]
	recordModel.FlightNumber = record[3]
	recordModel.From = record[4]
	recordModel.To = record[5]
	aircraftId, err := strconv.Atoi(record[6])
	if err != nil {
		return err
	}
	recordModel.AircraftID = aircraftId
	economyPrice, err := strconv.ParseFloat(record[7], 64)
	if err != nil {
		return err
	}
	recordModel.EconomyPrice = economyPrice
	if record[8] == "OK" {
		recordModel.Confirmed = 1
	} else {
		recordModel.Confirmed = 0
	}
	return nil
}
