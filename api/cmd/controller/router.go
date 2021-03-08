package controller

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/shops/1/services", getServices())

	router.Run(":3400")
}
