## App File Generator

### Overview
The App File Generator is a command-line tool that generates an application setup file for a specific feature within a Go project. The generated file integrates the feature's services, handlers, and routes, following a modular architecture.

### Flags and Parameters:

- `-feature <FeatureName>`: The name of the feature (e.g., SystemField).
- `-output <OutputDirectory>`: The directory where the generated file will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).

### Template Content:

- The template includes placeholders for the feature name and project name.
- {{ .FeatureName | ToLower }} is used to convert the feature name to lowercase for consistency in naming conventions.
- The AppContainer function initializes the Fiber app and sets up routes and services.
- The {{ .FeatureName }}App function configures the feature-specific components.

### Command

```bash
go run ./generators/app -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```
- example
```bash
go run ./generators/app -feature="Todo" -output ./internal/adapters/app -project my_project
```

### Generated File Structure
The tool generates a Go file in the specified output directory, following this structure:
```go
package app

import (
	"github.com/<ProjectName>/internal/adapters/database"
	handlers "github.com/<ProjectName>/internal/adapters/http/handlers/<featureName>"
	"github.com/<ProjectName>/internal/adapters/http/routers"
	repositories "github.com/<ProjectName>/internal/adapters/repositories/<featureName>"
	services "github.com/<ProjectName>/internal/core/services/<featureName>"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppContainer(app *fiber.App, db *gorm.DB) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	<FeatureName>App(route, db)
	return app
}

func <FeatureName>App(r routers.RouterImpl, db *gorm.DB) {
	transactorRepo := database.NewTransactorRepo(db)
	<featureName>Repo := repositories.New<FeatureName>Repo(db)
	<featureName>Srv := services.New<FeatureName>Service(<featureName>Repo, transactorRepo)
	<featureName>Handlers := handlers.New<FeatureName>Handler(<featureName>Srv)
	r.Create<FeatureName>Route(<featureName>Handlers)
}
```

### Explanation
- AppContainer Function: The entry point for integrating the feature into the application. It groups routes under the /v1 prefix and calls the specific feature's setup function.

- <FeatureName>App Function: Sets up the repositories, services, and handlers for the feature, and registers the routes using the router.


### Example
```bash
go run ./generators/app -feature User -output ./internal/app -project my_project
```
The tool will generate a file named `user_app.go` in the `./internal/app` directory, containing:
```go
package app

import (
	"github.com/my_project/internal/adapters/database"
	handlers "github.com/my_project/internal/adapters/http/handlers/user"
	"github.com/my_project/internal/adapters/http/routers"
	repositories "github.com/my_project/internal/adapters/repositories/user"
	services "github.com/my_project/internal/core/services/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppContainer(app *fiber.App, db *gorm.DB) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	UserApp(route, db)
	return app
}

func UserApp(r routers.RouterImpl, db *gorm.DB) {
	transactorRepo := database.NewTransactorRepo(db)
	userRepo := repositories.NewUserRepo(db)
	userSrv := services.NewUserService(userRepo, transactorRepo)
	userHandlers := handlers.NewUserHandler(userSrv)
	r.CreateUserRoute(userHandlers)
}
```


### App Generator Usage Notes
The generated file will replace placeholder text such as <ProjectName> and <FeatureName> with the actual project and feature names provided via the command-line flags.
The output directory must exist or will be created if it doesn't.