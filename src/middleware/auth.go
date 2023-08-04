package middleware

import (
	"net/http"

	"github.com/Yusup1907/banking-api/src/utils"
	"github.com/gin-gonic/gin"
)

func RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Baca cookie dengan nama "token"
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Missing token cookie",
			})
			return
		}

		// Ambil value dari cookie sebagai tokenString
		tokenString := cookie.Value

		// check token kosong
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Empty token",
			})
			return
		}

		// check verify token
		_, err = utils.VerifyAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Invalid token",
			})
			return
		}

		c.Next()
	}
}
