package domain

import "time"

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index;not null"`

	ParentId *uint     `json:"parent_id" gorm:"index"`
	Parent   *Category `json:"parent_category" gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ImageUrl     string    `json:"image_url"`
	Products     []Product `json:"products" gorm:"foreignKey:CategoryId"`
	DisplayOrder int       `json:"display_order"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
