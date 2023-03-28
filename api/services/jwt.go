package services

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/infrastructure/loggerinstance"
	"github.com/AlmazDefourten/goapp/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
	"time"
)

// JWTService struct of service for authorization
type JWTService struct {
}

func NewJWTService() *JWTService {
	return &JWTService{}
}

// Signin - authorization method
func (jwtService *JWTService) SignIn(username string) (*models.Tokens, error) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		loggerinstance.GlobalLogger.Error(err)
		panic(err)
	}

	var c models.Configurator
	err = container.Resolve(&c)
	if err != nil {
		loggerinstance.ServiceLogger.Error(err)
		return nil, err
	}
	//SIGNING_KEY - key for signing token
	SigningKey := c.GetString("jwt.signing_key")

	//create access token
	accessTokenTime := c.GetString("jwt.access_token_time")
	accessAddTime, err := time.ParseDuration(accessTokenTime)
	if err != nil {
		return nil, err
	}
	aClaims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessAddTime).Unix(),
		},
		Username: username,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	access, err := accessToken.SignedString([]byte(SigningKey))
	if err != nil {
		return nil, err
	}

	//create refresh token
	refreshTokenTime := c.GetString("jwt.refresh_token_time")
	refreshAddTime, err := time.ParseDuration(refreshTokenTime)
	if err != nil {
		return nil, err
	}
	refreshClaims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshAddTime).Unix(),
		},
		Username: username,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh, err := refreshToken.SignedString([]byte(SigningKey))

	var user models.User
	db.Model(&user).Where("login = ?", username).Update("refresh_token", refresh)

	if err != nil {
		return nil, err
	}
	return &models.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// method for refreshing tokens
func (jwtService *JWTService) ValidateAndRefreshTokens(refresh_token string) (*models.Tokens, error) {
	var c models.Configurator
	err := container.Resolve(&c)
	if err != nil {
		loggerinstance.ServiceLogger.Error(err)
		return nil, err
	}

	SigningKey := c.GetString("jwt.signing_key")

	// Parse takes the token string and a function for looking up the key
	token, err := jwt.Parse(refresh_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SigningKey, nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		username := claims.Username
		newTokenPair, err := jwtService.SignIn(username)
		if err != nil {
			return nil, err
		}

		return newTokenPair, err
	}

	return nil, err
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
