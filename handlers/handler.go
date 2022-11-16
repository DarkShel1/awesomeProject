package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/")
		api.GET("/")
		api.GET("/:id")
		api.PUT("/:id")
		api.DELETE("/:id")

		reservations := api.Group("/reserve")
		{
			reservations.POST("/")
			reservations.GET("/")
			reservations.GET("/:user_id")
			reservations.PUT("/:user_id")
			reservations.DELETE("/:user_id")
		}
	}

	return router
}
