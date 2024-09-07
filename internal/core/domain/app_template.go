package domain

type AppFlagDomain struct {
	FeatureName string
	ProjectName string
}

var AppTemplate = `
package app

import (
	"github.com/{{ .ProjectName }}/internal/adapters/database"
	handlers "github.com/{{ .ProjectName }}/internal/adapters/http/handlers/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/internal/adapters/http/routers"
	repositories "github.com/{{ .ProjectName }}/internal/adapters/repositories/{{ .FeatureName | ToLower }}"
	services "github.com/{{ .ProjectName }}/internal/core/services/{{ .FeatureName | ToLower }}"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppContainer(app *fiber.App, db *gorm.DB) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	{{ .FeatureName }}App(route, db)
	return app
}

func {{ .FeatureName }}App(r routers.RouterImpl, db *gorm.DB) {
	transactorRepo := database.NewTransactorRepo(db)
	{{ .FeatureName | ToLower }}Repo := repositories.New{{ .FeatureName }}Repo(db)
	{{ .FeatureName | ToLower }}Srv := services.New{{ .FeatureName }}Service({{ .FeatureName | ToLower }}Repo, transactorRepo)
	{{ .FeatureName | ToLower }}Handlers := handlers.New{{ .FeatureName }}Handler({{ .FeatureName | ToLower }}Srv)
	r.Create{{ .FeatureName }}Route({{ .FeatureName | ToLower }}Handlers)
}
`
