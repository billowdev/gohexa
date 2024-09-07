package domain

import (
	"hexagonal/internal/adapters/database/models"
	"time"
)

type TodoDomain struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
}

func ToTodoDomain(data *models.Todo) TodoDomain {
	if data == nil {
		return TodoDomain{

			ID: 0,
		}
	}

	return TodoDomain{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Name:      data.Name,
		Status:    data.Status,
	}
}

func ToTodoModel(data TodoDomain) *models.Todo {
	return &models.Todo{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Name:      data.Name,
		Status:    data.Status,
	}
}
