package router

import (
	"github.com/gin-gonic/gin"
	"github.com/karelpelcak/foodApI/ent"
	"github.com/karelpelcak/foodApI/handlers"
)

func SetupRouter(client *ent.Client) *gin.Engine {
	router := gin.Default()

	r := router.Group("/api/v1")
	authR := r.Group("/auth")
	{
		r.GET("/ping", handlers.Ping)
		authR.GET("/me", handlers.Me)
	}
	return router
}
