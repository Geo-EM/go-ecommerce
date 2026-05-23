package domain

import "time"

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"index;not null"`
	Description string `json:"description"`
	// TODO: change the float64 type to a more precise type to avoid rounding issues in financial calculations
	Price    float64 `json:"price" gorm:"not null"`
	ImageUrl string  `json:"image_url"`
	Stock    uint    `json:"stock" gorm:"not null"`

	SellerId string `json:"seller_id" gorm:"index;not null"`
	Seller   User   `json:"seller" gorm:"foreignKey:SellerId"`

	CategoryId string   `json:"category_id" gorm:"index;not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
