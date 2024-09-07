package domain

type RouteFlagDomain struct {
	FeatureName string
	ProjectName string
}

var RouteTemplate = `
package routers

import (
	handlers "{{ .ProjectName }}/internal/adapters/handlers/{{ .FeatureName | ToLower }}"
	"{{ .ProjectName }}/pkg/middlewares"
)

func (r RouterImpl) Create{{ .FeatureName }}Routes(h handlers.I{{ .FeatureName }}Handler) {
	r.route.Get("/{{ .FeatureName | Pluralize | ToLower }}", h.HandleGet{{ .FeatureName }}s)
	r.route.Get("/{{ .FeatureName | Pluralize | ToLower }}/:id", h.HandleGet{{ .FeatureName }})
	r.route.Post("/{{ .FeatureName | Pluralize | ToLower }}", h.HandleCreate{{ .FeatureName }})
	r.route.Put("/{{ .FeatureName | Pluralize | ToLower }}/:id", h.HandleUpdate{{ .FeatureName }})
	r.route.Delete("/{{ .FeatureName | Pluralize | ToLower }}/:id", h.HandleDelete{{ .FeatureName }})
}
`
