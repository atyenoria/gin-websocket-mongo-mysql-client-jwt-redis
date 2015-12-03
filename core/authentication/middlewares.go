package authentication

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RequireTokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authBackend := InitJWTAuthenticationBackend()
		fmt.Println(c.Request.Header.Get("Authorization"))
		token, err := jwt.Parse(c.Request.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			} else {
				return authBackend.PublicKey, nil
			}
		})

		fmt.Println(err)
		fmt.Println(token.Valid)
		if err == nil && token.Valid && !authBackend.IsInBlacklist(c.Request.Header.Get("Authorization")) {
			c.Next()
		} else {
			c.Writer.WriteHeader(http.StatusUnauthorized)
		}
	}
}
