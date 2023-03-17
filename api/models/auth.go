package models

type IAuthService interface {
	Registration(user *User) bool
	CheckIfUserExist(login string) bool
	Authorization(login string, password string) (bool, *Tokens)
	AuthCheck(token string) (bool, string)
	RefreshCheck(token string) (bool, *Tokens)
}

// Response is sent as a response with information about the success of the request
type AuthResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}
