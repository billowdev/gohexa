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
go run ./generators/models -feature <FeatureName> -output <OutputDirectory> -project <ProjectName> -uuid
```

### Example Commands
1. Generate Model File with UUID:
```bash
go run ./generators/models -feature="Todo" -output ./internal/adapters/database/models -project my_project -uuid
```
This command generates a todo.go file in the ./internal/adapters/database/models directory with UUID as the ID field.

2. Generate Model File with Auto-Increment ID:
```bash
go run ./generators/models -feature="Todo" -output ./internal/adapters/database/models -project my_project
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
