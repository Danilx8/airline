package internal

import (
	"app/app/domain"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
)

func ReadRecords(r io.Reader, formResponse *domain.FormResponse) ([][]string, error) {
	var err error
	var record []string
	reader := csv.NewReader(r)
	reader.Comma = ','

	records := make([][]string, 0, 10)
	for {
		formResponse.SuccessfulChanges += 1
		record, err = reader.Read()
		if errors.Is(err, io.EOF) {
			fmt.Println(err)
			formResponse.SuccessfulChanges -= 1
			break
		} else if errors.Is(err, csv.ErrFieldCount) {
			formResponse.RecordWithMissingFields += 1
			continue
		} else if err != nil {
			fmt.Println(err)
			break
		}
		records = append(records, record)
	}
	return records, err
}
