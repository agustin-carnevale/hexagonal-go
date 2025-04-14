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
	memoryRepository := memory.New()

	// Usecase: core app logic
	userService := usecases.NewUserService(memoryRepository)

	// Driver adapter: HTTP layer
	httpHandler := http_adapter.NewHandler(userService)

	// HTTP server
	mux := http.NewServeMux()
	httpHandler.RegisterRoutes(mux)

	log.Println("🚀 Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
