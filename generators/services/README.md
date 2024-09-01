## Services Generator

### Overview
The Services Generator tool creates Go service implementation files for a specified feature. Services handle the business logic and coordinate interactions between repositories and other components. This tool generates these service implementations based on the provided feature name and project details.

### Flags and Parameters
- `-feature <FeatureName>`: The name of the feature for which to generate a service (e.g., Order).
- `-output <OutputDirectory>`: The directory where the generated service files will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).

### Template Content
- The template generates a Go file with an implementation of the service interface for the specified feature.
- Includes methods for:
	- Create<FeatureName>: Create a new record and return a response.
	- Delete<FeatureName>: Delete a record by ID and return a response.
	- Get<FeatureName>: Get a record by ID and return a response.
	- Get<FeatureName>s: Get a list of records with pagination and return a response.
	- Update<FeatureName>: Update an existing record and return a response.
  
Command
To generate a service file, use the following command:
```bash
go run github.com/rapidstellar/gohexa/generators/services -generate service -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
1. Generate Service File:
```bash
go run github.com/rapidstellar/gohexa/generators/services -generate service -feature="Order" -output ./internal/core/services -project my_project
```
This command generates an order_service.go file in the ./internal/core/services directory.

### Template Example
Here is a sample of the generated service file based on the template:

```go
package services

import (
	"context"

	"github.com/my_project/internal/adapters/database"
	domain "github.com/my_project/internal/core/domain/order"
	ports "github.com/my_project/internal/core/ports/order"
	"github.com/my_project/pkg/configs"
	"github.com/my_project/pkg/helpers/pagination"
	"github.com/my_project/pkg/utils"
)

type OrderServiceImpl struct {
	repo       ports.IOrderRepository
	transactor database.IDatabaseTransactor
}

func NewOrderService(
	repo ports.IOrderRepository,
	transactor database.IDatabaseTransactor,
) ports.IOrderService {
	return &OrderServiceImpl{repo: repo, transactor: transactor}
}

// CreateOrder implements ports.IOrderService.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, payload domain.Order) utils.APIResponse {
	data := domain.ToOrderModel(payload)
	if err := s.repo.CreateOrder(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// DeleteOrder implements ports.IOrderService.
func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, id uint) utils.APIResponse {
	if err := s.repo.DeleteOrder(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// GetOrder implements ports.IOrderService.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, id uint) utils.APIResponse {
	data, err := s.repo.GetOrder(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	res := domain.ToOrderDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// GetOrders implements ports.IOrderService.
func (s *OrderServiceImpl) GetOrders(ctx context.Context) pagination.Pagination[[]domain.Order] {
	data, err := s.repo.GetOrders(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.Order]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.ToOrderDomain)
	return pagination.Pagination[[]domain.Order]{
		Rows:       newData,
		Links:      data.Links,
		Total:      data.Total,
		Page:       data.Page,
		PageSize:   data.PageSize,
		TotalPages: data.TotalPages,
	}
}

// UpdateOrder implements ports.IOrderService.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, payload domain.Order) utils.APIResponse {
	data := domain.ToOrderModel(payload)
	if err := s.repo.UpdateOrder(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.ToOrderDomain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}
```

### Services Generators Usage Notes
- Ensure that the output directory exists or is created by the tool.
- Adjust the featureName to match your domain model naming conventions.
