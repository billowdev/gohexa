package pagination

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type filterKey struct{}

// SetFilters sets filters in the context
func SetFilters[T interface{}](ctx context.Context, filters PaginationParams[T]) context.Context {
	return context.WithValue(ctx, filterKey{}, filters)
}

// GetFilters retrieves filters from the context
func GetFilters[T interface{}](ctx context.Context) PaginationParams[T] {
	if f, ok := ctx.Value(filterKey{}).(PaginationParams[T]); ok {
		return f
	}
	return PaginationParams[T]{}
}

// AddWhereClauseIfNotEmpty adds a WHERE clause to the query if the filter value is not empty,
// based on the provided column name, filter value, and filter type.
//
// Parameters:
//   - query: A pointer to the gorm.DB instance.
//   - columnName: The name of the column to filter.
//   - filterValue: The value to filter by.
//   - filterType: The type of filtering to apply (exact, like, date).
//
// Returns:
//   - *gorm.DB: A pointer to the modified gorm.DB instance.
func AddWhereClauseIfNotEmpty(query *gorm.DB, columnName string, filterValue string, filterType string) *gorm.DB {
	if filterValue != "" {
		switch filterType {
		case "exact":
			return query.Where(fmt.Sprintf("%s = ?", columnName), filterValue)
		case "like":
			return query.Where(fmt.Sprintf("LOWER(%s) LIKE LOWER(?)", columnName), "%"+strings.ToLower(filterValue)+"%")
		case "date":
			query = query.Where(fmt.Sprintf("DATE(%s) = ?", columnName), filterValue)
		default:
			return query
		}
	}
	return query
}

// ApplyFilter applies filtering to the query based on the provided column, value, and filter type.
//
// Parameters:
//   - query: A pointer to the gorm.DB instance.
//   - column: The name of the column to filter.
//   - value: The value to filter by.
//   - filterType: The type of filtering to apply (contains, exact).
//
// Returns:
//   - *gorm.DB: A pointer to the modified gorm.DB instance.
func ApplyFilter(query *gorm.DB, column string, value interface{}, filterType string) *gorm.DB {
	if filterType == "" {
		filterType = "contains"
	}
	switch filterType {
	case "contains":
		return AddWhereClauseIfNotEmpty(query, column, value.(string), "like")
	case "exact":
		return AddWhereClauseIfNotEmpty(query, column, value.(string), "exact")
	default:
		return query
	}
}

// ApplyCommaFilter applies filtering to the query for comma-separated values in the specified column.
//
// Parameters:
//   - query: A pointer to the gorm.DB instance.
//   - columnName: The name of the column to filter.
//   - filterValue: The comma-separated values to filter by.
//
// Returns:
//   - *gorm.DB: A pointer to the modified gorm.DB instance.
func ApplyCommaFilter(query *gorm.DB, columnName, filterValue string) *gorm.DB {
	if filterValue != "" {
		filterVals := strings.Split(filterValue, ",")
		if len(filterVals) != 0 {
			// Trim spaces from each element in the slice
			for i, val := range filterVals {
				filterVals[i] = strings.TrimSpace(val)
			}
			if filterVals[0] != "" {
				query = query.Where(columnName+" IN (?)", filterVals)
			} else {
				query = query.Where(columnName+" LIKE ?", "%"+filterValue+"%")
			}
		} else {
			query = query.Where(columnName+" LIKE ?", "%"+filterValue+"%")
		}
	}
	return query
}

// ApplyCommaFilterWithJoin applies filtering to the query with a JOIN condition for comma-separated values.
//
// Parameters:
//   - query: A pointer to the gorm.DB instance.
//   - joinTable: The table to join.
//   - joinCondition: The condition for joining.
//   - columnName: The name of the column to filter.
//   - filterValue: The comma-separated values to filter by.
//
// Returns:
//   - *gorm.DB: A pointer to the modified gorm.DB instance.
func ApplyCommaFilterWithJoin(query *gorm.DB, joinTable, joinCondition, columnName, filterValue string) *gorm.DB {
	if filterValue != "" {
		filterVals := strings.Split(filterValue, ",")
		if len(filterVals) != 0 {
			// Trim spaces from each element in the slice
			for i, val := range filterVals {
				filterVals[i] = strings.TrimSpace(val)
			}
			if filterVals[0] != "" {
				// Use raw SQL to apply filter with join condition
				query = query.Joins(joinTable).Where(joinCondition+" IN (?)", filterVals)
			} else {
				query = query.Joins(joinTable).Where(joinCondition+" LIKE ?", "%"+filterValue+"%")
			}
		} else {
			query = query.Joins(joinTable).Where(joinCondition+" LIKE ?", "%"+filterValue+"%")
		}
	}
	return query
}

