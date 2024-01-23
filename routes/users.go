package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vs45tech.com/event/models"
	"vs45tech.com/event/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		//fmt.Println(event)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request", "Error": err})
		return
	}
	id, err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create User", "Error": err})
		return
	}
	user.ID = id
	context.JSON(http.StatusCreated, gin.H{"message": "User Created !"})
}
func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		//fmt.Println(event)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request", "Error": err})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not Login User", "Error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Generate Token", "Error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Login Successful!", "token": token})
}
