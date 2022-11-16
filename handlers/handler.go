package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.createBalance)
		api.GET("/", h.getAllBalances)
		api.GET("/:id", h.getBalanceById)
		api.PUT("/:id", h.updateBalance)
		api.DELETE("/:id", h.deleteBalance)

		reservations := api.Group("/reserve")
		{
			reservations.POST("/", h.createReservation)
			reservations.GET("/", h.getAllReservations)
			reservations.GET("/:user_id", h.getReservationById)
			reservations.PUT("/:user_id", h.updateReservation)
			reservations.DELETE("/:user_id", h.deleteReservation)
		}
	}

	return router
}
