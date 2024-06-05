package routes

import (
	"Ntt_DATA/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/exoplanets", controller.AddExoplanet)
	r.GET("/exoplanets", controller.ListExoplanets)
	r.GET("/exoplanets/:id", controller.GetExoplanetByID)
	r.PUT("/exoplanets/:id", controller.UpdateExoplanet)
	r.DELETE("/exoplanets/:id", controller.DeleteExoplanet)
	r.GET("/exoplanets/:id/fuel-estimation", controller.FuelEstimation)

	return r
}
