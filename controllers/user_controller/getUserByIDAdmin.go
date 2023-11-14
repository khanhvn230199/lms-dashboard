package usercontroller

import (
	"kosei-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetMeIDByUserAdmin(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	if currentUser.Role == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "user not admin!"})
		return
	}

	userID := ctx.Param("userID")
	user := uc.User.GetByID(userID)
	if user != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "user not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": user}})
}
