
package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/my_project/internal/adapters/database/models"
	ports "github.com/my_project/internal/core/ports/"
	"github.com/my_project/pkg/helpers/filters"
	"github.com/my_project/pkg/helpers/pagination"
	"github.com/my_project/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	IHandler interface {
		HandleGet(c *fiber.Ctx) error
		HandleGets(c *fiber.Ctx) error
		HandleUpdate(c *fiber.Ctx) error
		HandleCreate(c *fiber.Ctx) error
		HandleDelete(c *fiber.Ctx) error
	}
	Impl struct {
		Service ports.IService
	}
)

func NewHandler(
	Service ports.IService,
) IHandler {
	return &Impl{
		Service: Service,
	}
}

// HandleCreate implements IHandler.
func (h *Impl) HandleCreate(c *fiber.Ctx) error {
	var payload domain.Domain
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.Service.Create(ctx, payload)
	return c.JSON(res)
}

// HandleDelete implements IHandler.
func (h *Impl) HandleDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.Service.Delete(ctx, uint(id))
	return c.JSON(res)
}

// HandleUpdate implements IHandler.
func (h *Impl) HandleUpdate(c *fiber.Ctx) error {
	var payload domain.Domain
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
	res := h.Service.Update(ctx, uint(id), payload)
	return c.JSON(res)
}

// HandleGet implements IHandler.
func (h *Impl) HandleGet(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.Service.Get(ctx, uint(id))
	return c.JSON(res)
}

// HandleGets implements IHandler.
func (h *Impl) HandleGets(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	params := pagination.NewPaginationParams[filters.Filter](c)
	paramCtx := pagination.SetFilters(ctx, params)
	res := h.Service.Gets(paramCtx)
	return c.JSON(res)
}
