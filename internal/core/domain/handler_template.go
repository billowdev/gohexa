package domain

type HandlerFlagDomain struct {
	FeatureName string
	ProjectName string
}

var HandlerTemplate = `
package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/{{ .ProjectName }}/internal/adapters/database/models"
	ports "{{ .ProjectName }}/internal/core/ports/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/pkg/helpers/filters"
	"github.com/{{ .ProjectName }}/pkg/helpers/pagination"
	"github.com/{{ .ProjectName }}/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	I{{ .FeatureName }}Handler interface {
		HandleGet{{ .FeatureName }}(c *fiber.Ctx) error
		HandleGet{{ .FeatureName }}s(c *fiber.Ctx) error
		HandleUpdate{{ .FeatureName }}(c *fiber.Ctx) error
		HandleCreate{{ .FeatureName }}(c *fiber.Ctx) error
		HandleDelete{{ .FeatureName }}(c *fiber.Ctx) error
	}
	{{ .FeatureName }}Impl struct {
		{{ .FeatureName | ToLower }}Service ports.I{{ .FeatureName }}Service
	}
)

func New{{ .FeatureName }}Handler(
	{{ .FeatureName | ToLower }}Service ports.I{{ .FeatureName }}Service,
) I{{ .FeatureName }}Handler {
	return &{{ .FeatureName }}Impl{
		{{ .FeatureName | ToLower }}Service: {{ .FeatureName | ToLower }}Service,
	}
}

// HandleCreate{{ .FeatureName }} implements I{{ .FeatureName }}Handler.
func (h *{{ .FeatureName }}Impl) HandleCreate{{ .FeatureName }}(c *fiber.Ctx) error {
	var payload models.{{ .FeatureName }}
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.{{ .FeatureName | ToLower }}Service.Create{{ .FeatureName }}(ctx, &payload)
	return c.JSON(res)
}

// HandleDelete{{ .FeatureName }} implements I{{ .FeatureName }}Handler.
func (h *{{ .FeatureName }}Impl) HandleDelete{{ .FeatureName }}(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.{{ .FeatureName | ToLower }}Service.Delete{{ .FeatureName }}(ctx, uint(id))
	return c.JSON(res)
}

// HandleUpdate{{ .FeatureName }} implements I{{ .FeatureName }}Handler.
func (h *{{ .FeatureName }}Impl) HandleUpdate{{ .FeatureName }}(c *fiber.Ctx) error {
	var payload models.{{ .FeatureName }}
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
	res := h.{{ .FeatureName | ToLower }}Service.Update{{ .FeatureName }}(ctx, uint(id), &payload)
	return c.JSON(res)
}

// HandleGet{{ .FeatureName }} implements I{{ .FeatureName }}Handler.
func (h *{{ .FeatureName }}Impl) HandleGet{{ .FeatureName }}(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.{{ .FeatureName | ToLower }}Service.Get{{ .FeatureName }}(ctx, uint(id))
	return c.JSON(res)
}

// HandleGet{{ .FeatureName }}s implements I{{ .FeatureName }}Handler.
func (h *{{ .FeatureName }}Impl) HandleGet{{ .FeatureName }}s(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	params := pagination.NewPaginationParams[filters.{{ .FeatureName }}Filter](c)
	paramCtx := pagination.SetFilters(ctx, params)
	res := h.{{ .FeatureName | ToLower }}Service.Get{{ .FeatureName }}s(paramCtx)
	return c.JSON(res)
}
`
