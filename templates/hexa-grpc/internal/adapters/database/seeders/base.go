package seeders

import (
	"encoding/json"

	"gorm.io/gorm"
)

func BaseSeeder[T any](db *gorm.DB, data *[]T, seed []byte) error {
	if err := json.Unmarshal(seed, &data); err != nil {
		return err
	}
	// Begin a transaction
	tx := db.Begin()

	// Iterate over the data slice and create each entry in the database
	for _, v := range *data {
		if err := tx.Create(&v).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
