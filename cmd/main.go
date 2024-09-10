package main

import (
	"github.com/LENOVO/go-wishlist-api/database"
	"github.com/LENOVO/go-wishlist-api/internal/api/handlers/routes"
	"github.com/LENOVO/go-wishlist-api/internal/api/server"
	"github.com/LENOVO/go-wishlist-api/internal/repository"
	"github.com/LENOVO/go-wishlist-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()

	e := echo.New()

	// Initialize repositories
	wishlistRepo := repository.NewWishlistRepository(database.DB)

	// Initialize use cases
	wishlistUsecase := usecase.NewWishlistUsecase(wishlistRepo)

	// Initialize handlers with use cases
	routes.RegisterRoutes(e, wishlistUsecase)

	// Start the server
	server.StartServer(e)
}
