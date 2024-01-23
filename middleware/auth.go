package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"vs45tech.com/event/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	fmt.Print(token)
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Un Authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token", "error": err.Error()})
		return
	}
	context.Set("userId", userId)
	context.Next()

}
