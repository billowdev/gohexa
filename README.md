# Get started

## Setup environment variables for `gohexa` CLI 

in each platform and run your CLI tool using those variables, you can follow these platform-specific instructions. Here's how you can do it:

### 1. macOS/Linux
Step 1: Set Environment Variables Temporarily
You can set environment variables temporarily for a single command by prepending the command with the environment variables.

```bash
MY_VAR="some_value" OTHER_VAR="another_value" ./build/gohexa-mac
```

Step 2: Set Environment Variables Persistently
To set environment variables persistently, you can add them to your shell profile (e.g., `~/.bashrc`, `~/.bash_profile`, `~/.zshrc` for Zsh users):

```bash
alias gohexa="~/path/to/build/gohexa-mac"
```

After adding the variables, source the profile to apply the changes:
```bash
source ~/.bashrc
```

Now, you can run your CLI tool:
```bash
gohexa -name new_project
```

### 2. Windows
Step 1: Set Environment Variables Temporarily
To set environment variables temporarily for a single command in Command Prompt:
```bash
set MY_VAR=some_value && set OTHER_VAR=another_value && build\gohexa.exe
```
In PowerShell:
```bash
$env:MY_VAR="some_value"; $env:OTHER_VAR="another_value"; .\build\gohexa.exe
```

Step 2: Set Environment Variables Persistently
To set environment variables persistently in Windows:

1. Open the Start menu, search for "Environment Variables," and select "Edit the system environment variables."
2. In the System Properties window, click "Environment Variables."
Under "User variables" or "System variables," click "New" and add your variable name and value.
3. Once set, these environment variables will be available to all command-line sessions, and you can run your CLI tool:

```bash
build\gohexa.exe
```

### 3. Using Environment Variables in Go
In your Go code, you can access these environment variables using the os.Getenv function:
```
package main

import (
	"fmt"
	"os"
)

func main() {
	myVar := os.Getenv("MY_VAR")
	otherVar := os.Getenv("OTHER_VAR")

	fmt.Println("MY_VAR:", myVar)
	fmt.Println("OTHER_VAR:", otherVar)

	// Your CLI logic here
}
```
### 4. Running the CLI with Environment Variables
After setting the environment variables, you can run the CLI on each platform, and it will pick up those variables.

macOS/Linux:
```bash
./build/gohexa-mac
```

Windows:
```bash
build\gohexa.exe
```

### Conclusion
By setting environment variables either temporarily or persistently on each platform, you can control the runtime environment of your CLI tool. The CLI can then access these variables using Go’s `os.Getenv` function, allowing you to configure the tool's behavior based on the environment.


## go-hexagonal
go hexagonal template

```bash
gohexa -generate project -output myproject
```
or
```bash
gohexa -generate project -output myproject -template hexagonal
```

### Golang Hexagonal Example

- Help
```
gohexa --help
```

- Fiber 
```bash
gohexa -generate project -output myproject -template hexa-fiber
```
- gRPC 
```bash
gohexa -generate project -output myproject -template hexa-grpc
```

### Hexagonal CRUD Files generator example

#### database transactor
```bash
gohexa -generate transactor -output="./internal/adapters/database"
```

#### models generator (gorm)
```bash
gohexa -generate model -feature="Todo" -output="./internal/adapters/database/models" -uuid=true
```
or 
```bash
gohexa -generate model -feature="Todo" -output="./internal/adapters/database/models" -uuid=false
```

#### domain generator (gorm)
```bash
gohexa -generate domain -feature="Todo" -output="./internal/core/domain" -project="my_project" -uuid=true
```
or uint id increamenter
```bash
gohexa -generate domain -feature="Todo" -output="./internal/core/domain" -project="my_project" -uuid=false
```

#### repo generator
```bash
gohexa -generate repository -feature="Todo" -output="./internal/adapters/repositories" -project="my_project"
```

#### services generator
```bash
gohexa -generate service -feature="Todo" -output="./internal/core/services" -project="my_project"
```

