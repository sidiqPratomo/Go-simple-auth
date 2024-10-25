package util

import (
	"strconv"

	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
)

type QueryParam struct {
	Id               int
	Sort             string
	SortBy           string
	Page             string
	Limit            string
	Offset           string
	SpecializationId string
	Search           string
	PharmacyId       string
}

func SetDefaultQueryParams(params QueryParam) (QueryParam, error) {
	if params.Page == "" {
		params.Page = "1"
	} else {
		pageInt, err := strconv.Atoi(params.Page)
		if err != nil || pageInt < 1 {
			return QueryParam{}, apperror.InvalidPageError()
		}
	}

	if params.Limit == "" {
		params.Limit = "12"
	} else {
		limitInt, err := strconv.Atoi(params.Limit)
		if err != nil || limitInt < 1 {
			return QueryParam{}, apperror.InvalidLimitError()
		}
	}

	pageInt, err := strconv.Atoi(params.Page)
	if err != nil {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(params.Limit)
	if err != nil {
		limitInt = 12
	}

	offsetInt := (pageInt - 1) * limitInt
	if offsetInt < 0 {
		offsetInt = 0
	}
	params.Offset = strconv.Itoa(offsetInt)

	return params, nil
}
