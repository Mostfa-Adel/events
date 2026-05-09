package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-events/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("events", getEvents)
	server.GET("go-events/:id", getEvent)

	authRoutes := server.Group("/").Use(middlewares.Authenticate)

	authRoutes.POST("events", createEvent)
	authRoutes.DELETE("go-events/:id", deleteEvent)
	authRoutes.PUT("go-events/:id", updateEvent)
	authRoutes.POST("go-events/:id/register", registerUserInEvent)
	authRoutes.DELETE("go-events/:id/register", cancelUserInEvent)

	server.POST("signup", signup)
	server.POST("login", login)

}
