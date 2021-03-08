package controller

import (
	"net/http"

	"github.com/alelaca/booking-backend/cmd/entities"
	"github.com/gin-gonic/gin"
)

func getServices() gin.HandlerFunc {
	return func(c *gin.Context) {
		services := []entities.Service{
			{ID: 1, Name: "Braids", Duration: 30},
			{ID: 2, Name: "Abuja", Duration: 30},
			{ID: 3, Name: "Blowdry", Duration: 45},
			{ID: 4, Name: "Haircut", Duration: 40},
			{ID: 5, Name: "Relaxer", Duration: 60},
			{ID: 6, Name: "Shampoo", Duration: 15},
			{ID: 7, Name: "Manicure", Duration: 20},
		}

		c.JSON(http.StatusOK, services)
	}
}
