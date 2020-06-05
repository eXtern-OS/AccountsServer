package auth

import (
	"../db"
	"../utils"
	"context"
	"fmt"
	"github.com/eXtern-OS/AMS"
	beatrix "github.com/eXtern-OS/Beatrix"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type DevToken struct {
	Id  string `json:"id"`
	UID string `json:"uid"`
}

func GetUserIdByEmailAndPassword(login, password string) (bool, string) {
	if t, c := db.NewDatabaseCollection("Users", "accounts", "GETUSERBYTOKENID"); t {
		hash := utils.Makehash(password)

		filter := bson.M{"email": login, "password": hash}

		fmt.Println(filter)

		var res AMS.Account

		if c.FindOne(context.Background(), filter).Decode(&res) == nil {
			return res.UID != "", res.UID
		} else {
			return false, ""
		}
	} else {
		return false, ""
	}
}

func (c *CookiesManager) LoadCookiesManager() {
	if t, collection := db.NewDBCollection("cookies", "LOADCOOKIESMANAGER"); t {
		filter := bson.M{"index": 0}

		var result ExportedManager

		if collection.FindOne(context.Background(), filter).Decode(&result) == nil {
			c.m = result.Map
			return
		} else {
			c.m = make(map[string]string)
			return
		}
	} else {
		c.m = make(map[string]string)
		return
	}
}

func (ex *ExportedManager) Dump() {
	if t, collection := db.NewDBCollection("cookies", "EXPORTEDMANAGER.DUMP"); t {
		if _, err := collection.InsertOne(context.Background(), ex); err == nil {
			return
		} else {
			log.Println(err)
			go beatrix.SendError("ERROR: CANNOT INSERT TO DB", "EXPORTEDMANAGER.DUMP")
			return
		}
	} else {
		return
	}

}
