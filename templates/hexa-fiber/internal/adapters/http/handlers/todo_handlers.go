package handlers

import (
	"context"
	"strconv"
	"time"

	"hexagonal/internal/core/domain"
	ports "hexagonal/internal/core/ports"
	"hexagonal/pkg/helpers/filters"
	"hexagonal/pkg/helpers/pagination"
	"hexagonal/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type (
	ITodoHandler interface {
		HandleGetTodo(c *fiber.Ctx) error
		HandleGetTodoes(c *fiber.Ctx) error
		HandleUpdateTodo(c *fiber.Ctx) error
		HandleCreateTodo(c *fiber.Ctx) error
		HandleDeleteTodo(c *fiber.Ctx) error
	}
	TodoImpl struct {
		todoService ports.ITodoService
	}
)

func NewTodoHandler(
	todoService ports.ITodoService,
) ITodoHandler {
	return &TodoImpl{
		todoService: todoService,
	}
}

// HandleCreateTodo implements ITodoHandler.
func (h *TodoImpl) HandleCreateTodo(c *fiber.Ctx) error {
	var payload domain.TodoDomain
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.todoService.CreateTodo(ctx, payload)
	return c.JSON(res)
}

// HandleDeleteTodo implements ITodoHandler.
func (h *TodoImpl) HandleDeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.todoService.DeleteTodo(ctx, uint(id))
	return c.JSON(res)
}

// HandleUpdateTodo implements ITodoHandler.
func (h *TodoImpl) HandleUpdateTodo(c *fiber.Ctx) error {
	var payload domain.TodoDomain
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.todoService.UpdateTodo(ctx, uint(id), payload)
	return c.JSON(res)
}

// HandleGetTodo implements ITodoHandler.
func (h *TodoImpl) HandleGetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.todoService.GetTodo(ctx, uint(id))
	return c.JSON(res)
}

// HandleGetTodos implements ITodoHandler.
func (h *TodoImpl) HandleGetTodoes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	params := pagination.NewPaginationParams[filters.TodoFilter](c)
	paramCtx := pagination.SetFilters(ctx, params)
	res := h.todoService.GetTodos(paramCtx)
	return c.JSON(res)
}
