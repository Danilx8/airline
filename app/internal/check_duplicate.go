package internal

import (
	"app/app/domain"
	"strings"
)

// Тупой вариант, если хотите более лучшую версию, то сделаю, на за платно
func CheckDuplicate(records [][]string, formResponse *domain.FormResponse) []bool {
	mask := make([]bool, len(records))
	for i := range records {
		mainRecord := strings.Join(records[i], " ")
		for j := i + 1; j < len(records); j++ {
			secondRecord := strings.Join(records[j], " ")
			if mainRecord == secondRecord {
				mask[i] = true
				formResponse.DuplicateRecords += 1
			}
		}
	}
	return mask
}
