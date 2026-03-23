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
