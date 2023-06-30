package routes

import (
	"github.com/gin-gonic/gin"
	"sber_ds/handlers"
)

func SetupRouters(router *gin.Engine) {
	router.GET("/api/v1/get_report", handlers.GetReport)
	router.POST("/api/v1/set_report", handlers.SetReport)
	router.GET("/api/v1/get_observation_time", handlers.GetObservationTime)
}
