package controller

import (
	"fmt"
	"net/http"
	"projects-subscribeme-backend/config/models"
	"projects-subscribeme-backend/middleware"
	"projects-subscribeme-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var subjectService service.SubjectService

func CreateSubjectController(service service.SubjectService) {
	subjectService = service
}

// GET /mata-kuliah
func GetAllSubjects(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	response, err := subjectService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// GET /mata-kuliah/:id
func GetSubjectByID(c *gin.Context) {
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

	subject, err := subjectService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": subject})

}

// POST /mata-kuliah
func CreateSubject(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	var subjectRequest models.SubjectRequest

	err = c.ShouldBindJSON(&subjectRequest)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting json to model")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	err = subjectService.Create(subjectRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Success"})
}

// DELETE /mata-kuliah/:id
func DeleteSubjectByID(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	err = subjectService.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Delete success."})
}

// POST /mata-kuliah/:id
func UpdateSubjectByID(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting string to int")
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	var subjectRequest models.SubjectRequest
	err = c.ShouldBindJSON(&subjectRequest)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting data to request model")
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	err = subjectService.UpdateByID(id, subjectRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Update success."})
}
