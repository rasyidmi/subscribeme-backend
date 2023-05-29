package main

import (
	"projects-subscribeme-backend/initializers"
	"projects-subscribeme-backend/routers"
)

// main
// controller
// service
// repository
// db

// func main() {
// 	config.SetupDatabase()
// 	// firebaseAuth := config.SetupFirebase()

// 	router := gin.New()
// 	router.SetTrustedProxies([]string{"127.0.0.1"})

// 	// router.Use(func(c *gin.Context) {
// 	// 	c.Set("firebaseAuth", firebaseAuth)
// 	// })
// 	router.Use(timeout.TimeoutHandler(5*time.Second, http.StatusRequestTimeout, gin.H{"data": "Request Timeout"}))

// 	v1Api := router.Group("/api/v1")

// 	routers.User(v1Api.Group("/user"))

// 	routers.Subject(v1Api.Group("/subject"))

// 	routers.Class(v1Api.Group("/class"))

// 	routers.Event(v1Api.Group("/event"))

// 	router.Run()

// }

func main() {
	initializers.Setup()
	router := routers.RouterSetup()
	router.Run()
}
