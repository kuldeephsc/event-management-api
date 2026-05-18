package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kuldeephsc/api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)

	server.GET("/events/:id", getEvent)

	// Use Groips to apply authentication middleware to specific routes
	// authenticated := server.Group("/")
	// authenticated.Use(middlewares.AuthenticationMiddleware)

	// authenticated.POST("/events", createEvent)

	// authenticated.PUT("/events/:id", updateEvent)

	// authenticated.DELETE("/events/:id", deleteEvent)

	// authenticated.DELETE("/events/:id/register", registerForEvent)

	//authenticated.DELETE("/events/:id", unregisterForEvent)

	server.POST("/events", middlewares.AuthenticationMiddleware, createEvent)

	server.PUT("/events/:id", middlewares.AuthenticationMiddleware, updateEvent)

	server.DELETE("/events/:id", middlewares.AuthenticationMiddleware, deleteEvent)

	server.POST("/events/:id/register", middlewares.AuthenticationMiddleware, registerForEvent)

	server.DELETE("/events/:id/register", middlewares.AuthenticationMiddleware, unregisterForEvent)

	server.POST("/signup", signup)

	server.POST("/login", login)

}
