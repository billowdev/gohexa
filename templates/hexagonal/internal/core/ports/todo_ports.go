
package ports

import (
	"context"

	"hexagonal/internal/adapters/database/models"
	domain "hexagonal/internal/core/domain"
	"hexagonal/pkg/helpers/pagination"
	"hexagonal/pkg/utils"
)

type ITodoRepository interface {
	GetTodo(ctx context.Context, id uint) (*models.Todo, error)
	GetTodos(ctx context.Context) (*pagination.Pagination[[]models.Todo], error)
	CreateTodo(ctx context.Context, payload *models.Todo) error
	UpdateTodo(ctx context.Context, payload *models.Todo) error
	DeleteTodo(ctx context.Context, id uint) error
}

type ITodoService interface {
	GetTodo(ctx context.Context, id uint) utils.APIResponse
	GetTodos(ctx context.Context) pagination.Pagination[[]domain.TodoDomain]
	CreateTodo(ctx context.Context, payload domain.TodoDomain) utils.APIResponse
	UpdateTodo(ctx context.Context, id uint, payload domain.TodoDomain) utils.APIResponse
	DeleteTodo(ctx context.Context, id uint) utils.APIResponse
}
