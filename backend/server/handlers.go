package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleHelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"hallo": true})
}
