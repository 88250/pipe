package util

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionData struct {
	UID   uint   // User ID
	UName string // Username
	URole int    // User role
	BID   uint   // Blog ID
}

func (sd *SessionData) Save(c *gin.Context) error {
	session := sessions.Default(c)
	sessionDataBytes, err := json.Marshal(sd)
	if nil != err {
		return err
	}
	session.Set("data", string(sessionDataBytes))

	return session.Save()
}

func GetSession(c *gin.Context) *SessionData {
	session, exists := c.Get("session")
	if nil == session || !exists {
		return nil
	}

	return session.(*SessionData)
}
