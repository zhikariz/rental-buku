package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ResetPasswordInput struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"password" binding:"required"`
}
