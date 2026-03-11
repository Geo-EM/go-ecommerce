package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	Code      string    `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:'buyer'"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
