package utils

// ConvertSlice converts a slice of model instances to a slice of domain instances.
// It uses the provided conversion function to transform each model instance into a domain instance.
//
// Parameters:
// - models: A slice of model instances to be converted.
// - convert: A function that takes a pointer to a model instance and returns a domain instance.
//
// Returns:
// - A slice of converted domain instances. If the input slice is nil, the function returns nil.
func ConvertSlice[TModel any, TDomain any](models []TModel, convert func(*TModel) TDomain) []TDomain {
	if models == nil {
		return nil
	}

	domains := make([]TDomain, len(models))
	for i, model := range models {
		domains[i] = convert(&model)
	}
	return domains
}
