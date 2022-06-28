package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := router.Run(); err != nil {
		log.Fatal("failed to start the server. err:", err)
	}
}
