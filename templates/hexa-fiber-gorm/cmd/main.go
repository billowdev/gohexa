package main

import (
	"fmt"
	"log"

	"hexagonal/internal/adapters/app"
	"hexagonal/internal/adapters/database"
	"hexagonal/pkg/configs"
)

func main() {
	params := configs.NewFiberHttpServiceParams()
	fiberHTTP := configs.NewFiberHTTPService(params)
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal("Failed to start Database:", err)
	}

	if err != nil {
		log.Fatal("Failed to connect Temporal client:", err)
	}

	app.AppContainer(fiberHTTP, db)
	portString := fmt.Sprintf(":%v", params.Port)

	err = fiberHTTP.Listen(portString)

	if err != nil {
		log.Fatal("Failed to start golang Fiber server:", err)
	}

}
