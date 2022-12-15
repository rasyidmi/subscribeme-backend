package controllers

import (
	"fmt"
	"log"
	"net/http"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/middleware"
	"projects-subscribeme-backend/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var eventModel models.EventRepository

func CreateEventController(model models.EventRepository) {
	eventModel = model
}

// CREATE /agenda
func CreateEvent(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	var eventRequest modelsConfig.EventDTO

	err = c.ShouldBindJSON(&eventRequest)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting json to model")
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	var classesID []int
	classesID = append(classesID, eventRequest.ClassesID...)

	event := modelsConfig.Event{
		Title:        eventRequest.Title,
		Description:  eventRequest.Description,
		DeadlineDate: eventRequest.DeadlineDate,
		SubjectID:    eventRequest.SubjectID,
		SubjectName:  eventRequest.SubjectName,
	}

	err = eventModel.Create(event, classesID)
	if err != nil {
		log.Println("ERROR HERE")
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Success"})
}

// GET /agenda/id
func GetEventByID(c *gin.Context) {
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

	event, err := eventModel.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": event})
}

// POST /agenda/id
func UpdateEventByID(c *gin.Context) {
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
	var eventRequest modelsConfig.EventDTO
	err = c.ShouldBindJSON(&eventRequest)
	if err != nil {
		fmt.Println("ERROR OCCURED: Error on converting data to request model")
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	data := map[string]interface{}{
		"title":         eventRequest.Title,
		"description":   eventRequest.Description,
		"deadline_date": eventRequest.DeadlineDate,
		"subject_id":    eventRequest.SubjectID,
		"subject_name":  eventRequest.SubjectName,
	}

	err = eventModel.UpdateByID(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Update success."})
}

// DELETE /agenda/id
func DeleteEventByID(c *gin.Context) {
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

	err = eventModel.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Delete success."})
}

func GetTodayDeadline(c *gin.Context) {
	err := middleware.AuthMiddleware(c)
	if err != nil {
		return
	}
	userId := c.Request.Header.Get("userId")
	events, err := eventModel.GetTodayDeadline(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
}