// filterDateRange applies date range filtering to the query based on the provided field and CommonTimeFilters.
//
// Parameters:
//   - field: The name of the field to filter.
//   - filter: The CommonTimeFilters containing filtering parameters.
//   - query: A pointer to the gorm.DB instance.
//
// Returns:
//   - *gorm.DB: A pointer to the modified gorm.DB instance.
func filterDateRange(field string, filter CommonTimeFilters, query *gorm.DB) *gorm.DB {
	if filter.StartDate != "" {
		query = query.Where(fmt.Sprintf("%s >= ?", field), filter.StartDate)
	}
	if filter.EndDate != "" {
		query = query.Where(fmt.Sprintf("%s <= ?", field), filter.EndDate)
	}
	return query
}

// ApplyDatetimeFilters applies datetime filtering to the query based on the provided CommonTimeFilters.
//
// Parameters:
//   - query: A pointer to the gorm.DB instance.
//   - filter: The CommonTimeFilters containing filtering parameters.
//
// Returns:
//   - *gorm.DB: A pointer to the modified gorm.DB instance.
func ApplyDatetimeFilters(query *gorm.DB, filter CommonTimeFilters) *gorm.DB {
	switch filter.DateField {
	case "created_at":
		query = filterDateRange("created_at", filter, query)
	case "updated_at":
		query = filterDateRange("updated_at", filter, query)
	default:
		query = filterDateRange("created_at", filter, query)
	}

	if filter.CreatedAfter != "" {
		query = query.Where(fmt.Sprintf("%s %s ?", "created_at", ">="), filter.CreatedAfter)
	}
	if filter.UpdatedAfter != "" {
		query = query.Where(fmt.Sprintf("%s %s ?", "updated_at", ">="), filter.UpdatedAfter)
	}

	if filter.CreatedBefore != "" {
		query = query.Where(fmt.Sprintf("%s %s ?", "created_at", "<="), filter.CreatedBefore)
	}
	if filter.UpdatedBefore != "" {
		query = query.Where(fmt.Sprintf("%s %s ?", "updated_at", "<="), filter.UpdatedBefore)
	}

	if filter.CreatedAt != "" {
		// PostgreSQL
		query = query.Where(fmt.Sprintf("DATE(%s) = ?", "created_at"), filter.CreatedAt)
	}
	if filter.UpdatedAt != "" {
		//  PostgreSQL
		query = query.Where(fmt.Sprintf("DATE(%s) = ?", "updated_at"), filter.UpdatedAt)
	}

	return query
}

// ApplyDatetimePreloadFilters applies datetime filtering with preloading to the query based on the provided CommonTimeFilters.
//
// Parameters:
//   - query: A pointer to the gorm.DB instance.
//   - filter: The CommonTimeFilters containing filtering parameters.
//   - preloadKey: The key for preloading.
//
// Returns:
//   - *gorm.DB: A pointer to the modified gorm.DB instance.
func ApplyDatetimePreloadFilters(query *gorm.DB, filter CommonTimeFilters, preloadKey string) *gorm.DB {
	switch filter.DateField {
	case "created_at":
		query = filterDateRange(fmt.Sprintf("%s.created_at", preloadKey), filter, query)
	case "updated_at":
		query = filterDateRange(fmt.Sprintf("%s.updated_at", preloadKey), filter, query)
	default:
		query = filterDateRange(fmt.Sprintf("%s.created_at", preloadKey), filter, query)
	}

	// Apply >= filters
	if filter.CreatedAfter != "" {
		query = query.Where(fmt.Sprintf("%s >= ?", fmt.Sprintf("%s.created_at", preloadKey)), filter.CreatedAfter)
	}
	if filter.UpdatedAfter != "" {
		query = query.Where(fmt.Sprintf("%s >= ?", fmt.Sprintf("%s.updated_at", preloadKey)), filter.UpdatedAfter)
	}

	// Apply <= filters
	if filter.CreatedBefore != "" {
		query = query.Where(fmt.Sprintf("%s <= ?", fmt.Sprintf("%s.created_at", preloadKey)), filter.CreatedBefore)
	}
	if filter.UpdatedBefore != "" {
		query = query.Where(fmt.Sprintf("%s <= ?", fmt.Sprintf("%s.updated_at", preloadKey)), filter.UpdatedBefore)
	}

	// Apply DATE filters (for PostgreSQL)
	if filter.CreatedAt != "" {
		query = query.Where(fmt.Sprintf("DATE(%s) = ?", fmt.Sprintf("%s.created_at", preloadKey)), filter.CreatedAt)
	}
	if filter.UpdatedAt != "" {
		query = query.Where(fmt.Sprintf("DATE(%s) = ?", fmt.Sprintf("%s.updated_at", preloadKey)), filter.UpdatedAt)
	}

	return query
}
