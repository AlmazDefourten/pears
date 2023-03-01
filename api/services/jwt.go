package services

import (
	"fmt"
	"github.com/AlmazDefourten/goapp/models"
	"github.com/AlmazDefourten/goapp/models/container_models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// JWTService struct of service for authorization
type JWTService struct {
	Container  *container_models.Container
	JWTService models.IJWTService
}

func NewJWTService(container *container_models.Container) *JWTService {
	return &JWTService{
		Container: container,
	}
}

// Signin - authorization method
func (jwtService *JWTService) SignIn(username string) (*models.Tokens, error) {
	//SIGNING_KEY - key for signing token
	SIGNING_KEY := jwtService.Container.ConfigProvider.GetString("jwt.signing_key")

	//create access token
	access_token_time := jwtService.Container.ConfigProvider.GetString("jwt.access_token_time")
	access_add_time, err := time.ParseDuration(access_token_time)
	if err != nil {
		return nil, err
	}
	a_claims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(access_add_time).Unix(),
		},
		Username: username,
	}
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, a_claims)
	access, err := access_token.SignedString([]byte(SIGNING_KEY))
	if err != nil {
		return nil, err
	}

	//create refresh token
	refresh_token_time := jwtService.Container.ConfigProvider.GetString("jwt.refresh_token_time")
	refresh_add_time, err := time.ParseDuration(refresh_token_time)
	if err != nil {
		return nil, err
	}
	refresh_claims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refresh_add_time).Unix(),
		},
		Username: username,
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)
	refresh, err := refresh_token.SignedString([]byte(SIGNING_KEY))

	if err != nil {
		return nil, err
	}
	return &models.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// method for refreshing tokens
func (jwtService *JWTService) RefreshTokens(refresh_token string) (*models.Tokens, error) {
	SIGNING_KEY := jwtService.Container.ConfigProvider.GetString("jwt.signing_key")

	// Parse takes the token string and a function for looking up the key
	token, err := jwt.Parse(refresh_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SIGNING_KEY, nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		username := claims.Username
		newTokenPair, err := jwtService.JWTService.SignIn(username)
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
