package controllers

import (
	"fmt"
	"log"
	"net/http"
	modelsConfig "projects-subscribeme-backend/pkg/config/models"
	"projects-subscribeme-backend/pkg/middleware"
	"projects-subscribeme-backend/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var eventService service.EventService

func CreateEventController(service service.EventService) {
	eventService = service
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

	err = eventService.Create(eventRequest)
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

	event, err := eventService.FindByID(id)
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

	err = eventService.UpdateByID(id, eventRequest)
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

	err = eventService.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Delete success."})
}
