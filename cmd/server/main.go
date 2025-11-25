package main

import (
	router "GolangBackendEcommerce/internal/routers"
	"log"
)

func main() {
	r := router.NewRouter()

	if err := r.Run(":8002"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
