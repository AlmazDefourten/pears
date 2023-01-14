package jwt

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	iris "github.com/kataras/iris/v12"
	"..\models"
	"..\infrastructure\container"
)
// JWTService struct of service for authorization
type JWTService struct {
	Container *models.Container
}

//Signin - authorization method
func (a *Authorizer) SignIn(username, password string) (string, error) *JWTService {
	//check if user exists and password is correct
	//connect to db
	db := container.NewConnection()

	//get user from db
	var user models.User
	db.Where("login = ?", username).First(&user)
	if user.password != password {
		return "", fmt.Errorf("invalid username or password")
	}

	//create token
	claims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssueAt:   jwt.At(time.Now()),
		},
		Username: username,
	}

	//return token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.signingkey)
}

//ParseToken - parse token method
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

//CheckToken - method for checking token
func CheckToken (c iris.Context) {
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











