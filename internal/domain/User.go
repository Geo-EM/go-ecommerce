package domain

import "time"

type User struct {
	ID                     uint      `json:"id" gorm:"primaryKey"`
	FirstName              string    `json:"first_name" gorm:"not null"`
	LastName               string    `json:"last_name" gorm:"not null"`
	Email                  string    `json:"email" gorm:"index;unique;not null"`
	Phone                  string    `json:"phone" gorm:"unique"`
	Password               string    `json:"password" gorm:"not null"`
	VerificationCode       uint      `json:"verification_code" gorm:"default:null"`
	VerificationCodeExpiry time.Time `json:"verification_code_expiry" gorm:"default:null"`
	Verified               bool      `json:"verified" gorm:"default:false"`
	UserType               string    `json:"user_type" gorm:"default:'buyer'"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
