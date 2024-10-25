package util

import (
	"strconv"
	"strings"
	"time"

	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
)

type GetProductQuery struct {
	SortBy    *string `form:"sort-by" validate:"omitempty"`
	Search    *string `form:"search" validate:"omitempty"`
	Sort      *string `form:"sort" validate:"omitempty"`
	Latitude  *string `form:"lat" validate:"required,latitude"`
	Longitude *string `form:"long" validate:"required,longitude"`
	Category  *string `form:"category" validate:"omitempty,number"`
	MinPrice  *string `form:"min-price" validate:"omitempty,number"`
	MaxPrice  *string `form:"max-price" validate:"omitempty,number"`
	Limit     *string `form:"limit" validate:"omitempty,number"`
	Page      *string `form:"page" validate:"omitempty,number"`
}

type ValidatedGetProductQuery struct {
	SortBy    *string
	Search    *string
	Sort      *string
	Latitude  string
	Longitude string
	Category  *int64
	MinPrice  *int
	MaxPrice  *int
	Limit     int
	Page      int
}

func ValidateGetProductQuery(query GetProductQuery) (*ValidatedGetProductQuery, error) {
	resQuery := ValidatedGetProductQuery{}

	if query.SortBy != nil {
		if strings.ToLower(*query.SortBy) != "price" {
			return nil, apperror.InvalidSortByError()
		}
		resQuery.SortBy = query.SortBy
	}

	if query.Sort != nil {
		if strings.ToLower(*query.Sort) != "asc" && strings.ToLower(*query.Sort) != "desc" {
			return nil, apperror.InvalidSortError()
		}
		resQuery.Sort = query.Sort
	}

	if (query.SortBy != nil && query.Sort == nil) || (query.SortBy == nil && query.Sort != nil) {
		return nil, apperror.InvalidSortPairError()
	}

	if query.Search != nil {
		resQuery.Search = query.Search
	}

	resQuery.Latitude = *query.Latitude
	resQuery.Longitude = *query.Longitude

	if query.Category != nil {
		categoryInt, err := strconv.Atoi(*query.Category)
		categoryInt64 := int64(categoryInt)
		if err != nil {
			return nil, apperror.InvalidCategoryError()
		}
		if categoryInt < 1 {
			return nil, apperror.InvalidCategoryError()
		}
		resQuery.Category = &categoryInt64
	}

	if query.MinPrice != nil {
		minPriceInt, err := strconv.Atoi(*query.MinPrice)
		if err != nil {
			return nil, apperror.InvalidMinPriceError()
		}
		if minPriceInt < 0 {
			return nil, apperror.InvalidMinPriceError()
		}
		resQuery.MinPrice = &minPriceInt
	}

	if query.MaxPrice != nil {
		maxPriceInt, err := strconv.Atoi(*query.MaxPrice)
		if err != nil {
			return nil, apperror.InvalidMaxPriceError()
		}
		if maxPriceInt < 0 {
			return nil, apperror.InvalidMaxPriceError()
		}
		resQuery.MaxPrice = &maxPriceInt
	}

	if resQuery.MinPrice != nil && resQuery.MaxPrice != nil {
		if *resQuery.MaxPrice < *resQuery.MinPrice {
			return nil, apperror.InvalidPriceRangeError()
		}
	}

	if query.Limit == nil {
		resQuery.Limit = 12
	}
	if query.Limit != nil {
		limitInt, err := strconv.Atoi(*query.Limit)
		if err != nil {
			return nil, apperror.InvalidLimitError()
		}
		if limitInt < 1 {
			return nil, apperror.InvalidLimitError()
		}
		resQuery.Limit = limitInt
	}

	if query.Page == nil {
		resQuery.Page = 1
	}
	if query.Page != nil {
		pageInt, err := strconv.Atoi(*query.Page)
		if err != nil {
			return nil, apperror.InvalidPageError()
		}
		if pageInt < 1 {
			return nil, apperror.InvalidPageError()
		}
		resQuery.Page = pageInt
	}

	return &resQuery, nil
}

type GetDrugsAdminQuery struct {
	Search *string `form:"search" validate:"omitempty"`
	Limit  *string `form:"limit" validate:"omitempty,number"`
	Page   *string `form:"page" validate:"omitempty,number"`
}

type ValidatedGetDrugAdminQuery struct {
	Search *string
	Limit  int
	Page   int
}

