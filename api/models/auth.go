package models

type IAuthService interface {
	Registration(user *User) (bool, string)
	CheckIfUserExist(login string) (bool, error)
	Authorization(login string, password string) (bool, *Tokens, string)
	AuthCheck(token string) (bool, string)
	RefreshCheck(token string) (bool, *Tokens)
}

// Response is sent as a response with information about the success of the request
type AuthResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Access  string `json:"access_token"`
}
