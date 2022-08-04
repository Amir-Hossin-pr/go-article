package main

import (
	"article/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.Run("localhost:6060")
}
