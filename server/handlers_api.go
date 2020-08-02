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

type RegisterCredentials struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Website   string `json:"website"`
	Email     string `json:"email"`
	Avatarurl string `json:"avatarurl"`
	Password  string `json:"password"`
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
		return
	} else {
		c.Status(http.StatusBadRequest)
	}
}

func handleApiRegister(c *gin.Context) {
	var rgc RegisterCredentials
	if c.BindJSON(&rgc) == nil {
		if rgc.Name == "" || rgc.Username == "" || rgc.Email == "" || rgc.Password == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		if !AMS.Register(rgc.Name, rgc.Username, rgc.Username, rgc.Avatarurl, rgc.Password, rgc.Website, rgc.Email) {
			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
		return
	}
}

func handleApiAlive(c *gin.Context) {
	c.Status(http.StatusOK)
}
