package services

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/infra/logger_instance"
	"github.com/AlmazDefourten/goapp/models/user_models"
	"github.com/AlmazDefourten/goapp/models/util_adapters"
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
func (jwtService *JWTService) SignIn(username string) (*user_models.Tokens, error) {
	var db gorm.DB
	err := container.Resolve(&db)
	if err != nil {
		logger_instance.GlobalLogger.Error(err)
		panic(err)
	}

	var c util_adapters.Configurator
	err = container.Resolve(&c)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
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
	aClaims := user_models.Claims{
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
	refreshClaims := user_models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshAddTime).Unix(),
		},
		Username: username,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh, err := refreshToken.SignedString([]byte(SigningKey))

	var user user_models.User
	db.Model(&user).Where("login = ?", username).Update("refresh_token", refresh)

	if err != nil {
		return nil, err
	}
	return &user_models.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// method for refreshing tokens
func (jwtService *JWTService) ValidateAndRefreshTokens(refresh_token string) (*user_models.Tokens, error) {
	var c util_adapters.Configurator
	err := container.Resolve(&c)
	if err != nil {
		logger_instance.ServiceLogger.Error(err)
		return nil, err
	}

	SigningKey := c.GetString("jwt.signing_key")

	if refresh_token == "" {
		return nil, fmt.Errorf("empty token")
	}

	// Parse takes the token string and a function for looking up the key
	token, err := jwt.Parse(refresh_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SigningKey, nil
	})

	if claims, ok := token.Claims.(*user_models.Claims); ok && token.Valid {
		username := claims.Username
		newTokenPair, err := jwtService.SignIn(username)
		if err != nil {
			return nil, fmt.Errorf("token refresh error: %v", err)
		}

		return newTokenPair, err
	}

	return nil, fmt.Errorf("invalid token")
}

// ParseToken - parse token method
func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &user_models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*user_models.Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", fmt.Errorf("invalid token")
}
