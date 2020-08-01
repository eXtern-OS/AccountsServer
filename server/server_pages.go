package server

import "github.com/gin-gonic/gin"

func setPages(r *gin.Engine) {
	r.GET("/", handleMainGet)

	r.POST("/", handleMainPost)

	r.GET("/login", handleLoginGet)

	r.POST("/login", handleLoginPost)

	r.GET("/register", handleRegisterGet)

	r.POST("/register", handleRegisterPost)
}
