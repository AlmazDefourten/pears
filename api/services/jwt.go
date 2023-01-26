package services

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/models/container_models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	// iris "github.com/kataras/iris/v12"
	"github.com/AlmazDefourten/goapp/models"
)

// JWTService struct of service for authorization
type JWTService struct {
	Container *container_models.Container
}

func NewJWTService(container *container_models.Container) *JWTService {
	return &JWTService{
		Container: container,
	}
}

// Signin - authorization method
func (jwtService *JWTService) SignIn(username string) (string, error) {
	//create token
	token_time := jwtService.Container.ConfigProvider.GetString("jwt.token_time")
	claims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.ParseDuration(token_time)).Unix(),
		},
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//SIGNING_KEY - key for signing token
	SIGNING_KEY := jwtService.Container.ConfigProvider.GetString("jwt.signing_key")
	//return token
	return token.SignedString([]byte(SIGNING_KEY))
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
