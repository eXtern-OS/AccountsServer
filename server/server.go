package server

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine, weburl string) {
	websiteURL = weburl
	setApi(r)
	setPages(r)
}
