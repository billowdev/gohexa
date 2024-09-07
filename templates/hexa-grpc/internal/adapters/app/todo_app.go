
package app

import (
	"hexagonal/internal/adapters/database"
	handlers "hexagonal/internal/adapters/http/handlers"
	"hexagonal/internal/adapters/http/routers"
	repositories "hexagonal/internal/adapters/repositories"
	services "hexagonal/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppContainer(app *fiber.App, db *gorm.DB) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	TodoApp(route, db)
	return app
}

func TodoApp(r routers.RouterImpl, db *gorm.DB) {
	transactorRepo := database.NewTransactorRepo(db)
	todoRepo := repositories.NewTodoRepository(db)
	todoSrv := services.NewTodoService(todoRepo, transactorRepo)
	todoHandlers := handlers.NewTodoHandler(todoSrv)
	r.CreateTodoRoutes(todoHandlers)
}
