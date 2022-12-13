package controllers

import (
	"fmt"
	"net/http"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/middleware"
	"projects-subscribeme-backend/pkg/models"
	subjectDTO "projects-subscribeme-backend/pkg/utils/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

var subjectModel models.SubjectModel

func CreateSubjectController(model models.SubjectModel) {
	subjectModel = model
}

// GET /mata-kuliah
func GetAllSubjects(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	response, err := subjectModel.GetAll()
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

	subject, err := subjectModel.FindByID(id)
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
	var subjectData subjectDTO.SubjectRequest

	err = c.ShouldBindJSON(&subjectData)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting json to model")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	newSubject := modelsConfig.Subject{
		Title: subjectData.Title,
		Term:  subjectData.Term,
		Major: subjectData.Major,
	}

	err = subjectModel.Create(newSubject, subjectData.Classes)
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

	err = subjectModel.DeleteByID(id)
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
	var subjectData subjectDTO.SubjectRequest
	err = c.ShouldBindJSON(&subjectData)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting data to request model")
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	subjectData.ID = id
	err = subjectModel.UpdateByID(id, subjectData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Update success."})
}
