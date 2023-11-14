package usercontroller

import (
	"kosei-jwt/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) DeleteMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var updatedUserPhoto models.User
	result := uc.DB.First(&updatedUserPhoto, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
		return
	}
	now := time.Now()
	userResponse := &models.User{
		Deleted:   int8(1),
		CreatedAt: now,
		UpdatedAt: now,
	}
	uc.DB.Model(&updatedUserPhoto).Updates(userResponse)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
