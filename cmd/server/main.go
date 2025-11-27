package main

import "GolangBackendEcommerce/internal/initialize"

func main() {
	// r := router.NewRouter()

	// if err := r.Run(); err != nil {
	// 	log.Fatalf("failed to run server: %v", err)
	// }
	initialize.Run()
}
