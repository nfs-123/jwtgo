package controllers

import (
	auth "jwtgo/helpers"
	"jwtgo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Username string `json:"username"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.Bind(&request); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	tokenstring, err := auth.GenerateJWT(user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"token":    tokenstring,
		"username": request.Username,
	})
}
