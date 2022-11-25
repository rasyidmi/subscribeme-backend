package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	configModels "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/middleware"
	"projects-subscribeme-backend/pkg/models"
	"projects-subscribeme-backend/pkg/utils"
	userDTO "projects-subscribeme-backend/pkg/utils/dto"
)

var userModel models.UserModel

func CreateUserController(model models.UserModel) {
	userModel = model
}

func CreateUser(c *gin.Context) {
	var userData userDTO.UserRequest
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting json to model")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on hasing the password.")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	newUser := configModels.User{
		Email:    userData.Email,
		Password: hashedPassword,
		Name:     userData.Name,
		Role:     userData.Role,
		Avatar:   userData.Avatar,
	}
	err = userModel.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Success"})
}

func Login(c *gin.Context) {
	var userData userDTO.UserRequest
	err := c.ShouldBindJSON(&userData)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting json to model")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	// Find the user.
	user, err := userModel.FindByEmail(userData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	userResponse := userDTO.UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Name:   user.Name,
		Role:   user.Role,
		Avatar: user.Avatar,
	}
	// Compare password
	isCorrect := utils.ComparePassword(userData.Password, user.Password)
	if !isCorrect {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Password incorrect."})
		return
	}
	// Generate access token.
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	// Generate refresh token.
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"userData": userResponse, "refreshToken": refreshToken, "accessToken": accessToken}})
}

func AutoLogin(c *gin.Context) {
	token := middleware.ExtractToken(c.GetHeader("Authorization"))
	payload, err := utils.VerifyAccessToken(token)
	if err != nil && strings.Contains(err.Error(), "expired") {
		c.JSON(http.StatusUnauthorized, gin.H{"data": err.Error()})
		return
	}
	user, err := userModel.FindByID(payload.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	userResponse := userDTO.UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Name:   user.Name,
		Role:   user.Role,
		Avatar: user.Avatar,
	}
	// Generate new access token.
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	// Generate new refressh token.
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"userData": userResponse, "refreshToken": refreshToken, "accessToken": accessToken}})
}

func RefreshToken(c *gin.Context) {
	refreshToken := middleware.ExtractToken(c.GetHeader("Authorization"))
	payload, err := utils.VerifyRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"data": err.Error()})
	}
	user, err := userModel.FindByID(payload.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	// Generate new access token.
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	// Generate new refressh token.
	newRefreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	userResponse := userDTO.UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Name:   user.Name,
		Role:   user.Role,
		Avatar: user.Avatar,
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"userData": userResponse, "refreshToken": newRefreshToken, "accessToken": accessToken}})
}

func FindUserByID(c *gin.Context) {
	user, err := userModel.FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
