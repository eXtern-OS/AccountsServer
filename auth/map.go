package auth

import "sync"

type CookiesManager struct {
	mx sync.RWMutex
	m  map[string]string
}

type ExportedManager struct {
	Index int               `json:"index"`
	Map   map[string]string `json:"map"`
}

func (c *CookiesManager) Load(key string) (bool, string) {
	c.mx.RLock()
	val, ok := c.m[key]
	c.mx.RUnlock()
	return ok, val
}

func (c *CookiesManager) Write(key string, value string) {
	c.mx.Lock()
	c.m[key] = value
	c.mx.Unlock()
}

func (c *CookiesManager) Dump() ExportedManager {
	var ex ExportedManager
	ex.Index = 0
	ex.Map = make(map[string]string)
	c.mx.RLock()
	ex.Map = c.m
	c.mx.RUnlock()
	return ex
}
