package util

import (
	"strconv"

	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
)

func CheckPharmacyDrugPagination(page, limit string) (limitInt, offsetInt int, err error) {
	pageInt := 1

	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return 0, 0, apperror.InvalidPageError()
		}
		if pageInt < 1 {
			return 0, 0, apperror.InvalidPageError()
		}
	}

	limitInt = 10

	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, apperror.InvalidLimitError()
		}
		if limitInt < 1 {
			return 0, 0, apperror.InvalidLimitError()
		}
	}

	offsetInt = limitInt * (pageInt - 1)

	return limitInt, offsetInt, nil
}
