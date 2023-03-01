package models

type IAuthService interface {
	Registration(user *User) bool
	CheckIfUserExist(login string) bool
	Authorization(login string, password string) (bool, *Tokens)
	AuthCheck(token string) (bool, string)
}
