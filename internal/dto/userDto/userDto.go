package userDto

type LoginUserDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type RegisterUserDto struct {
	LoginUserDto
	Phone string `json:"phone"`
}

type VerifyCodeUserDto struct {
	VerificationCode uint `json:"verification_code"`
}

type SellerDto struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	BankAccountNumber string `json:"bank_account_number"`
	SwiftCode         string `json:"swift_code"`
	PaymentMethod     string `json:"payment_method"`
}