#### ports generator
```bash
gohexa -generate port -feature="Todo" -output="./internal/core/ports" -project="my_project"
```

#### handlers generator
```bash
gohexa -feature="Todo" -output="./internal/adapters/http/handlers" -project=my_project
```
#### routes generator
```bash
gohexa -feature="Todo" -output="./internal/adapters/http/routers" -project=my_project
```

#### app generator
```bash
gohexa -feature="Todo" -output ./internal/adapters/app -project my_project
```
or
```bash
gohexa -generate app -feature="Todo" -output ./internal/adapters/app -project my_project
```


# Project Generator

## Overview
The Project Generator tool creates a new project directory structure based on a specified template. It sets up a project with pre-defined folders and files, replacing placeholder values with the provided project name.

## Flags and Parameters
- `-name <ProjectName>`: The name of the new project.
- `-template <TemplateName>`: The name of the template to use (default is hexagonal).

## Command
To generate a new project, use the following command:
```bash
gohexa -output <ProjectName> -template <TemplateName>
```

## Example Commands
1. Generate Project Using Default Template:
```bash
gohexa -output MyNewProject
```
This command creates a new project named MyNewProject using the default hexagonal template.

2. Generate Project Using a Custom Template:
```bash
gohexa -output MyCustomProject -template custom_template
```
This command creates a new project named MyCustomProject using the custom_template template.

## Template Structure
- Template Directory: The template directory contains the folder structure and files to be copied to the new project.
- Placeholder Replacement: All instances of the placeholder go-template in files within the template directory will be replaced with the specified project name.

## Usage Notes
- Ensure the template directory exists and is structured as desired before running the command.
- The tool will create the new project directory and copy all files from the template directory, replacing placeholders in the files.


## Example
Given a template directory structure like this:

```bash
hexagonal/
├── cmd/
│   └── main.go
├── internal/
│   ├── adapters/
│   └── core/
└── README.md
```


# Get started with RPS Hexa generator

## Transactor File Generator

### Overview
The Transactor File Generator is a command-line tool that creates a transactor.go file in the specified output directory. This file contains a set of transaction management utilities for use with the Gorm ORM in a hexagonal architecture.

### Features
- Generates a transactor.go file with pre-defined transaction management functions.
- Supports injecting, extracting, and managing database transactions within a context.
- Provides functions for handling transactions with timeout contexts.

### Usage
#### Flags
`-output <OutputDirectory>`: The output directory where the transactor.go file will be generated.

### Command
To generate the transactor.go file, run the following command:

```bash
gohexa -generate transactor -output <OutputDirectory>
```

### Example Command
Generate the transactor.go file in the ./database directory:

```bash
gohexa -generate transactor -output ./database
```

### Output
The tool generates a file named transactor.go in the specified directory. The generated file includes:

- Transaction Management: Functions for beginning, committing, and rolling back transactions.
- Context Injection/Extraction: Utilities for injecting and extracting the transaction from the context.
- Timeout Handling: Functions for executing transactions with a context timeout.

### Generated File Structure
The transactor.go file includes the following structure:

- Package: database
- Transaction Utilities:
	- InjectTx(ctx context.Context, tx *gorm.DB) context.Context
	- ExtractTx(ctx context.Context) *gorm.DB
	- HelperExtractTx(ctx context.Context, db *gorm.DB) *gorm.DB

- Transaction Management:
	- BeginTransaction() (*gorm.DB, error)
	- RollbackTransaction(tx *gorm.DB) error
	- WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	- WithTransactionContextTimeout(ctx context.Context, timeout time.Duration, tFunc func(ctx context.Context) error) error
- Interfaces:
	- IDatabaseTransactor

### Notes
Ensure that the output directory exists or the tool will create it.
The generated transactor.go file is designed to work with the Gorm ORM in a hexagonal architecture.

## Models Generator

### Overview
The Models Generator tool creates Go model files for a specified feature. It supports generating models with either UUID or auto-incrementing ID fields and uses GORM for ORM functionality. The generated models are designed to work with your project's database setup and include timestamp fields for tracking creation and updates.

