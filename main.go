package main

// main
// controller
// service
// repository
// db

import (
	"net/http"
	"projects-subscribeme-backend/config"
	"projects-subscribeme-backend/routers"
	"time"

	timeout "github.com/s-wijaya/gin-timeout"

	"github.com/gin-gonic/gin"
)

func main() {
	config.SetupDatabase()
	// firebaseAuth := config.SetupFirebase()

	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// router.Use(func(c *gin.Context) {
	// 	c.Set("firebaseAuth", firebaseAuth)
	// })
	router.Use(timeout.TimeoutHandler(5*time.Second, http.StatusRequestTimeout, gin.H{"data": "Request Timeout"}))

	v1Api := router.Group("/api/v1")

	routers.User(v1Api.Group("/user"))

	routers.Subject(v1Api.Group("/subject"))

	routers.Class(v1Api.Group("/class"))

	routers.Event(v1Api.Group("/event"))

	router.Run()
}
