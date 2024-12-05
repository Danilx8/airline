package internal

import (
	"app/app/domain"
	"strings"
)

// Тупой вариант, если хотите более лучшую версию, то сделаю, на за платно
func CheckDuplicate(records [][]string, formResponse *domain.FormResponse) []bool {
	mask := make([]bool, len(records))
	seen := make(map[string]int)

	for i := range records {
		mainRecord := strings.Join(records[i], " ")
		seen[mainRecord]++
	}

	for i := range records {
		mainRecord := strings.Join(records[i], " ")
		if seen[mainRecord] > 1 {
			mask[i] = true
			formResponse.DuplicateRecords++
		}
	}

	return mask
}
