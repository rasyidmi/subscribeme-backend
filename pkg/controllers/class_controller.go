package controllers

import (
	"fmt"
	"net/http"
	"projects-subscribeme-backend/pkg/service"
	"projects-subscribeme-backend/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var classService service.ClassService

func CreateClassController(service service.ClassService) {
	classService = service
}

// GET /kelas/:id
func GetClassByID(c *gin.Context) {
	// err := middleware.AuthMiddleware(c)
	// if err != nil {
	// 	return
	// }
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"data": "Error on converting string to int"})
		return
	}

	class, err := classService.GetClassByID(id)
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
