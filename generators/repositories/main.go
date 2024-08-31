// generators repositories
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	featureName := flag.String("feature", "", "The name of the feature to generate a repository for (e.g., Order)")
	outputDir := flag.String("output", "", "The output directory for the generated repository files")
	projectName := flag.String("project", "my_project", "The name of the project (default: my_project)")
	flag.Parse()

	if *featureName == "" {
		fmt.Println("Please provide a feature name using the -feature flag.")
		return
	}

	if *outputDir == "" {
		fmt.Println("Please provide an output directory using the -output flag.")
		return
	}

	// Ensure the output directory exists, create if not
	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	generateRepoFile(*outputDir, *featureName, *projectName)
}

var RepoTemplate = `
package repositories

import (
	"context"

	"github.com/{{ .ProjectName }}/internal/adapters/database"
	"github.com/{{ .ProjectName }}/internal/adapters/database/models"
	ports "github.com/{{ .ProjectName }}/internal/core/ports/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/pkg/helpers/filters"
	"github.com/{{ .ProjectName }}/pkg/helpers/pagination"
	"gorm.io/gorm"
)

type {{ .FeatureName }}Impl struct {
	db *gorm.DB
}

func New{{ .FeatureName }}Repository(db *gorm.DB) ports.I{{ .FeatureName }}Repository {
	return &{{ .FeatureName }}Impl{db: db}
}

// Create{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Create{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

// Delete{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Delete{{ .FeatureName }}(ctx context.Context, id uint) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Where("id=?", id).Delete(&models.{{ .FeatureName }}{}).Error; err != nil {
		return err
	}
	return nil
}

// Get{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Get{{ .FeatureName }}(ctx context.Context, id uint) (*models.{{ .FeatureName }}, error) {
	tx := database.HelperExtractTx(ctx, o.db)

	var data models.{{ .FeatureName }}
	if err := tx.WithContext(ctx).Where("id =?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// Get{{ .FeatureName }}s implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Get{{ .FeatureName }}s(ctx context.Context) (*pagination.Pagination[[]models.{{ .FeatureName }}], error) {
	tx := database.HelperExtractTx(ctx, o.db)

	p := pagination.GetFilters[filters.{{ .FeatureName }}Filter](ctx)
	fp := p.Filters

	orderBy := pagination.NewOrderBy(pagination.SortParams{
		Sort:           p.Sort,
		Order:          p.Order,
		DefaultOrderBy: "updated_at DESC",
	})
	tx = pagination.ApplyFilter(tx, "id", fp.ID, "contains")
	tx = tx.WithContext(ctx).Order(orderBy)
	data, err := pagination.Paginate[filters.{{ .FeatureName }}Filter, []models.{{ .FeatureName }}](p, tx)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Update{{ .FeatureName }} implements ports.I{{ .FeatureName }}Repository.
func (o *{{ .FeatureName }}Impl) Update{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error {
	tx := database.HelperExtractTx(ctx, o.db)
	if err := tx.WithContext(ctx).Save(&payload).Error; err != nil {
		return err
	}
	return nil
}
`

func generateRepoFile(dir string, featureName string, projectName string) {
	// Define the template for the repository file

	// Prepare the data for template rendering
	data := struct {
		FeatureName string
		ProjectName string
	}{
		FeatureName: featureName,
		ProjectName: projectName,
	}

	// Parse and execute the template
	tmpl, err := template.New("repo").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(RepoTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_repository.go", strings.ToLower(featureName))
	filePath := filepath.Join(dir, fileName)

	// Create the output file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Execute the template and write to the file
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	} else {
		fmt.Printf("Repository file '%s' created successfully!\n", filePath)
	}
}
