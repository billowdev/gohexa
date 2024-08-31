## Repositories Generator

### Overview
The Repositories Generator tool creates Go implementation files for repositories in a specified feature. Repositories handle interactions with the database and provide methods for CRUD operations on your models. This tool generates these implementations based on the provided feature name and project details.

### Flags and Parameters
- `-feature <FeatureName>`: The name of the feature for which to generate a repository (e.g., Order).
- `-output <OutputDirectory>`: The directory where the generated repository files will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).

### Template Content
- The template generates a Go file with an implementation of the repository interface for the specified feature.
- Includes methods for:
	- Create<FeatureName>: Create a new record.
	- Delete<FeatureName>: Delete a record by ID.
	- Get<FeatureName>: Get a record by ID.
	- Get<FeatureName>s: Get a list of records with pagination and filters.
	- Update<FeatureName>: Update an existing record.

### Command
To generate a repository file, use the following command:
```bash
go run github.com/rapidstellar/gohexa/generators/repositories -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
Generate Repository File:
```bash
go run github.com/rapidstellar/gohexa/generators/repositories -feature="Order" -output ./internal/adapters/repositories -project my_project
```
This command generates an `order_repository.go` file in the `./internal/adapters/repositories` directory.

### Template Example
Here is a sample of the generated repository file based on the template:
```go
package repositories

import (
	"context"

	"github.com/my_project/internal/adapters/database"
	"github.com/my_project/internal/adapters/database/models"
	ports "github.com/my_project/internal/core/ports/order"
	"github.com/my_project/pkg/helpers/filters"
	"github.com/my_project/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type OrderImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.IOrderRepository {
	return &OrderImpl{db: db}
}

// CreateOrder implements ports.IOrderRepository.
func (o *OrderImpl) CreateOrder(ctx context.Context, payload *models.Order) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder implements ports.IOrderRepository.
func (o *OrderImpl) DeleteOrder(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}

// GetOrder implements ports.IOrderRepository.
func (o *OrderImpl) GetOrder(ctx context.Context, id uint) (*models.Order, error) {
	tx := database.HelperExtractTx(ctx, o.db)

	var data models.Order
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrders implements ports.IOrderRepository.
func (o *OrderImpl) GetOrders(ctx context.Context) (*pagination.Pagination[[]models.Order], error) {
	tx := database.HelperExtractTx(ctx, o.db)

	p := pagination.GetFilters[filters.OrderFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.OrderFilter, []models.Order](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateOrder implements ports.IOrderRepository.
func (o *OrderImpl) UpdateOrder(ctx context.Context, payload *models.Order) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
```

### Repositories Generators Usage Notes
- Ensure that the output directory exists or is created by the tool.
- Adjust the featureName to match your domain model naming conventions.