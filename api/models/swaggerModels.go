package models

// UserAuthInfo it`s swagger only struct to describe body of queries with auth user
type UserAuthInfo struct {
	Login    string `json:"login" example:"ivan.petrov@mail.ru"`
	Password string `json:"password" example:"mypaSSword123876"`
}
