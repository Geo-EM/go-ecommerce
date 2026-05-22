package domain

import "time"

type BankAccount struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserId uint `json:"user_id" gorm:"not null"`

	BankAccountNumber uint   `json:"bank_account_number" gorm:"index;unique;not null"`
	SwiftCode         string `json:"swift_code"`
	PaymentType       string `json:"payment_type"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
