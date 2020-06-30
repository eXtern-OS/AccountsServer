package auth

import (
	"externos.io/AccountsServer/utils"
	"fmt"
	"strconv"
	"time"
)

var C CookiesManager

func Init() {
	C.LoadCookiesManager()
}

func AuthenticateCookie(cookieID string) (bool, string) {
	return C.Load(cookieID)
}

func AuthenticatePassword(login, password string) (bool, string) {
	return GetUserIdByEmailAndPassword(login, password)
}

func NewCookie(uid string) string {
	t := utils.Makehash(uid + strconv.Itoa(int(time.Now().Unix())))
	fmt.Println(t, uid)
	C.Write(t, uid)
	return t
}
