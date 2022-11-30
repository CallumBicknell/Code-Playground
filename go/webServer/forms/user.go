package forms

type UserSignup struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"username" binding:"required"`
	Pin      string `json:"Pin" binding:"required"`
}
