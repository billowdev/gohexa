package pagination

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"hexagonal/pkg/configs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewPaginationResponse[T any](c *fiber.Ctx, message string, data Pagination[[]T]) error {
	if data.Rows == nil || reflect.ValueOf(data.Rows).Len() == 0 {
		return c.Status(200).JSON(APIV2PaginationResponse{
			StatusCode:    configs.API_SUCCESS_CODE,
			StatusMessage: "The process of pagination was success",
			Data:          make([]interface{}, 0),
			Pagination: GetPaginationInfo(Pagination[interface{}]{
				Total:      0,
				Links:      PaginationLinks{},
				Page:       0,
				PageSize:   0,
				TotalPages: 0,
				Rows:       make([]interface{}, 0),
			}),
		})
	}
	response := APIV2PaginationResponse{
		StatusCode:    configs.API_SUCCESS_CODE,
		StatusMessage: message,
		Data:          data.Rows,
		Pagination:    GetPaginationInfo(data),
	}
	return c.Status(200).JSON(response)
}

func GetPaginationInfo[T any](payload Pagination[T]) PaginationInfo {
	return PaginationInfo{
		Links:      payload.Links,
		Total:      payload.Total,
		Page:       payload.Page,
		PageSize:   payload.PageSize,
		TotalPages: payload.TotalPages,
	}
}

// GetOffset calculates the offset for pagination.
func (p *PaginationParams[FilterType]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit returns the pagination limit with a default value if not set.
func (p *PaginationParams[FilterType]) GetLimit() int {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

// GetPage returns the current page with a default value if not set.
func (p *PaginationParams[FilterType]) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

// GetSort returns the sorting criteria with a default value if not set.
func (p *PaginationParams[FilterType]) GetSort() string {
	if p.Sort == "" {
		return "Id desc"
	}
	return p.Sort
}

func GetAPIEndpoint(c *fiber.Ctx) string {
	// Get the original URL from the request (excluding the query string)
	originalURL := c.OriginalURL()

	// Remove the query string (if any)
	if index := strings.Index(originalURL, "?"); index != -1 {
		originalURL = originalURL[:index]
	}

	// Construct the desired URL with the current host and path
	return fmt.Sprintf("https://%s%s", c.Hostname(), originalURL)
}

// Paginate is a function that returns a gorm.DB modifier for pagination.
func Paginate[FT any, T any](p PaginationParams[FT], query *gorm.DB) (Pagination[T], error) {
	var value T
	var totalRows int64
	if err := query.Model(value).Count(&totalRows).Error; err != nil {
		return Pagination[T]{}, err
	}
	p.TotalRows = totalRows
	p.TotalPages = int(math.Ceil(float64(totalRows) / float64(p.GetLimit())))

	if err := query.Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort()).Find(&value).Error; err != nil {
		return Pagination[T]{}, err

	}
	var nextLink, prevLink string
	if p.Page < p.TotalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page+1, p.Limit)
	}
	if p.Page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", p.BaseURL, p.Page-1, p.Limit)
	}
	return Pagination[T]{
		Links: PaginationLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      p.TotalRows,
		Page:       p.GetPage(),
		PageSize:   p.GetLimit(),
		TotalPages: p.TotalPages,
		Rows:       value,
	}, nil
}

func NewPaginationParams[FilterType interface{}](c *fiber.Ctx) PaginationParams[FilterType] {
	host := GetAPIEndpoint(c) // Assuming utils.GetAPIEndpoint is a function you have
	defaultLimit := 10
	defaultPage := 1
	defaultSort := "created_at desc"

	paginationParams := PaginationParams[FilterType]{
		Limit:   defaultLimit,
		Page:    defaultPage,
		Sort:    defaultSort,
		BaseURL: host,
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err == nil && limit > 0 {
		paginationParams.Limit = limit
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err == nil && page > 0 {
		paginationParams.Page = page
	}

	sort := c.Query("sort")
	if sort != "" {
		paginationParams.Sort = sort
	}

	order := c.Query("order")
	if sort != "" {
		paginationParams.Order = order
	}
	return paginationParams
}

func PaginateArray[T any](data []T, page, pageSize int, endpoint string) (PaginationInfo, []T) {
	totalItems := len(data)
	if pageSize <= 0 || page <= 0 {
		return PaginationInfo{}, nil // or return an error
	}

	totalPages := (totalItems + pageSize - 1) / pageSize

	if page > totalPages {
		return PaginationInfo{}, nil // or return an error
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > totalItems {
		end = totalItems
	}

	pageData := data[start:end]

	// Construct next and previous links
	var nextLink, prevLink string
	if page < totalPages {
		nextLink = fmt.Sprintf("%s?page=%d&page_size=%d", endpoint, page+1, pageSize)
	}
	if page > 1 {
		prevLink = fmt.Sprintf("%s?page=%d&page_size=%d", endpoint, page-1, pageSize)
	}

	pagination := PaginationInfo{
		Links: PaginationLinks{
			Next:     nextLink,
			Previous: prevLink,
		},
		Total:      int64(totalItems),
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
	return pagination, pageData
}
