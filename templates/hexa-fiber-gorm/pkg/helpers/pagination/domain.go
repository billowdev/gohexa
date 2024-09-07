package pagination

import (
	"fmt"
	"strings"
)

// CommonTimeFilters contains common time-based filters such as date ranges and timestamps.
type CommonTimeFilters struct {
	DateField     string `json:"date_field"`
	CreatedAfter  string `json:"created_after"`
	UpdatedAfter  string `json:"updated_after"`
	CreatedBefore string `json:"created_before"`
	UpdatedBefore string `json:"updated_before"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}

// PaginationParams defines the parameters for pagination.
type PaginationParams[FilterType interface{}] struct {
	Limit      int        `json:"limit"`
	Page       int        `json:"page"`
	Sort       string     `json:"sort"`
	Order      string     `json:"order"`
	TotalRows  int64      `json:"total_rows"`
	TotalPages int        `json:"total_pages"`
	Filters    FilterType `json:"filters"`
	BaseURL    string     `json:"base_url"`
}

type SortParams struct {
	Sort           string // Field to sort by
	Order          string // Sorting order ("ASC" or "DESC")
	DefaultOrderBy string // Default
}

// NewOrderBy constructs the ORDER BY clause based on sorting parameters
func NewOrderBy(params SortParams) string {
	// DefaultOrderBy represents the default ORDER BY clause
	// Validate and sanitize sorting parameters
	sortField := strings.TrimSpace(params.Sort)
	sortOrder := strings.ToUpper(strings.TrimSpace(params.Order))

	// Construct the ORDER BY clause based on parameters
	if sortField != "" {
		if sortOrder == "ASC" || sortOrder == "DESC" {
			return fmt.Sprintf("%s %s", sortField, sortOrder)
		}
	}

	// If sorting parameters are invalid or not provided, fallback to default order
	return params.DefaultOrderBy
}

type PaginationResponseSwagger[T any] struct {
	StatusCode    string                `json:"status_code"`
	StatusMessage string                `json:"status_message"`
	Data          T                     `json:"data"`
	Pagination    PaginationInfoSwagger `json:"pagination"`
}
type PaginationInfoSwagger struct {
	Links      PaginationLinks `json:"links"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
}

// PaginationResponse represents the response structure for paginated data.
type PaginationResponse struct {
	StatusCode    string         `json:"status_code"`
	StatusMessage string         `json:"status_message"`
	Data          interface{}    `json:"data"`
	Pagination    PaginationInfo `json:"pagination"`
}

// PaginationInfo contains information about the pagination.
type PaginationInfo struct {
	Links      PaginationLinks `json:"links"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
	Rows       interface{}     `json:"rows,omitempty"`
}

type Pagination[T any] struct {
	Links      PaginationLinks `json:"links"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
	Rows       T               `json:"rows"`
}

// PaginationLinks contains links for next and previous pages.
type PaginationLinks struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type APIV2PaginationResponse struct {
	StatusCode    string      `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Data          interface{} `json:"data"`
	Pagination    interface{} `json:"pagination"`
}
