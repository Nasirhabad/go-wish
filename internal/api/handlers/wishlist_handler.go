package handlers

import (
	"go-wishlist-api/internal/domain"
	"go-wishlist-api/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistHandler struct {
	usecase usecase.WishlistUsecase
}

func NewWishlistHandler(usecase usecase.WishlistUsecase) *WishlistHandler {
	return &WishlistHandler{usecase: usecase}
}

func (h *WishlistHandler) GetAll(c echo.Context) error {
	wishlists, err := h.usecase.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"data": wishlists})
}

func (h *WishlistHandler) Create(c echo.Context) error {
	var input domain.Wishlist
	if err := c.Bind(&input); err != nil {
		return err
	}
	wishlist, err := h.usecase.Create(input)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{"data": wishlist})
}