func ValidateGetDrugAdminQuery(query GetDrugsAdminQuery) (*ValidatedGetDrugAdminQuery, error) {
	validatedQuery := ValidatedGetDrugAdminQuery{}

	if query.Search != nil {
		validatedQuery.Search = query.Search
	}

	if query.Limit == nil {
		validatedQuery.Limit = 12
	}
	if query.Limit != nil {
		limitInt, err := strconv.Atoi(*query.Limit)
		if err != nil {
			return nil, apperror.InvalidLimitError()
		}
		if limitInt < 1 {
			return nil, apperror.InvalidLimitError()
		}
		validatedQuery.Limit = limitInt
	}

	if query.Page == nil {
		validatedQuery.Page = 1
	}
	if query.Page != nil {
		pageInt, err := strconv.Atoi(*query.Page)
		if err != nil {
			return nil, apperror.InvalidPageError()
		}
		if pageInt < 1 {
			return nil, apperror.InvalidPageError()
		}
		validatedQuery.Page = pageInt
	}

	return &validatedQuery, nil
}

type GetOrderQuery struct {
	StatusId     *int64  `form:"status_id" validate:"omitempty"`
	PharmacyName *string `form:"pharmacy_name" validate:"omitempty"`
	Limit        *string `form:"limit" validate:"omitempty,number"`
	Page         *string `form:"page" validate:"omitempty,number"`
}

type ValidatedGetOrderQuery struct {
	StatusId     *int64
	PharmacyName *string
	Limit        int
	Page         int
}

func ValidateGetOrderQuery(query GetOrderQuery) (*ValidatedGetOrderQuery, error) {
	validatedGetOrderQuery := ValidatedGetOrderQuery{}

	if query.StatusId != nil {
		if *query.StatusId > appconstant.OrderStatusCanceled || *query.StatusId < appconstant.OrderStatusWaitingForPayment {
			return nil, apperror.InvalidOrderStatusError()
		}
		validatedGetOrderQuery.StatusId = query.StatusId
	}

	if query.PharmacyName != nil {
		validatedGetOrderQuery.PharmacyName = query.PharmacyName
	}

	if query.Limit == nil {
		validatedGetOrderQuery.Limit = 10
	}
	if query.Limit != nil {
		limitInt, err := strconv.Atoi(*query.Limit)
		if err != nil {
			return nil, apperror.InvalidLimitError()
		}
		if limitInt < 1 {
			return nil, apperror.InvalidLimitError()
		}
		validatedGetOrderQuery.Limit = limitInt
	}

	if query.Page == nil {
		validatedGetOrderQuery.Page = 1
	}
	if query.Page != nil {
		pageInt, err := strconv.Atoi(*query.Page)
		if err != nil {
			return nil, apperror.InvalidPageError()
		}
		if pageInt < 1 {
			return nil, apperror.InvalidPageError()
		}
		validatedGetOrderQuery.Page = pageInt
	}

	return &validatedGetOrderQuery, nil
}

type GetReportQuery struct {
	PharmacyId int64   `form:"pharmacy_id" validate:"required"`
	Sort       *string `form:"sort"`
	MaxDate    string  `form:"max_date" validate:"required,datetime=2006-01-02"`
	MinDate    string  `form:"min_date" validate:"required,datetime=2006-01-02"`
}

type ValidatedGetReportQuery struct {
	PharmacyId int64
	Sort       *string
	MaxDate    string
	MinDate    string
}

func ValidateGetReportQuery(query GetReportQuery) (*ValidatedGetReportQuery, error) {
	validatedGetReportQuery := ValidatedGetReportQuery{}

	if query.PharmacyId > 0 {
		validatedGetReportQuery.PharmacyId = query.PharmacyId
	}

	if query.Sort != nil {
		if strings.ToLower(*query.Sort) != "asc" && strings.ToLower(*query.Sort) != "desc" {
			return nil, apperror.InvalidSortError()
		}
		validatedGetReportQuery.Sort = query.Sort
	}

	maxDate, err := time.Parse("2006-01-02", query.MaxDate)
	if err != nil {
		return nil, apperror.InvalidMaxDateError()
	}

	minDate, err := time.Parse("2006-01-02", query.MinDate)
	if err != nil {
		return nil, apperror.InvalidMinDateError()
	}

	if minDate.After(maxDate) {
		return nil, apperror.InvalidMinDateError()
	}

	maxDate = maxDate.UTC().Add(time.Hour*time.Duration(23) +
		time.Minute*time.Duration(59) +
		time.Second*time.Duration(59))

	maxDateString := maxDate.Format("2006-01-02 15:04:05")
	validatedGetReportQuery.MaxDate = maxDateString

	minDateString := minDate.Format("2006-01-02 15:04:05")
	validatedGetReportQuery.MinDate = minDateString

	return &validatedGetReportQuery, nil
}
