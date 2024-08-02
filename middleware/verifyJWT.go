package middleware

import (
	"mypackages/controllers/actions"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var seccretKey, _ = os.LookupEnv("SECRET_KEY")

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Header["Access-Token"])<1{
			c.AbortWithStatusJSON(401, gin.H{"error": "message"})
			return
		}

		token, err := jwt.Parse(c.Request.Header["Access-Token"][0], func(token *jwt.Token) (interface{}, error) {
			return []byte(seccretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "message"})
			return
		}

		if !token.Valid{
			c.AbortWithStatusJSON(401, gin.H{"error": "message"})
			return 
		}
		actions.ParseJWT(c.Request.Header["Access-Token"][0])
		c.Next()
	}
}

