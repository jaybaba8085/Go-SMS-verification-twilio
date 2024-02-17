package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaybaba8085/Go-SMS-verification-twilio/api"
)

func main() {
	router := gin.Default()
	app := api.Config{Router: router}
	app.Routes()
	router.Run(":8000")
}
