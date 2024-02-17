package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaybaba8085/sms-verification/api"
)

func main() {
	serverInIt := gin.Default()
	app := api.Config{Server: serverInIt}
	app.Routes()
	serverInIt.Run(":8000")
}