### Flags and Parameters
- `-feature <FeatureName>`: The name of the feature for which to generate the model (e.g., Order).
- `-output <OutputDirectory>`: The directory where the generated model file will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).
- `-uuid`: Whether to use UUID for the ID field (default is false).

### Template Content
- The template generates a Go file with a model struct that includes:
	- gorm.Model fields (ID, CreatedAt, UpdatedAt, DeletedAt).
	- Optionally uses UUIDs for the ID field if -uuid flag is set.
	- Custom table name function TableName() for GORM.
- The model includes JSON tags for serialization and GORM tags for database mapping.

### Command
To generate a model file, use the following command:
```bash
gohexa -generate model -feature <FeatureName> -output <OutputDirectory> -project <ProjectName> -uuid
```

### Example Commands
1. Generate Model File with UUID:
```bash
gohexa -generate model -feature="Todo" -output ./internal/adapters/database/models -project my_project -uuid
```
This command generates a todo.go file in the ./internal/adapters/database/models directory with UUID as the ID field.

2. Generate Model File with Auto-Increment ID:
```bash
gohexa -generate model -feature="Todo" -output ./internal/adapters/database/models -project my_project
```
This command generates a `todo.go` file in the `./internal/adapters/database/models` directory with auto-incrementing ID.

### Template Example
Here is a sample of the generated model file based on the template:
```go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID                 uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt          time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

var TnTodo = "todos"

func (st *Todo) TableName() string {
	return TnTodo
}
```

### Models Genrators Usage Notes
- Ensure that the output directory exists or is created by the tool.
- Adjust the featureName to match your domain model naming conventions.
- The -uuid flag should be used if your project requires UUIDs for IDs; otherwise, the default is auto-incrementing IDs.


## Domain Generator

### Overview
This tool generates Go domain files based on a specified feature name. It supports both UUID and auto-increment integer IDs for your domain models. The generated files include a domain struct and helper functions for converting between domain and model representations.

### Flags and Parameters
- `-feature <FeatureName>`: The name of the feature for which to generate the domain file (e.g., Order).
- `-output <OutputDirectory>`: The directory where the generated domain file will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).
- `-uuid`: Use UUID for the ID field instead of an auto-incrementing integer. Add this flag to use UUIDs.

### Template Content
- The template generates a Go file with a domain struct for the specified feature.
- {{ .FeatureName | ToLower }} is used to convert the feature name to lowercase.
- The domain struct includes an ID field that can be either UUID or uint, based on the -uuid flag.
- The To{{ .FeatureName }}Domain function converts a model to a domain struct.
- The To{{ .FeatureName }}Model function converts a domain struct to a model.
- A helper function defaultStringIfEmpty is provided to handle default values for empty strings.

### Command
To generate a domain file, use the following command:
```bash
gohexa -generate domain -feature <FeatureName> -output <OutputDirectory> -project <ProjectName> [-uuid]
```
Example Commands
Without UUID:
```bash
gohexa -generate domain -feature="Todo" -output ./internal/adapters/domain -project my_project
```
This command generates a todo_domain.go file in the ./internal/adapters/domain directory with an auto-increment integer ID.

With UUID:
```bash
gohexa -generate domain -feature="Todo" -output ./internal/adapters/domain -project my_project -uuid
```
This command generates a todo_domain.go file in the ./internal/adapters/domain directory with a UUID as the ID.


### Template Example
Here is a sample of the generated domain file based on the template:

