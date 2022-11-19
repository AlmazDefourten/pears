package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/auth/pkg/parser"
)

struct User {
	id    int   `json:"id"`
	login string `json:"login"`
	password string `json:"password"` 
	mail  string `json:"mail"`
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

//создать токен по логину паролю
func (a *Authorizer) SignIn(username, password string) (string, error) {
	//проверить логин пароль
	//подключение к бд
	dsn := "host=localhost user=postgres password=mypas dbname=pears port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//получить пользователя из бд
	var user User
	db.Where("login = ?", username).First(&user)
	if user.password != password {
		return "", fmt.Errorf("invalid username or password")
	}

	//создать токен
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssueAt:   jwt.At(time.Now()),
		},
		Username: username,
	}

	//вернуть токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.signingkey)
}

//расшифровка токена
func ParseToken(accessToken string, signingKey []byte) (string, error){
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	}

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", fmt.Errorf("invalid token")
}

//middleware для проверки токена
func Middleware (c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err := parser.ParseToken(headerParts[1], SIGNING_KEY)
	if err != nil {
		status := http.StatusBadRequest
		if err == jwt.ErrInvalidToken {
			status = http.StatusUnauthorized
		}
		c.AbortWithStatus(status)
		return
	}
}











