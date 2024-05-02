package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const key = "Key"

func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(key))
	return sessions.Sessions("SessionId", store)

}
