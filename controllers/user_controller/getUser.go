package usercontroller

import (
	"kosei-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:          currentUser.ID,
		Name:        currentUser.Name,
		Email:       currentUser.Email,
		Photo:       currentUser.Photo,
		Role:        currentUser.Role,
		District:    currentUser.District,
		DateOfBirth: currentUser.DateOfBirth,
		Address:     currentUser.Address,
		City:        currentUser.City,
		FullName:    currentUser.FullName,
		Type:        currentUser.TypeUser,
		CreatedAt:   currentUser.CreatedAt,
		UpdatedAt:   currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
