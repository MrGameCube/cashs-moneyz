package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var db *gorm.DB

func main() {

	db = InitDB()
	defer db.Close()
	router := initializeGin()
	server := initializeHTTPServer(router)
	waitForShutdown(server)

	log.Println("Server stopped!")
}

func initializeGin() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLFiles("frontend/public/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/helloWorld", handleHelloWorld)

	router.Static("/static", "frontend/public/static")
	initGroupHandlers(router)
	return router
}

func initializeHTTPServer(router *gin.Engine) *http.Server {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		log.Printf("HTTP Server Listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	return server
}

func waitForShutdown(server *http.Server) {
	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	// Wartet auf Stoppsignal des OS (Bsp. Strg+C oder kill)
	<-quitChannel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
