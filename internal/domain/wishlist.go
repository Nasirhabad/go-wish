package domain

import "time"

type Wishlist struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Title      string     `json:"title" gorm:"not null"`            // Ensure Title is not null
	IsAchieved bool       `json:"is_achieved" gorm:"default:false"` // Default value for IsAchieved
}
