package routes

import (
	"github.com/gin-gonic/gin"
	"vs45tech.com/event/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	authenticated := server.Group("/", middleware.Authenticate)
	authenticated.POST("/events", CreateEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deletEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
