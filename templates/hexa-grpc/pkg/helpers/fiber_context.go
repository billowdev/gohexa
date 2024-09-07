package helpers

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Define a custom type for the context key
type contextKey string

// Helper function to set the subject in the context
func HelperSetSubject(ctx context.Context, c *fiber.Ctx) context.Context {
	const subKey contextKey = "sub"

	// Extract 'sub' from the Fiber context
	sub := c.Locals("sub")

	// Add 'sub' to the provided context with the custom key
	return context.WithValue(ctx, subKey, sub)
}

// Helper function to get the subject from the context
func HelperGetSubject(ctx context.Context) (string, error) {
	const subKey contextKey = "sub"

	// Retrieve the 'sub' value from the context
	sub, ok := ctx.Value(subKey).(string)
	if !ok || sub == "" {
		return "", errors.New("subject not found in context")
	}

	return sub, nil
}
