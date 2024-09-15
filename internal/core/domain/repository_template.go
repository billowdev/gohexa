package domain

type RepositoryFlagDomain struct {
	FeatureName string
	ProjectName string
}

var RepoTemplate = `
package repositories

import (
	"context"

	"github.com/{{ .ProjectName }}/internal/adapters/database"
	"github.com/{{ .ProjectName }}/internal/adapters/database/models"
	ports "github.com/{{ .ProjectName }}/internal/core/ports/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/pkg/helpers/filters"
	"github.com/{{ .ProjectName }}/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type {{ .FeatureName }}Impl struct {
	db *gorm.DB
}

func New{{ .FeatureName }}Repository(db *gorm.DB) ports.I{{ .FeatureName }}Repository {
	return &{{ .FeatureName }}Impl{db: db}
}

// BulkCreate{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
// It accepts a slice of {{ .FeatureName }} and inserts them into the database in bulk.
func (o *{{ .FeatureName }}Impl) BulkCreate{{ .FeatureName }}(ctx context.Context, payloads []*models.{{ .FeatureName }}) error {
	tx := database.HelperExtractTx(ctx, o.db)

	// Start a new transaction context
	tx = tx.WithContext(ctx)

	// Perform bulk creation
	if err := tx.Create(&payloads).Error; err != nil {
		return err
	}

	return nil
}

// GetOneByFields implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) GetOneByFields(ctx context.Context, filters map[string]interface{}) (*models.{{ .FeatureName }}, error) {
	tx := c.db.WithContext(ctx)
	var data models.{{ .FeatureName }}
	// Build the query dynamically based on the filters using the utility function
	query, args := helpers.BuildQueryConditions(filters)

	// Execute the query
	if err := tx.Where(query, args...).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// Create{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Create{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// Delete{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Delete{{ .FeatureName }}(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.{{ .FeatureName }}{}).Error; err != nil {
		return err
	}
	return nil
}

// Get{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Get{{ .FeatureName }}(ctx context.Context, id uint) (*models.{{ .FeatureName }}, error) {
	tx := database.HelperExtractTx(ctx, o.db)

	var data models.{{ .FeatureName }}
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// Get{{ .FeatureName }}s implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Get{{ .FeatureName }}s(ctx context.Context) (*pagination.Pagination[[]models.{{ .FeatureName }}], error) {
	tx := database.HelperExtractTx(ctx, o.db)

	p := pagination.GetFilters[filters.{{ .FeatureName }}Filter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.{{ .FeatureName }}Filter, []models.{{ .FeatureName }}](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Update{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Update{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
`
