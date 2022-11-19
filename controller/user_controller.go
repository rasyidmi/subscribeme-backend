package controller

import (
	"fmt"
	"net/http"
	"projects-subscribeme-backend/models"
	"projects-subscribeme-backend/service"

	"github.com/gin-gonic/gin"
)

var userService service.UserService

func CreateUserController(service service.UserService) {
	userService = service
}

func CreateUser(c *gin.Context) {
	var user models.UserRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting json to model")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	err = userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Success"})
}

func FindUserByID(c *gin.Context) {
	user, err := userService.FindByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
