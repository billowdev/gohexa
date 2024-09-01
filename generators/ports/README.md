## Ports Generator
### Overview
The Ports Generator tool creates Go interface files for ports in a specified feature. Ports are used to define the interfaces for repositories and services in your application. This tool helps in generating these interfaces with the necessary methods for interacting with the data and performing business logic.

### Flags and Parameters
- `-feature <FeatureName>`: The name of the feature for which to generate ports (e.g., Order).
- `-output <OutputDirectory>`: The directory where the generated port files will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).

### Template Content
- The template generates a Go file with interfaces for:
	- I<FeatureName>Repository: Methods for CRUD operations on the model.
	- I<FeatureName>Service: Methods for CRUD operations and service-level logic.
- Includes placeholders for context and data types.
- The IDType field defaults to uint. Adjustments can be made if using UUIDs.

### Command
To generate a ports file, use the following command:
```bash
go run github.com/rapidstellar/gohexa/generators/ports -generate port -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
1. Generate Ports File:
```bash
go run github.com/rapidstellar/gohexa/generators/ports -generate port -feature="Order" -output ./internal/core/ports -project my_project
```

This command generates an order_ports.go file in the `./internal/core/ports` directory.

### Template Example
Here is a sample of the generated ports file based on the template:
```go
package ports

import (
	"context"

	"github.com/my_project/internal/adapters/database/models"
	domain "github.com/my_project/internal/core/domain/order"
	"github.com/my_project/pkg/helpers/pagination"
	"github.com/my_project/pkg/utils"
)

type IOrderRepository interface {
	GetOrder(ctx context.Context, id uint) (*models.Order, error)
	GetOrders(ctx context.Context) (*pagination.Pagination[[]models.Order], error)
	CreateOrder(ctx context.Context, payload *models.Order) error
	UpdateOrder(ctx context.Context, payload *models.Order) error
	DeleteOrder(ctx context.Context, id uint) error
}

type IOrderService interface {
	GetOrder(ctx context.Context, id uint) utils.APIResponse
	GetOrders(ctx context.Context) pagination.Pagination[[]domain.OrderDomain]
	CreateOrder(ctx context.Context, payload domain.OrderDomain) utils.APIResponse
	UpdateOrder(ctx context.Context, payload domain.OrderDomain) utils.APIResponse
	DeleteOrder(ctx context.Context, id uint) utils.APIResponse
}
```

### Ports Generators Usage Notes
- Ensure that the output directory exists or is created by the tool.
- Adjust the featureName to match your domain model naming conventions.
- The IDType defaults to uint. Modify the generatePortsFile function if UUID is used to change IDType accordingly.