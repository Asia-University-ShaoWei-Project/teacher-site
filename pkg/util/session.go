package util

import "github.com/gin-contrib/sessions"

const (
	sessionTokenKey = "token"
)

func GetSessionToken(s sessions.Session) interface{} {
	return s.Get(sessionTokenKey)
}

func SetSessionToken(s sessions.Session, token string) {
	s.Set(sessionTokenKey, token)
}

func DeleteSessionToken(s sessions.Session) {
	s.Delete(sessionTokenKey)
}
