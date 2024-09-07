package domain

type ServiceFlagDomain struct {
	FeatureName string
	ProjectName string
}

var ServiceTemplate = `
package services

import (
	"context"

	"github.com/{{ .ProjectName }}/internal/adapters/database"
	domain "github.com/{{ .ProjectName }}/internal/core/domain/{{ .FeatureName | ToLower }}"
	ports "github.com/{{ .ProjectName }}/internal/core/ports/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/pkg/configs"
	"github.com/{{ .ProjectName }}/pkg/helpers/pagination"
	"github.com/{{ .ProjectName }}/pkg/utils"
)

type {{ .FeatureName }}ServiceImpl struct {
	repo       ports.I{{ .FeatureName }}Repository
	transactor database.IDatabaseTransactor
}

func New{{ .FeatureName }}Service(
	repo ports.I{{ .FeatureName }}Repository,
	transactor database.IDatabaseTransactor,
) ports.I{{ .FeatureName }}Service {
	return &{{ .FeatureName }}ServiceImpl{repo: repo, transactor: transactor}
}

// Create{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Create{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}) utils.APIResponse {
	data := domain.To{{ .FeatureName }}Model(payload)
	if err := s.repo.Create{{ .FeatureName }}(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// Delete{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Delete{{ .FeatureName }}(ctx context.Context, id uint) utils.APIResponse {
	if err := s.repo.Delete{{ .FeatureName }}(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// Get{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Get{{ .FeatureName }}(ctx context.Context, id uint) utils.APIResponse {
	data, err := s.repo.Get{{ .FeatureName }}(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	res := domain.To{{ .FeatureName }}Domain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// Get{{ .FeatureName }}s implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Get{{ .FeatureName }}s(ctx context.Context) pagination.Pagination[[]domain.{{ .FeatureName }}] {
	data, err := s.repo.Get{{ .FeatureName }}s(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.{{ .FeatureName }}]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.To{{ .FeatureName }}Domain)
	return pagination.Pagination[[]domain.{{ .FeatureName }}]{
		Rows:       newData,
		Links:      data.Links,
		Total:      data.Total,
		Page:       data.Page,
		PageSize:   data.PageSize,
		TotalPages: data.TotalPages,
	}
}

// Update{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Update{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}) utils.APIResponse {
	data := domain.To{{ .FeatureName }}Model(payload)
	if err := s.repo.Update{{ .FeatureName }}(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.To{{ .FeatureName }}Domain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}
`
