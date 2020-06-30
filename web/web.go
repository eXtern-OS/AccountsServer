package web

import (
	"github.com/eXtern-OS/AMS"
	"github.com/gin-gonic/gin"
)

var SADDR, PseudoP string

func Init(saddr, pseudoP string) {
	SADDR = saddr
	PseudoP = pseudoP
}

func RenderImgPath(img string) string {
	return "http://" + SADDR + "/" + PseudoP + "/" + img
}

func RenderIndex(uid string) gin.H {
	acc := AMS.GetUserByID(uid)
	return gin.H{
		"avatar_url":      RenderImgPath(acc.AvatarURL),
		"name":            acc.Name,
		"registered_date": acc.Registered,
		"login":           acc.Login,
		"email":           acc.Email,
		"website":         acc.Website,
	}
}
