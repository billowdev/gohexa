package database

import (
	"context"

	"gorm.io/gorm"
)

// Idea https://www.kaznacheev.me/posts/en/clean-transactions-in-hexagon/
type txKey struct{}

// injectTx injects the transaction into the context
func InjectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts the transaction from the context
func ExtractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

func HelperExtractTx(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx := ExtractTx(ctx)
	if tx == nil {
		tx = db
	}
	return tx
}
