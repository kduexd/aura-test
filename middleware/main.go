package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
)

// METHODS the methods of server accepted
var METHODS = []string{"GET", "POST", "PUT", "DELETE"}

// CORS the rules of server
var CORS = cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowMethods:    append(METHODS, "OPTIONS"),
	AllowHeaders: []string{"Content-Type", "Origin", "Authorization", "Upgrade", "Connection",
		"Accept-Encoding", "Accept-Language", "Host", "token", "X-Requested-With"},
	MaxAge: 6 * time.Hour,
})
