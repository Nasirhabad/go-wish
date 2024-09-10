package repository

import (
	"github.com/LENOVO/go-wishlist-api/internal/domain"

	"gorm.io/gorm"
)

type WishlistRepository interface {
	GetAll() ([]domain.Wishlist, error)
	Create(domain.Wishlist) (domain.Wishlist, error)
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	return &wishlistRepository{db: db}
}

func (r *wishlistRepository) GetAll() ([]domain.Wishlist, error) {
	var wishlists []domain.Wishlist
	if err := r.db.Find(&wishlists).Error; err != nil {
		return nil, err
	}
	return wishlists, nil
}

func (r *wishlistRepository) Create(wishlist domain.Wishlist) (domain.Wishlist, error) {
	if err := r.db.Create(&wishlist).Error; err != nil {
		return domain.Wishlist{}, err
	}
	return wishlist, nil
}
