package web

import (
	"github.com/eXtern-OS/AMS"
	"github.com/gin-gonic/gin"
)

func RenderIndex(uid string) gin.H {
	acc := AMS.GetUserByID(uid)
	return gin.H{
		"avatar_url":      acc.AvatarURL,
		"name":            acc.Name,
		"registered_date": acc.Registered,
		"login":           acc.Login,
		"email":           acc.Email,
		"website":         acc.Website,
	}
}
