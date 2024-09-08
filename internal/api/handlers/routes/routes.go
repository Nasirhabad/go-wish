package routes

import (
	"go-wishlist-api/internal/api/handlers"
	"go-wishlist-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, usecase usecase.WishlistUsecase) {
	handler := handlers.NewWishlistHandler(usecase)
	e.GET("/wishlists", handler.GetAll)
	e.POST("/wishlists", handler.Create)
}
