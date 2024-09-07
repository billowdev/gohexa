package database

import (
	"fmt"

	"hexagonal/internal/adapters/database/models"
	"hexagonal/internal/adapters/database/seeders"
	"hexagonal/pkg/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {
	if configs.DB_SCHEMA == "" {
		configs.DB_SCHEMA = "public"
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v search_path=%v",
		configs.DB_HOST,
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_NAME,
		configs.DB_PORT,
		configs.DB_SSL_MODE,
		configs.DB_SCHEMA,
	)
	loggerDBLevel := logger.Silent
	if configs.APP_DEBUG_MODE {
		loggerDBLevel = logger.Info
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		// PreferSimpleProtocol: DB_DRY_RUN,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(loggerDBLevel), // or logger.Silent if you don't want logs
		// Logger: logger.Default.LogMode(logger.Info), // or logger.Silent if you don't want logs
	})

	if err != nil {
		return nil, err // instead of panic, return the error
	}
	// !: ENABLE MIGRATIONS DB
	if configs.DB_RUN_MIGRATION {
		if err := RunMigrations(db); err != nil {
			return nil, err
		}
	}
	if configs.DB_RUN_SEEDER {
		RunSeeds(db)
	}
	return db, nil
}

func RunSeeds(db *gorm.DB) {
	seeders.SeedTodo(db)

}

func resetSeeder(db *gorm.DB) error {
	if err := helperDeleteInfo(db, models.TNTodo); err != nil {
		return err
	}
	return nil
}

func helperDeleteInfo(db *gorm.DB, table string) error {
	err := db.Exec(fmt.Sprintf("DELETE FROM %s", table)).Error
	if err != nil {
		return err
	}
	err = db.Exec(fmt.Sprintf("SELECT setval('%s_id_seq', 1, false)", table)).Error
	if err != nil {
		return err
	}
	return nil
}

func RunMigrations(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
		if err != nil {
			return err
		}

		err = tx.AutoMigrate(
			// TODO START USER
			&models.Todo{},
		)
		if err != nil {
			return err
		}
		// TODO SHOULD COMMENT OUT IF NOT SIMULATE DATA
		// err = resetSeeder(db)
		// if err != nil {
		// 	return err
		// }
		return err
	})

	return err
}
