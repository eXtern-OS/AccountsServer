package server

import "github.com/gin-gonic/gin"

func setApi(r *gin.Engine) {
	r.GET("/alive", handleApiAlive)
	api := r.Group("/api")
	{
		api.GET("/newToken", handleApiNewToken)
		api.GET("/register", handleApiRegister)
	}
}
