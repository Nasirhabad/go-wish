package usecase

import "github.com/LENOVO/go-wishlist-api/internal/domain"

// WishlistRepository defines the interface for wishlist repository operations.
type WishlistRepository interface {
	GetAll() ([]domain.Wishlist, error)
	Create(domain.Wishlist) (domain.Wishlist, error)
}

// WishlistUsecase defines the use case methods for managing wishlists.
type WishlistUsecase interface {
	GetAll() ([]domain.Wishlist, error)              // GetAll retrieves all wishlists.
	Create(domain.Wishlist) (domain.Wishlist, error) // Create adds a new wishlist.
}

type wishlistUsecase struct {
	repo WishlistRepository
}

// NewWishlistUsecase creates a new instance of WishlistUsecase.
func NewWishlistUsecase(repo WishlistRepository) WishlistUsecase {
	return &wishlistUsecase{repo: repo}
}

// GetAll retrieves all wishlists from the repository.
func (u *wishlistUsecase) GetAll() ([]domain.Wishlist, error) {
	return u.repo.GetAll()
}

// Create adds a new wishlist to the repository.
func (u *wishlistUsecase) Create(wishlist domain.Wishlist) (domain.Wishlist, error) {
	return u.repo.Create(wishlist)
}
