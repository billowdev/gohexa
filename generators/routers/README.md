## Routes Generator

### Overview

The Routes Generator tool creates Go route handling files for a specified feature. Routes define the HTTP endpoints for interacting with your service, including creating, reading, updating, and deleting resources. This tool generates route handling implementations based on the provided feature name and project details.

### Flags and Parameters
- `feature <FeatureName>`: The name of the feature for which to generate routes (e.g., SeaPort).
- `output <OutputDirectory>`: The directory where the generated route files will be saved.
- `project <ProjectName>`: The name of the project (default is my_project).

### Template Content
- The template generates a Go file with route definitions for the specified feature.
- Includes routes for:
	- GET /<feature>: List all resources.
	- GET /<feature>/:id: Get a single resource by ID.
	- POST /<feature>: Create a new resource.
	- PUT /<feature>/:id: Update an existing resource by ID.
	- DELETE /<feature>/:id: Delete a resource by ID.

### Command
To generate a route file, use the following command:
```bash
go run github.com/rapidstellar/gohexa/generators/routes -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
1. Generate Route File:
```bash
go run github.com/rapidstellar/gohexa/generators/routes -feature="SeaPort" -output ./internal/routers -project my_project
```
This command generates a seaport_routes.go file in the ./internal/routers directory.

### Template Example
Here is a sample of the generated route file based on the template:

```go
package routers

import (
	handlers "my_project/internal/adapters/handlers/seaport"
	"my_project/pkg/middlewares"
)

func (r RouterImpl) CreateSeaPortRoutes(h handlers.ISeaPortHandler) {
	r.route.Get("/seaports", h.HandleGetSeaPorts)
	r.route.Get("/seaports/:id", h.HandleGetSeaPort)
	r.route.Post("/seaports", h.HandleCreateSeaPort)
	r.route.Put("/seaports/:id", h.HandleUpdateSeaPort)
	r.route.Delete("/seaports/:id", h.HandleDeleteSeaPort)
}
```
### Routers Generators Usage Notes
Ensure that the output directory exists or is created by the tool.
Adjust the featureName to match your domain naming conventions and route requirements.

