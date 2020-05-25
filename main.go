package main

import (
	"encoding/json"
	"fmt"
	"github.com/eXtern-OS/AMS"
	beatrix "github.com/eXtern-OS/Beatrix"
	tm "github.com/eXtern-OS/TokenMaster"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
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

	r.Run()
}
