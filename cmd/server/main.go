package main

import (
	"github.com/gin-gonic/gin"
	"transPact/pkg/server"
)

func main() {
	router := gin.Default()
	router.GET("/endpoint", server.GetResponse)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
