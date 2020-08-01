package server

import (
	"externos.io/AccountsServer/auth"
	"externos.io/AccountsServer/web"
	"github.com/eXtern-OS/AMS"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var websiteURL = ""

func handleMainGet(c *gin.Context) {
	if cid, err := c.Cookie("userID"); err == nil {
		if t, uid := auth.AuthenticateCookie(cid); t {
			if uid != "" {
				c.HTML(http.StatusOK, "index.html", web.RenderIndex(uid))
				return
			}
		}
	}
	c.Redirect(http.StatusTemporaryRedirect, "/login")
	c.Abort()
	return
}

func handleMainPost(c *gin.Context) {
	if cid, err := c.Cookie("userID"); err == nil {
		if t, uid := auth.AuthenticateCookie(cid); t {
			name := c.PostForm("name")
			username := c.PostForm("username")
			url := c.PostForm("avatar_url")
			addr := c.PostForm("website")
			pwd := c.PostForm("pwd")
			pwdc := c.PostForm("cpwd")

			if name == "" || url == "" || addr == "" || pwd == "" || pwdc == "" || pwd != pwdc {
				c.Status(http.StatusBadRequest)
				return
			}
			AMS.UpdateDatabase(name, username, url, pwd, uid)

			c.HTML(http.StatusOK, "index.html", web.RenderIndex(uid))
		}
	}
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
	return
}

func handleLoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
	return
}

func handleLoginPost(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	if t, uid := auth.GetUserIdByEmailAndPassword(login, password); t {
		c.SetCookie("userID", auth.NewCookie(uid), int(time.Now().Add(12*30*time.Hour).Unix()), "/", websiteURL, false, false)
		c.Redirect(http.StatusFound, "/")
		c.Abort()
		return
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	}
}

func handleRegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
	return
}

func handleRegisterPost(c *gin.Context) {
	name := c.PostForm("name")
	username := c.PostForm("username")
	website := c.PostForm("website")
	email := c.PostForm("email")
	avatarurl := c.PostForm("avatarURL")
	password := c.PostForm("password")

	if name == "" || username == "" || email == "" || password == "" {
		c.Redirect(http.StatusFound, "/register")
		c.Abort()
		return
	}
	AMS.Register(name, username, username, avatarurl, password, website, email)
	c.Redirect(http.StatusTemporaryRedirect, "/login")
	c.Abort()
	return
}
