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
go run ./generators/domain -feature <FeatureName> -output <OutputDirectory> -project <ProjectName> [-uuid]
```
Example Commands
Without UUID:
```bash
go run ./generators/domain -feature="Todo" -output ./internal/adapters/domain -project my_project
```
This command generates a todo_domain.go file in the ./internal/adapters/domain directory with an auto-increment integer ID.

With UUID:
```bash
go run ./generators/domain -feature="Todo" -output ./internal/adapters/domain -project my_project -uuid
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