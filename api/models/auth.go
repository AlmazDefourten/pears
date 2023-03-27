package models

type IAuthService interface {
	Registration(user *User) (bool, error)
	CheckIfUserExist(login string) (bool, error)
	Authorization(login string, password string) (bool, *Tokens)
	AuthCheck(token string) (bool, string)
}

// Response is sent as a response with information about the success of the request
type AuthResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
