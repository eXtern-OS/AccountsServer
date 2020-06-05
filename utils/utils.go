package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"time"
)

func Makehash(data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func KeyOffset() string {
	t := time.Now()
	m := int(t.Month())
	l := "1"

	if m == 1 {
		return strconv.Itoa(t.Year()-1) + "December"
	} else {
		con, err := time.Parse(l, strconv.Itoa(m-1))
		if err != nil {
			return ""
		}
		return strconv.Itoa(t.Year()) + con.Month().String()
	}
}

func KeySetOffset(offset int) (string, string) {
	t := time.Now()
	m := int(t.Month())
	l := "1"
	if m-offset < 0 {
		return "", ""
	} else if m == offset {
		return "", "December"
	} else {
		cv, err := time.Parse(l, strconv.Itoa(m-offset))
		if err != nil {
			return "", ""
		}
		return strconv.Itoa(t.Year()) + cv.Month().String(), cv.Month().String()
	}
}
