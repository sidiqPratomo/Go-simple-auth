package util

import (

	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
)

type QueryParam struct {
	Id         int
	SortBy     string
	SortOrder  string
	Page       int
	Limit      int32
	Offset     int32
	Search     string
	Status     *int8
}

func SetDefaultQueryParams(raw QueryParam) (QueryParam, error) {
	params := raw

	// Default & validation for Limit
	if raw.Limit == 0 {
		params.Limit = 12
	} else if raw.Limit < 1 {
		return QueryParam{}, apperror.InvalidLimitError()
	}

	// Default & validation for Page
	if raw.Page == 0 {
		params.Page = 1
	} else if raw.Page < 1 {
		return QueryParam{}, apperror.InvalidPageError()
	}

	return params, nil
}
