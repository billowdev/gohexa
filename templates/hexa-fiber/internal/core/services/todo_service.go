package services

import (
	"context"

	"hexagonal/internal/adapters/database"
	domain "hexagonal/internal/core/domain"
	ports "hexagonal/internal/core/ports"
	"hexagonal/pkg/configs"
	"hexagonal/pkg/helpers/pagination"
	"hexagonal/pkg/utils"
)

type TodoServiceImpl struct {
	repo       ports.ITodoRepository
	transactor database.IDatabaseTransactor
}

func NewTodoService(
	repo ports.ITodoRepository,
	transactor database.IDatabaseTransactor,
) ports.ITodoService {
	return &TodoServiceImpl{repo: repo, transactor: transactor}
}

// CreateTodo implements ports.ITodoService.
func (s *TodoServiceImpl) CreateTodo(ctx context.Context, payload domain.TodoDomain) utils.APIResponse {
	data := domain.ToTodoModel(payload)
	if err := s.repo.CreateTodo(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// DeleteTodo implements ports.ITodoService.
func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, id uint) utils.APIResponse {
	if err := s.repo.DeleteTodo(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// GetTodo implements ports.ITodoService.
func (s *TodoServiceImpl) GetTodo(ctx context.Context, id uint) utils.APIResponse {
	data, err := s.repo.GetTodo(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	res := domain.ToTodoDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// GetTodos implements ports.ITodoService.
func (s *TodoServiceImpl) GetTodos(ctx context.Context) pagination.Pagination[[]domain.TodoDomain] {
	data, err := s.repo.GetTodos(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.TodoDomain]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.ToTodoDomain)
	return pagination.Pagination[[]domain.TodoDomain]{
		Rows:       newData,
		Links:      data.Links,
		Total:      data.Total,
		Page:       data.Page,
		PageSize:   data.PageSize,
		TotalPages: data.TotalPages,
	}
}

// UpdateTodo implements ports.ITodoService.
func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, id uint, payload domain.TodoDomain) utils.APIResponse {
	if _, err := s.repo.GetTodo(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	payload.ID = id
	data := domain.ToTodoModel(payload)
	if err := s.repo.UpdateTodo(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.ToTodoDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}
