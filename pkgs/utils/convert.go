package utils

import "strings"

// ToLower returns the lowercase version of the input string
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Pluralize returns the plural form of the input string (simple example)
func Pluralize(s string) string {
	if strings.HasSuffix(s, "s") {
		return s
	}
	return s + "s"
}
