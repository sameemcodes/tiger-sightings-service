package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")
	fmt.Println("tokenString", tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("token", token)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	fmt.Println("outer token", token)

	if err != nil {
		fmt.Println("token err", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Token Login Attempt!! Please Try Again",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims ", claims["sub"])

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("id", claims["sub"])

		c.Next()
	}

}
