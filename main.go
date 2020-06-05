package main

import (
	"./auth"
	"./db"
	"./web"
	"encoding/json"
	"fmt"
	"github.com/eXtern-OS/AMS"
	beatrix "github.com/eXtern-OS/Beatrix"
	tm "github.com/eXtern-OS/TokenMaster"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Config struct {
	BeatrixToken   string `json:"beatrix-token"`
	BeatrixChannel string `json:"beatrix-channelID"`
	MongoURI       string `json:"mongo-uri"`
	SQLURI         string `json:"sql-uri"`
	WebsiteURL     string `json:"website_url"`
}

func LoadConfig() Config {
	var config Config
	configFile, err := os.Open("ams-credentials.json")
	if err != nil {
		log.Panic(err)
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	fmt.Println(config)
	err = configFile.Close()
	if err != nil {
		log.Panic(err)
	}
	return config
}

func main() {
	r := gin.Default()

	cx := LoadConfig()

	beatrix.Init("AccountsServer", cx.BeatrixToken, cx.BeatrixChannel)

	AMS.Init(cx.MongoURI, cx.SQLURI)

	tm.Init(cx.MongoURI)

	db.Init(cx.MongoURI)

	auth.Init()

	r.LoadHTMLGlob("static/*.html")
	r.Static("/assets", "./static/assets")

	r.GET("/alive", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/newToken", func(c *gin.Context) {
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
	})

	r.GET("/", func(c *gin.Context) {
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
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	})

	r.POST("/", func(c *gin.Context) {
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
	})

	r.POST("/login", func(c *gin.Context) {
		login := c.PostForm("login")
		password := c.PostForm("password")

		if t, uid := auth.GetUserIdByEmailAndPassword(login, password); t {
			c.SetCookie("userID", auth.NewCookie(uid), int(time.Now().Add(12*30*time.Hour).Unix()), "/", cx.WebsiteURL, false, false)
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{})
			return
		}
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
		return
	})

	r.POST("/register", func(c *gin.Context) {
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
	})

	r.Run()
}
