package seeders

import (
	"hexagonal/internal/adapters/database/models"

	"gorm.io/gorm"
)

var SEED_TODO = []byte(`[
	{
		"id": 1,
		"name": "Golang CLI",
		"status": false
	},
	{
		"id": 2,
		"name": "Golang Template",
		"status": true
    }
]`)

func SeedTodo(db *gorm.DB) error {
	var data []models.Todo
	seed := SEED_TODO
	if err := BaseSeeder(db, &data, seed); err != nil {
		return err
	}
	return nil
}
