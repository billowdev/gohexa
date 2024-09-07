package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type TransactorImpl struct {
	db *gorm.DB
}

// BeginTransaction implements IDatabasePorts.
func (d *TransactorImpl) BeginTransaction() (*gorm.DB, error) {
	tx := d.db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	return tx, nil
}

// RollbackTransaction rolls back the transaction if it was started and returns any error encountered.
func (d *TransactorImpl) RollbackTransaction(tx *gorm.DB) error {
	if tx == nil {
		return nil // No transaction to rollback
	}
	if tx.Error != nil {
		return tx.Error // If there was an error, return it
	}

	// Rollback the transaction
	if err := tx.Rollback().Error; err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}

// WithinTransaction implements IDatabaseTransactor.
// WithinTransaction runs the provided function within a transaction context.
// The transaction is automatically committed if the function completes successfully, or rolled back if an error occurs.
func (d *TransactorImpl) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	// begin transaction
	tx, err := d.BeginTransaction()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", tx.Error)
	}

	// Ensure that the transaction is finalized properly
	defer func() {
		if r := recover(); r != nil {
			_ = d.RollbackTransaction(tx)
			panic(r) // Re-panic after rollback
		} else if tx.Error != nil {
			_ = d.RollbackTransaction(tx)
		} else {
			if commitErr := tx.Commit().Error; commitErr != nil {
				log.Printf("failed to commit transaction: %v", commitErr)
				err = commitErr
			}
		}
	}()

	// Run the callback function with the transaction context
	err = tFunc(InjectTx(ctx, tx))
	if err != nil {
		tx.Error = err // Set the error to indicate a rollback is needed
		return err
	}

	return nil
}

// WithTransactionContextTimeout executes a function within a transaction with a specified context timeout.
// The transaction is committed if successful, or rolled back if an error occurs or the context times out.
func (d *TransactorImpl) WithTransactionContextTimeout(ctx context.Context, timeout time.Duration, tFunc func(ctx context.Context) error) error {
	// Create a new context with timeout
	transactionCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Start a new transaction
	tx, err := d.BeginTransaction()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Ensure that the transaction is finalized properly
	defer func() {
		select {
		case <-transactionCtx.Done():
			// Rollback if the transaction context is done (timeout or cancel)
			if rollbackErr := d.RollbackTransaction(tx); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
			}
		default:
			// Commit if no error and context is still valid
			if commitErr := tx.Commit().Error; commitErr != nil {
				log.Printf("failed to commit transaction: %v", commitErr)
				err = commitErr
			}
		}
	}()

	// Run the callback function with the transaction context
	err = tFunc(InjectTx(transactionCtx, tx))
	if err != nil {
		tx.Error = err // Mark the transaction as needing a rollback
		return err
	}

	return nil
}

type IDatabaseTransactor interface {
	// InjectTx(ctx context.Context, tx *gorm.DB) context.Context
	// ExtractTx(ctx context.Context) *gorm.DB
	WithinTransaction(context.Context, func(ctx context.Context) error) error
	WithTransactionContextTimeout(ctx context.Context, timeout time.Duration, tFunc func(ctx context.Context) error) error
	BeginTransaction() (*gorm.DB, error)
	RollbackTransaction(tx *gorm.DB) error
}

func NewTransactorRepo(db *gorm.DB) IDatabaseTransactor {
	return &TransactorImpl{db: db}
}
