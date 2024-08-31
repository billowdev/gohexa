## Handlers Generator

### Overview
This tool generates Go handler files for a specified feature. The handlers include CRUD operations (Create, Read, Update, Delete) for interacting with the service layer. The generated handlers are compatible with the Fiber framework and are designed to work with the service interfaces and models provided in the project.

### Flags and Parameters
- `-feature <FeatureName>`: The name of the feature for which to generate the handlers (e.g., Order).
- `-output <OutputDirectory>`: The directory where the generated handler file will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).

### Template Content
The template generates a Go file with handlers for CRUD operations:
- HandleCreate{{ .FeatureName }}
- HandleDelete{{ .FeatureName }}
- HandleUpdate{{ .FeatureName }}
- HandleGet{{ .FeatureName }}
- HandleGet{{ .FeatureName }}s
- Handlers are implemented as methods on a struct that implements the I{{ .FeatureName }}Handler interface.
- Error handling is included for invalid requests and responses.
- Utility functions for creating and deleting handlers are used, including context timeout and error handling.

### Command
To generate a handler file, use the following command:
```bash
go run github.com/rapidstellar/gohexa/generators/handlers -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```
### Example Commands
1. Generate Handlers File:

```bash
go run github.com/rapidstellar/gohexa/generators/handlers -feature="Todo" -output ./internal/adapters/handlers -project my_project
```
This command generates a todo_handlers.go file in the ./internal/adapters/handlers directory.

### Template Example
Here is a sample of the generated handlers file based on the template:
```go
package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/my_project/internal/adapters/database/models"
	ports "github.com/my_project/internal/core/ports/todo"
	"github.com/my_project/pkg/helpers/filters"
	"github.com/my_project/pkg/helpers/pagination"
	"github.com/my_project/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	ITodoHandler interface {
		HandleGetTodo(c *fiber.Ctx) error
		HandleGetTodos(c *fiber.Ctx) error
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
	var payload models.Todo
	if err := c.BodyParser(&payload); err != nil {
		return utils.NewErrorResponse(c, "Invalid request payload", err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ctx.Err(); err != nil {
		return c.Context().Err()
	}
	res := h.todoService.CreateTodo(ctx, &payload)
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
	var payload models.Todo
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
	res := h.todoService.UpdateTodo(ctx, uint(id), &payload)
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
func (h *TodoImpl) HandleGetTodos(c *fiber.Ctx) error {
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
```


