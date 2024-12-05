package domain

type FormResponse struct {
	SuccessfulChanges       int
	DuplicateRecords        int
	RecordWithMissingFields int
}