```go
package domain

import (
	"time"

	"github.com/my_project/internal/adapters/database/models"
)

type TodoDomain struct {
	ID                 string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Field1             string    `json:"field_1"`
	Field2             string    `json:"field_2"`
}

func ToTodoDomain(data *models.Todo) TodoDomain {
	if data == nil {
		return TodoDomain{
			ID: "00000000-0000-0000-0000-000000000000",
		}
	}

	return TodoDomain{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Field1:    data.Field1,
		Field2:    defaultStringIfEmpty(data.Field2, "No Field2"),
	}
}

func ToTodoModel(data TodoDomain) *models.Todo {
	return &models.Todo{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Field1:    data.Field1,
		Field2:    defaultStringIfEmpty(data.Field2, "No Field2"),
	}
}

func defaultStringIfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
```

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
gohexa -generate port -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
1. Generate Ports File:
```bash
gohexa -generate port -feature="Order" -output ./internal/core/ports -project my_project
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

## Repositories Generator

### Overview
The Repositories Generator tool creates Go implementation files for repositories in a specified feature. Repositories handle interactions with the database and provide methods for CRUD operations on your models. This tool generates these implementations based on the provided feature name and project details.

### Flags and Parameters
- `-feature <FeatureName>`: The name of the feature for which to generate a repository (e.g., Order).
- `-output <OutputDirectory>`: The directory where the generated repository files will be saved.
- `-project <ProjectName>`: The name of the project (default is my_project).

### Template Content
- The template generates a Go file with an implementation of the repository interface for the specified feature.
- Includes methods for:
	- Create<FeatureName>: Create a new record.
	- Delete<FeatureName>: Delete a record by ID.
	- Get<FeatureName>: Get a record by ID.
	- Get<FeatureName>s: Get a list of records with pagination and filters.
	- Update<FeatureName>: Update an existing record.

### Command
To generate a repository file, use the following command:
```bash
gohexa -generate repository -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
Generate Repository File:
```bash
gohexa -generate repository -feature="Order" -output ./internal/adapters/repositories -project my_project
```
This command generates an `order_repository.go` file in the `./internal/adapters/repositories` directory.

### Template Example
Here is a sample of the generated repository file based on the template:
```go
package repositories

import (
	"context"

	"github.com/my_project/internal/adapters/database"
	"github.com/my_project/internal/adapters/database/models"
	ports "github.com/my_project/internal/core/ports/order"
	"github.com/my_project/pkg/helpers/filters"
	"github.com/my_project/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type OrderImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.IOrderRepository {
	return &OrderImpl{db: db}
}

// CreateOrder implements ports.IOrderRepository.
func (o *OrderImpl) CreateOrder(ctx context.Context, payload *models.Order) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder implements ports.IOrderRepository.
func (o *OrderImpl) DeleteOrder(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.Order{}).Error; err != nil {
		return err
	}
	return nil
}

// GetOrder implements ports.IOrderRepository.
func (o *OrderImpl) GetOrder(ctx context.Context, id uint) (*models.Order, error) {
	tx := database.HelperExtractTx(ctx, o.db)

	var data models.Order
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// GetOrders implements ports.IOrderRepository.
func (o *OrderImpl) GetOrders(ctx context.Context) (*pagination.Pagination[[]models.Order], error) {
	tx := database.HelperExtractTx(ctx, o.db)

	p := pagination.GetFilters[filters.OrderFilter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.OrderFilter, []models.Order](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// UpdateOrder implements ports.IOrderRepository.
func (o *OrderImpl) UpdateOrder(ctx context.Context, payload *models.Order) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
```

### Repositories Generators Usage Notes
- Ensure that the output directory exists or is created by the tool.
- Adjust the featureName to match your domain model naming conventions.


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
gohexa -generate service -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
1. Generate Service File:
```bash
gohexa -generate service -feature="Order" -output ./internal/core/services -project my_project
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
gohexa -generate handler -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```
### Example Commands
1. Generate Handlers File:

```bash
gohexa -generate handler -feature="Todo" -output ./internal/adapters/handlers -project my_project
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
gohexa -generate route -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```

### Example Commands
1. Generate Route File:
```bash
gohexa -generate route -feature="SeaPort" -output ./internal/routers -project my_project
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
gohexa -generate app -feature <FeatureName> -output <OutputDirectory> -project <ProjectName>
```
- example
```bash
gohexa -generate app -feature="Todo" -output ./internal/adapters/app -project my_project
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
gohexa -generate app -feature User -output ./internal/app -project my_project
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
