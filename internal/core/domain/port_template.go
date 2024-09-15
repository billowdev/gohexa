package domain

type PortFlagDomain struct {
	FeatureName string
	ProjectName string
	IDType      string
}

var PortsTemplate = `
package ports

import (
	"context"

	"github.com/{{ .ProjectName }}/internal/adapters/database/models"
	domain "github.com/{{ .ProjectName }}/internal/core/domain/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/pkg/helpers/pagination"
	"github.com/{{ .ProjectName }}/pkg/utils"
)

type I{{ .FeatureName }}Repository interface {
	BulkCreate{{ .FeatureName }}(ctx context.Context, payloads []*models.{{ .FeatureName }}) error
	GetAllByFields(ctx context.Context, filters map[string]interface{}, limit int)
	GetOneByFields(ctx context.Context, filters map[string]interface{})
	Get{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) (*models.{{ .FeatureName }}, error)
	Get{{ .FeatureName }}s(ctx context.Context) (*pagination.Pagination[[]models.{{ .FeatureName }}], error)
	Create{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error
	Update{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error
	Delete{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) error
}

type I{{ .FeatureName }}Service interface {
	Get{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) utils.APIResponse
	Get{{ .FeatureName }}s(ctx context.Context) pagination.Pagination[[]domain.{{ .FeatureName }}Domain]
	Create{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}Domain) utils.APIResponse
	Update{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}Domain) utils.APIResponse
	Delete{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) utils.APIResponse
}
`
