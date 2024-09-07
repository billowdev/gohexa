
package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/OrderTest/internal/adapters/database/models"
	ports "github.com/OrderTest/internal/core/ports/document"
	"github.com/OrderTest/pkg/helpers/filters"
	"github.com/OrderTest/pkg/helpers/pagination"
	"github.com/OrderTest/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	IDocumentHandler interface {
		HandleGetDocument(c *fiber.Ctx) error
		HandleGetDocuments(c *fiber.Ctx) error
		HandleUpdateDocument(c *fiber.Ctx) error
		HandleCreateDocument(c *fiber.Ctx) error
		HandleDeleteDocument(c *fiber.Ctx) error
	}
	DocumentImpl struct {
		documentService ports.IDocumentService
	}
)

func NewDocumentHandler(
	documentService ports.IDocumentService,
) IDocumentHandler {
	return &DocumentImpl{
		documentService: documentService,
	}
}

// HandleCreateDocument implements IDocumentHandler.
func (h *DocumentImpl) HandleCreateDocument(c *fiber.Ctx) error {
	var payload domain.DocumentDomain
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.documentService.CreateDocument(ctx, payload)
	return c.JSON(res)
}

// HandleDeleteDocument implements IDocumentHandler.
func (h *DocumentImpl) HandleDeleteDocument(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.documentService.DeleteDocument(ctx, uint(id))
	return c.JSON(res)
}

// HandleUpdateDocument implements IDocumentHandler.
func (h *DocumentImpl) HandleUpdateDocument(c *fiber.Ctx) error {
	var payload domain.DocumentDomain
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
	res := h.documentService.UpdateDocument(ctx, uint(id), payload)
	return c.JSON(res)
}

// HandleGetDocument implements IDocumentHandler.
func (h *DocumentImpl) HandleGetDocument(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.NewErrorResponse(c, "Invalid ID", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.documentService.GetDocument(ctx, uint(id))
	return c.JSON(res)
}

// HandleGetDocuments implements IDocumentHandler.
func (h *DocumentImpl) HandleGetDocuments(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	params := pagination.NewPaginationParams[filters.DocumentFilter](c)
	paramCtx := pagination.SetFilters(ctx, params)
	res := h.documentService.GetDocuments(paramCtx)
	return c.JSON(res)
}
