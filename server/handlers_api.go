package server

import (
	"github.com/eXtern-OS/AMS"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func handleApiNewToken(c *gin.Context) {
	var cr Credentials
	if c.BindJSON(&cr) == nil {
		if cr.Login == "" || cr.Password == "" {
			c.Status(http.StatusNoContent)
		} else {
			code, t := AMS.GetToken(cr.Login, cr.Password, c.ClientIP())
			if code != http.StatusOK {
				c.Status(code)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"token": t,
				})
			}
		}
	}
}

func handleApiAlive(c *gin.Context) {
	c.Status(http.StatusOK)
}
