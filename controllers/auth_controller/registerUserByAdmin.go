package auth_controller

import (
	"kosei-jwt/models"
	"kosei-jwt/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (ac *AuthController) RegisterUserByAdmin(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	if currentUser.Role == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "user not admin!"})
		return
	}
	var payload models.SignUpInputAdmin

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := ctx.ShouldBindUri(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := models.ValidateUserAdmin(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	filename := ""
	if payload.Photo != nil {
		if err := ctx.SaveUploadedFile(payload.Photo, "uploads/"+payload.Photo.Filename); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}
		filename = payload.Photo.Filename
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      payload.Role,
		Photo:     filename,
		TypeUser:  payload.TypeUser,
		Deleted:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err = ac.Auth.GetUserByName(payload.Name)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "name already exists"})
		return
	}

	err = ac.Auth.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "create user error !"})
		return
	}

	userResponse := &models.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Photo:     newUser.Photo,
		Role:      newUser.Role,
		Type:      newUser.TypeUser,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
