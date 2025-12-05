package dto

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Something went wrong"`
}

type SuccessMessageResponse struct {
	Message string `json:"message" example:"OK"`
}

const (
	MsgUnauthorized        = "Something went wrong while authenticating user!"
	MsgInvalidPassword     = "Invalid password!"
	MsgForbidden           = "Forbidden"
	MsgBadRequest          = "Something went wrong while processing your request!"
	MsgUserAlreadyLoggedIn = "User already logged in!"
	MsgRegistered          = "Registered successfully! You're now logged in."
	MsgLoggedIn            = "Logged in successfully"
)
