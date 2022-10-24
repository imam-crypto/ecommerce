package middleware

import (
	"ecommerce/helpers"
	"ecommerce/services"
	"ecommerce/utils"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type ServiceAuth interface {
	GenerateToken(userID string) (string, time.Time, error)
	ValidateToken(tokenEncode string) (*jwt.Token, error)
	//RefreshToken(userID string) (string, time.Time, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var config, _ = utils.LoadConfig(".", false)
var SECRET_KEY = []byte(config.SecretKey)

func (s jwtService) GenerateToken(userID string) (string, time.Time, error) {

	expirationTime := time.Now().Add(1 * time.Hour)

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["role"] = "ADMIN"
	claim["StandardClaims"] = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	//fmt.Println("secret keynya", config.SecretKey)
	if err != nil {
		return signedToken, expirationTime, err
	}
	fmt.Println("err nya", err)
	return signedToken, expirationTime, nil
}

func (s jwtService) ValidateToken(tokenEncode string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenEncode, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, nil
	}

	return token, nil
}
func AuthAdmin(authService ServiceAuth, userService services.UserServices) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		exp := int(claim["ExpiresAt"].(float64))
		if exp < int(time.Now().Local().Unix()) {
			response := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "token expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := claim["user_id"].(float64)
		convUserID := fmt.Sprintf("%f", userID)
		user, err := userService.FindUser(convUserID)
		if err != nil {
			response := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if user.Role == "MEMBER" {
			response := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("current_user", user)
		c.Next()
	}

}
