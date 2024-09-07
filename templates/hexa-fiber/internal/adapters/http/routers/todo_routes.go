// gohexa -generate route -output ./internal/adapters/http/routers -feature Todo -project hexagonal
package routers


import "hexagonal/internal/adapters/http/handlers"

func (r RouterImpl) CreateTodoRoutes(h handlers.ITodoHandler) {
	r.route.Get("/todos", h.HandleGetTodoes)
	r.route.Get("/todos/:id", h.HandleGetTodo)
	r.route.Post("/todos", h.HandleCreateTodo)
	r.route.Put("/todos/:id", h.HandleUpdateTodo)
	r.route.Delete("/todos/:id", h.HandleDeleteTodo)
}
