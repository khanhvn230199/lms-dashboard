package usercontroller

import (
	"kosei-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetMeByUserAdmin(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	if currentUser.Role == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "user not admin!"})
		return
	}
	var user models.User
	userResponse := make([]models.UserResponse, 0)
	result := uc.DB.Model(&user).Find(&userResponse)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "get user error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
