package domain

import "time"

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"index;not null"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`

	// TODO: change the float64 type to a more precise type to avoid rounding issues in financial calculations
	Price float64 `json:"price" gorm:"not null"`
	Stock uint    `json:"stock" gorm:"not null"`

	SellerId uint `json:"seller_id" gorm:"index;not null"`
	Seller   User `json:"seller" gorm:"foreignKey:SellerId"`

	CategoryId uint     `json:"category_id" gorm:"index;not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
