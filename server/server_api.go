package server

import "github.com/gin-gonic/gin"

func setApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/alive", handleApiAlive)
		api.GET("/newToken", handleApiNewToken)
	}
}
