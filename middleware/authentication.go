package middleware

import (
	"ecommerce/helpers"
	"ecommerce/services"
	"ecommerce/utils"
	"errors"
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
	claim := jwt.MapClaims{
		"user_id":   userID,
		"ExpiresAt": expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, expirationTime, err
	}
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
		response := helpers.RESPONSE_AUTH
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			authHeader = arrayToken[1]
		}
		token, err := authService.ValidateToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//fmt.Println(claim, "claimnya ")
		//fmt.Println("xp nya", claim["ExpiresAt"].(float64))
		idUser, _ := claim["user_id"].(string)
		expNew, _ := claim["ExpiresAt"].(float64)
		exp := int(expNew)
		if exp < int(time.Now().Local().Unix()) {
			responseExp := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "token expired")
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseExp)
			return
		}
		user, err := userService.FindUserByID(idUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("current_user", user)
		c.Next()
	}
}
func AuthCustomer(authService ServiceAuth, userService services.UserServices) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		response := helpers.ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "unauthorized")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//fmt.Println(claim, "claimnya ")
		//fmt.Println("xp nya", claim["ExpiresAt"].(float64))
		idUser, _ := claim["user_id"].(string)
		expNew, _ := claim["ExpiresAt"].(float64)
		exp := int(expNew)
		if exp < int(time.Now().Local().Unix()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		user, err := userService.FindUserByID(idUser)
		if user.Role != "CUSTOMER" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("current_user", user)
		c.Next()
	}
}
