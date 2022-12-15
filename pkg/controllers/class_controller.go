package controllers

import (
	"fmt"
	"net/http"
	"projects-subscribeme-backend/pkg/middleware"
	"projects-subscribeme-backend/pkg/models"
	"projects-subscribeme-backend/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var classModel models.ClassModel

func CreateClassController(model models.ClassModel) {
	classModel = model
}

// GET /kelas/:id
func GetClassByID(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"data": "Error on converting string to int"})
		return
	}
	userId := c.Request.Header.Get("userId")

	class, err := classModel.GetClassByID(id, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	} else if class.ID == 0 {
		err = utils.DataNotFound{}
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": class})
}

func Subscribe(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	classId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"data": "Error on converting string to int"})
		return
	}

	userId := c.Request.Header.Get("userId")
	err = classModel.Subscribe(classId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Successfully subscribe the class."})
}
