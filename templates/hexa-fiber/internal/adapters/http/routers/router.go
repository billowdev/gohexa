package routers

import "github.com/gofiber/fiber/v2"

type RouterImpl struct {
	route fiber.Router
}

func NewRoute(r fiber.Router) RouterImpl {
	return RouterImpl{r}
}
