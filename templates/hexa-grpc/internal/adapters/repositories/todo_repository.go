// gohexa -generate repository -output ./internal/adapters/repositories -feature Todo -project hexagonal
package repositories

import (
	"context"

	"hexagonal/internal/adapters/database"
	"hexagonal/internal/adapters/database/models"
	ports "hexagonal/internal/core/ports"
	"hexagonal/pkg/helpers/filters"
	"hexagonal/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type TodoImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ports.ITodoRepository {
	return &TodoImpl{db: db}
}

// CreateTodo implements ports.ITodoRepository.
func (o *TodoImpl) CreateTodo(ctx context.Context, payload *models.Todo) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteTodo implements ports.ITodoRepository.
func (o *TodoImpl) DeleteTodo(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

// GetTodo implements ports.ITodoRepository.
func (o *TodoImpl) GetTodo(ctx context.Context, id uint) (*models.Todo, error) {
	tx := database.HelperExtractTx(ctx, o.db)

	var data models.Todo
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetTodos implements ports.ITodoRepository.
func (o *TodoImpl) GetTodos(ctx context.Context) (*pagination.Pagination[[]models.Todo], error) {
	tx := database.HelperExtractTx(ctx, o.db)

	p := pagination.GetFilters[filters.TodoFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.TodoFilter, []models.Todo](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateTodo implements ports.ITodoRepository.
func (o *TodoImpl) UpdateTodo(ctx context.Context, payload *models.Todo) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
