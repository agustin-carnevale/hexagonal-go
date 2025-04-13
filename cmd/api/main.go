package main

import (
	"log"
	"net/http"

	"github.com/agustin-carnevale/hexagonal-go/internal/user/adapters/drivens/memory"
	http_adapter "github.com/agustin-carnevale/hexagonal-go/internal/user/adapters/drivers/http"
	"github.com/agustin-carnevale/hexagonal-go/internal/user/usecases"
)

// entry point
func main() {
	// Driven adapter: memory repo
	repo := memory.New()

	// Usecase: core app logic
	userService := usecases.NewUserService(repo)

	// Driver adapter: HTTP layer
	handler := http_adapter.NewHandler(userService)

	// HTTP server
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println("🚀 Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
