package services

import (
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	// iris "github.com/kataras/iris/v12"
	"github.com/AlmazDefourten/goapp/models"
)

// SIGNING_KEY - key for signing token
var SIGNING_KEY string = "secret"

// JWTService struct of service for authorization
type JWTService struct {
	Container *models.Container
}

func NewJWTService(container *models.Container) *JWTService {
	return &JWTService{
		Container: container,
	}
}

// Signin - authorization method
func (jwtService *JWTService) SignIn(username, password string) (string, error) {
	//check if user exists and password is correct
	//get user from db
	var user models.User
	jwtService.Container.AppConnection.Where("login = ?", username).First(&user)
	if user.Password != password {
		return "", fmt.Errorf("invalid username or password")
	}

	//create token
	claims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//return token
	return token.SignedString(SIGNING_KEY)
}

// ParseToken - parse token method
func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", fmt.Errorf("invalid token")
}

//CheckToken - method for checking token
// func CheckToken(c iris.Context) {
// 	authHeader := c.GetHeader("Authorization")
// 	if authHeader == "" {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	headerParts := strings.Split(authHeader, " ")
// 	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	err := parser.ParseToken(headerParts[1], SIGNING_KEY)
// 	if err != nil {
// 		status := http.StatusBadRequest
// 		if err == jwt.ErrInvalidToken {
// 			status = http.StatusUnauthorized
// 		}
// 		c.AbortWithStatus(status)
// 		return
// 	}
// }


