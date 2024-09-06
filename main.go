package main

import (
	"aura-test/middleware"
	"aura-test/pkg/log"
	"fmt"

	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the mode of gin
	gin.SetMode(gin.DebugMode)
	g := gin.New()

	g.Use(ginzap.Ginzap(log.Logger, time.RFC3339, true))
	g.Use(middleware.CORS)
	g.Use(gin.Recovery())

	// Routing the gin
	Routing(g)

	// Log the server is online
	log.Info("server is online")

	// Start the server
	server := &http.Server{Addr: ":3000", ReadHeaderTimeout: 3 * time.Second, Handler: g}
	log.Fatal(fmt.Sprintf("Server in fatal. %v", server.ListenAndServe()))
}
